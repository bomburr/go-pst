// This file is part of go-pst (https://www.go-pst.org/)
// Copyright (C) 2020 Marten Mooij (https://www.mooijtech.com/)
package pst

import (
	"bytes"
	"encoding/binary"
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

// This struct represents a PST file.
type ParsableFile struct {
	Path string
}

// Constructor for creating PST files.
func NewPSTFile(pstFilePath string) ParsableFile {
	return ParsableFile{
		Path: pstFilePath,
	}
}

// The file header common to both the 32-bit and 64-bit PFF format consists of 24 bytes.
func (pstFile *ParsableFile) GetHeader() ([]byte, error) {
	inputFile, err := os.Open(pstFile.Path)

	if err != nil {
		log.Errorf("Failed to open file: %s", pstFile.Path)
		return nil, err
	}

	outputBuffer := make([]byte, 24)
	count, err := inputFile.Read(outputBuffer)

	if err != nil {
		log.Errorf("Failed to read file (%d of 24 bytes read).", count)
		return nil, err
	}

	if err := inputFile.Close(); err != nil {
		log.Errorf("Failed to close file: %s", err)
		return nil, err
	}

	return outputBuffer, nil
}

// The first 4 bytes of the file header contain the unique signature "!BDN" signifying the PFF format.
func (pstFile *ParsableFile) IsValidSignature(fileHeader []byte) bool {
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
func (pstFile *ParsableFile) GetContentType(fileHeader []byte) string {
	return string(fileHeader[8:10])
}

// Constants for identifying format types (32-bit or 64-bit).
const (
	FormatType32 = "32"
	FormatType64 = "64"
)

// The 11th and 12th byte of the file header contains the format type.
// This can be either 32-bit (ANSI) or 64-bit (Unicode).
func (pstFile *ParsableFile) GetFormatType(fileHeader []byte) (string, error) {
	formatType := fileHeader[10:12]

	// References "2.2. Format types" from the file format specification
	if bytes.Equal(formatType, []byte{14, 0}) {
		return FormatType32, nil
	} else if bytes.Equal(formatType, []byte{15, 0}) {
		return FormatType32, nil
	} else if bytes.Equal(formatType, []byte{21, 0}) {
		return FormatType64, nil
	} else if bytes.Equal(formatType, []byte{23, 0}) {
		return FormatType64, nil
	} else if bytes.Equal(formatType, []byte{36, 0}) {
		return FormatType64, nil
	} else {
		return "", errors.New("failed to get format type")
	}
}

// The file header data bytes size may be 540 (64-bit) or 488 (32-bit).
func (pstFile *ParsableFile) GetHeaderData(formatType string) ([]byte, error) {
	inputFile, err := os.Open(pstFile.Path)

	if err != nil {
		log.Errorf("Failed to open file: %s", pstFile.Path)
		return nil, err
	}

	// File header output buffer
	var outputBufferSize int

	if formatType == FormatType64 {
		outputBufferSize = 540
	} else if formatType == FormatType32 {
		outputBufferSize = 488
	} else {
		return nil, errors.New("unsupported format type")
	}

	outputBuffer := make([]byte, outputBufferSize)
	count, err := inputFile.Read(outputBuffer)

	if err != nil {
		log.Errorf("Failed to read file (%d of %d bytes read).", count, outputBufferSize)
		return nil, err
	}

	if err := inputFile.Close(); err != nil {
		log.Errorf("Failed to close file: %s", err)
		return nil, err
	}

	return outputBuffer, nil
}

// Constants for identifying encryption types.
const (
	EncryptionTypeNone = "none"
	EncryptionTypePermute = "permute"
	EncryptionTypeCyclic = "cyclic"
)

// Reads the encryption type.
// Compressible encryption (permute) is on by default with newer versions of Outlook.
func (pstFile *ParsableFile) GetEncryptionType(fileHeaderData []byte, formatType string) (string, error) {
	var encryptionType []byte

	if formatType == FormatType64 {
		encryptionType = fileHeaderData[513:514]
	} else if formatType == FormatType32 {
		encryptionType = fileHeaderData[461:462]
	} else {
		return "", errors.New("unsupported format type")
	}

	if bytes.Equal(encryptionType, []byte{0}) {
		return EncryptionTypeNone, nil
	} else if bytes.Equal(encryptionType, []byte{1}) {
		return EncryptionTypePermute, nil
	} else if bytes.Equal(encryptionType, []byte{2}) {
		return EncryptionTypeCyclic, nil
	} else {
		return "", errors.New("unsupported encryption type")
	}
}

// Returns the offset where the b-tree starts.
func (pstFile *ParsableFile) GetBTreeStartOffset(fileHeaderData []byte, formatType string) (int, error) {
	if formatType == FormatType64 {
		return int(binary.LittleEndian.Uint64(fileHeaderData[240:248])), nil
	} else if formatType == FormatType32 {
		return int(binary.LittleEndian.Uint32(fileHeaderData[196:200])), nil
	} else {
		return -1, errors.New("unsupported format type")
	}
}

