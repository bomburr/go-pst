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

// Returns true if the given file signature matches "!BDN" (signifying the PFF format).
func IsValidSignature(fileHeader []byte) bool {
	return bytes.HasPrefix(fileHeader, []byte("!BDN"))
}