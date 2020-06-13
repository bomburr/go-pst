// This file is part of go-pst (https://www.go-pst.org/)
// Copyright (C) 2020 Marten Mooij (https://www.mooijtech.com/)
package pst

import (
	log "github.com/sirupsen/logrus"
)

// Parses the given file
func ParseFile(path string) {
	pstFile := NewPSTFile(path)

	log.Infof("Using file: %s...", pstFile.Path)

	// Get the file header
	fileHeader, err := pstFile.GetHeader()

	if err != nil {
		log.Fatalf("Failed to get file header: %s", err)
	}

	if !pstFile.IsValidSignature(fileHeader) {
		log.Fatal("Invalid file signature.")
	}

	// Get the content type
	fileContentType, err := pstFile.GetContentType(fileHeader)

	if err != nil {
		log.Fatalf("Failed to get content type: %s", err)
	}

	if fileContentType == ContentTypePST {
		log.Info("Identified content type as Personal Storage Table (PST).")
	} else if fileContentType == ContentTypeOST {
		log.Info("Identified content type as Offline Storage Table (OST).")
	} else if fileContentType == ContentTypePAB {
		log.Info("Identified content type as Public Address Book (PAB).")
	}

	// Get the format type
	fileFormatType, err := pstFile.GetFormatType(fileHeader)

	if err != nil {
		log.Fatalf("Failed to get format type: %s", err)
	}

	if fileFormatType == FormatType64 {
		log.Info("Identified format type as 64-bit (Unicode).")
	} else if fileFormatType == FormatType64With4k {
		log.Info("Identified format type as 64-bit with 4k pages.")
	} else if fileFormatType == FormatType32 {
		log.Info("Identified format type as 32-bit (ANSI).")
	}

	// Get the file header data
	fileHeaderData, err := pstFile.GetHeaderData(fileFormatType)

	if err != nil {
		log.Fatalf("Failed to get header data: %s", err)
	}

	// Get file encryption type
	fileEncryptionType, err := pstFile.GetEncryptionType(fileHeaderData, fileFormatType)

	if err != nil {
		log.Fatalf("Failed to get encryption type: %s", err)
	}

	if fileEncryptionType == EncryptionTypeNone {
		log.Info("Identified encryption type as none.")
	} else if fileEncryptionType == EncryptionTypePermute {
		log.Info("Identified encryption type as permute.")
	} else if fileEncryptionType == EncryptionTypeCyclic {
		log.Info("Identified encryption type as cyclic.")
	}

	// Walk the B-Tree
	fileBTreeStartOffset, err := pstFile.GetBTreeStartOffset(fileHeaderData, fileFormatType)

	if err != nil {
		log.Fatalf("Failed to get b-tree start offset: %s", err)
	}

	log.Infof("Walking b-tree at start offset: %d...", fileBTreeStartOffset)
}