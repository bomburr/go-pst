// This file is part of go-pst (https://www.go-pst.org/)
// Copyright (C) 2020 Marten Mooij (https://www.mooijtech.com/)
package pst

import (
	"bytes"
	"encoding/binary"
	"errors"
	"os"
)

// ParsableFile represents a PST file.
type ParsableFile struct {
	Path string
}

// NewPSTFile is a constructor for creating PST files.
func NewPSTFile(pstFilePath string) ParsableFile {
	return ParsableFile{
		Path: pstFilePath,
	}
}

// Read reads the parsable file into the specified buffer.
func (parsableFile *ParsableFile) Read(outputBufferSize int, offset int) ([]byte, error) {
	inputFile, err := os.Open(parsableFile.Path)

	if err != nil {
		return nil, err
	}

	outputBuffer := make([]byte, outputBufferSize)

	_, err = inputFile.Seek(int64(offset), 0)

	if err != nil {
		return nil, err
	}

	_, err = inputFile.Read(outputBuffer)

	if err != nil {
		return nil, err
	}

	if err := inputFile.Close(); err != nil {
		return nil, err
	}

	return outputBuffer, nil
}

// GetHeader returns the file header.
// References "2. File header".
func (parsableFile *ParsableFile) GetHeader() ([]byte, error) {
	return parsableFile.Read(24, 0)
}

// IsValidSignature checks if the file header contains the unique signature "!BDN".
// References "2. File header".
func (parsableFile *ParsableFile) IsValidSignature(fileHeader []byte) bool {
	return bytes.HasPrefix(fileHeader, []byte("!BDN"))
}

// Constants for identifying content types (PST, OST or PAB).
// References "2.1. Content types".
const (
	ContentTypePST = "PST"
	ContentTypeOST = "PST"
	ContentTypePAB = "PAB"
)

// GetContentType returns the content type which may be PST, OST or PAB.
// References "2. File header".
func (parsableFile *ParsableFile) GetContentType(fileHeader []byte) (string, error) {
	contentType := fileHeader[8:10]

	if bytes.Equal(contentType, []byte("SM")) {
		return ContentTypePST, nil
	} else if bytes.Equal(contentType, []byte("SO")) {
		return ContentTypeOST, nil
	} else if bytes.Equal(contentType, []byte("AB")) {
		return ContentTypePAB, nil
	} else {
		return "", errors.New("failed to get content type")
	}
}

// Constants for identifying format types (32-bit or 64-bit).
const (
	FormatType32 = "32-bit"
	FormatType64 = "64-bit"
	FormatType64With4k = "64-bit-with-4k"
)

// GetFormatType returns the format type which can be either 32-bit (ANSI) or 64-bit (Unicode).
// References "2. File header" and "2.2. Format types".
func (parsableFile *ParsableFile) GetFormatType(fileHeader []byte) (string, error) {
	formatType := fileHeader[10:12]

	if bytes.Equal(formatType, []byte{14, 0}) {
		return FormatType32, nil
	} else if bytes.Equal(formatType, []byte{15, 0}) {
		return FormatType32, nil
	} else if bytes.Equal(formatType, []byte{21, 0}) {
		return FormatType64, nil
	} else if bytes.Equal(formatType, []byte{23, 0}) {
		return FormatType64, nil
	} else if bytes.Equal(formatType, []byte{36, 0}) {
		return FormatType64With4k, nil
	} else {
		return "", errors.New("failed to get format type")
	}
}

// GetHeaderData returns the file header data (in bytes).
// References "2.3. The 32-bit header data" and "2.4. The 64-bit header data".
func (parsableFile *ParsableFile) GetHeaderData(formatType string) ([]byte, error) {
	var outputBufferSize int

	if formatType == FormatType64 {
		outputBufferSize = 540
	} else if formatType == FormatType32 {
		outputBufferSize = 488
	} else {
		return nil, errors.New("unsupported format type")
	}

	return parsableFile.Read(outputBufferSize, 0)
}

// Constants for identifying encryption types.
const (
	EncryptionTypeNone = "none"
	EncryptionTypePermute = "permute"
	EncryptionTypeCyclic = "cyclic"
)

// GetEncryptionType returns the encryption type.
// Compressible encryption (permute) is on by default with newer versions of Outlook.
// References "2.3. The 32-bit header data", "2.4. The 64-bit header data" and "2.7. Encryption types".
func (parsableFile *ParsableFile) GetEncryptionType(fileHeaderData []byte, formatType string) (string, error) {
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

// GetBTreeStartOffset returns the start offset of the b-tree.
// References "2.3. The 32-bit header data" and "2.4. The 64-bit header data".
func (parsableFile *ParsableFile) GetBTreeStartOffset(fileHeaderData []byte, formatType string) (int, error) {
	if formatType == FormatType64 {
		return int(binary.LittleEndian.Uint64(fileHeaderData[240:248])), nil
	} else if formatType == FormatType32 {
		return int(binary.LittleEndian.Uint32(fileHeaderData[196:200])), nil
	} else {
		return -1, errors.New("unsupported format type")
	}
}

