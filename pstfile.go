// This file is part of go-pst (https://github.com/mooijtech/go-pst)
// Copyright (C) 2020 Marten Mooij (https://www.mooijtech.com/)
package main

import (
	"bytes"
	"log"
	"os"
)

// This struct represents a PST file.
type PSTFile struct {
	Path string
}

// Reads the PST file header.
func ReadFileHeader(pstFile PSTFile) []byte {
	inputFile, err := os.Open(pstFile.Path)

	if err != nil {
		log.Fatalf("Failed to open file: %s", pstFile.Path)
	}

	outputBuffer := make([]byte, 16)
	count, err := inputFile.Read(outputBuffer)

	if err != nil {
		log.Fatalf("Failed to read file (%d of 16 bytes read).", count)
	}

	return outputBuffer
}

// The first 4 bytes of the file header contain the unique signature "!BDN" signifying the PFF format.
func IsValidSignature(fileHeader []byte) bool {
	return bytes.HasPrefix(fileHeader, []byte("!BDN"))
}

// Constants for identifying content types.
const (
	ContentTypePST = "SM"
	ContentTypeOST = "SO"
	ContentTypePAB = "AB"
)

// The 9th and 10th byte of the file header contains the content type.
// The content type signifies if the file contains the PST, OST or PAB format.
func ReadContentType(fileHeader []byte) string {
	return string(fileHeader[8:10])
}