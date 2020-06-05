// This file is part of go-pst (https://www.go-pst.org/)
// Copyright (C) 2020 Marten Mooij (https://www.mooijtech.com/)
package pst

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
)

// This struct represents a PST file.
type PSTFile struct {
	Path string
}

// Constructor for creating PST files.
func NewPSTFile(pstFilePath string) PSTFile {
	return PSTFile{
		Path: pstFilePath,
	}
}

// The file header common to both the 32-bit and 64-bit PFF format consists of 24 bytes.
func (pstFile *PSTFile) GetHeader() []byte {
	inputFile, err := os.Open(pstFile.Path)

	if err != nil {
		log.Fatalf("Failed to open file: %s", pstFile.Path)
	}

	outputBuffer := make([]byte, 24)
	count, err := inputFile.Read(outputBuffer)

	if err != nil {
		log.Fatalf("Failed to read file (%d of 24 bytes read).", count)
	}

	return outputBuffer
}

// The first 4 bytes of the file header contain the unique signature "!BDN" signifying the PFF format.
func (pstFile *PSTFile) IsValidSignature(fileHeader []byte) bool {
	return bytes.HasPrefix(fileHeader, []byte("!BDN"))
}

// Constants for identifying content types (PST, OST or PAB).
const (
	ContentTypePST = "SM"
	ContentTypeOST = "SO"
	ContentTypePAB = "AB"
)

// The 9th and 10th byte of the file header contains the content type.
// The content type signifies if the file contains the PST, OST or PAB format.
func (pstFile *PSTFile) GetContentType(fileHeader []byte) string {
	return string(fileHeader[8:10])
}

// Constants for identifying format types (32-bit or 64-bit).
const (
	FormatType32 = "32"
	FormatType64 = "64"
)

// The 11th and 12th byte of the file header contains the format type.
// This can be either 32-bit (ANSI) or 64-bit (Unicode).
func (pstFile *PSTFile) GetFormatType(fileHeader []byte) string {
	formatType := fileHeader[10:12]

	// Values from "2.2. Format types"
	if bytes.Equal(formatType, []byte{14, 0}) {
		return FormatType32
	} else if bytes.Equal(formatType, []byte{15, 0}) {
		return FormatType32
	} else if bytes.Equal(formatType, []byte{21, 0}) {
		return FormatType64
	} else if bytes.Equal(formatType, []byte{23, 0}) {
		return FormatType64
	} else if bytes.Equal(formatType, []byte{36, 0}) {
		return FormatType64
	} else {
		return "unknown"
	}
}

// The file header data bytes size may be 540 (64-bit) or 488 (32-bit).
func (pstFile *PSTFile) GetHeaderData(formatType string) []byte {
	inputFile, err := os.Open(pstFile.Path)

	if err != nil {
		log.Fatalf("Failed to open file: %s", pstFile.Path)
	}

	outputBufferSize := 540

	// 32-bit
	if formatType == FormatType32 {
		outputBufferSize = 488
	}

	outputBuffer := make([]byte, outputBufferSize)
	count, err := inputFile.Read(outputBuffer)

	if err != nil {
		log.Fatalf("Failed to read file (%d of %d bytes read).", count, outputBufferSize)
	}

	return outputBuffer
}

// Constants for identifying encryption types.
const (
	EncryptionTypeNone = "none"
	EncryptionTypePermute = "permute"
	EncryptionTypeCyclic = "cyclic"
)

// Reads the encryption type.
// Compressible encryption (permute) is on by default with newer versions of Outlook.
func (pstFile *PSTFile) GetEncryptionType(fileHeaderData []byte) string {
	encryptionType := fileHeaderData[513:514]

	// 32-bit
	if len(fileHeaderData) == 488 {
		encryptionType = fileHeaderData[461:462]
	}

	if bytes.Equal(encryptionType, []byte{0}) {
		return EncryptionTypeNone
	} else if bytes.Equal(encryptionType, []byte{1}) {
		return EncryptionTypePermute
	} else if bytes.Equal(encryptionType, []byte{2}) {
		return EncryptionTypeCyclic
	} else {
		return "unknown"
	}
}

// Returns the offset where the b-tree starts.
func (pstFile *PSTFile) GetBTreeStartOffset(fileHeaderData []byte) int {
	if pstFile.GetFormatType(fileHeaderData) == FormatType64 {
		return int(binary.LittleEndian.Uint64(fileHeaderData[240:248]))
	} else {
		return int(binary.LittleEndian.Uint32(fileHeaderData[196:200]))
	}
}

// Walks the b-tree at the given index.
// Returns th
func (pstFile *PSTFile) FindBTreeItem() {
	// The 64-bit/32-bit page is 512 bytes of size.
	pageSize := 496


}