// This file is part of go-pst (https://www.go-pst.org/)
// Copyright (C) 2020 Marten Mooij (https://www.mooijtech.com/)
package pst

import (
	"log"
)

func ParseFile(path string) {
	pstFile := NewPSTFile(path)

	log.Printf("Starting go-pst v%s...", version)
	log.Printf("Using file: %s...", pstFile.Path)

	fileHeader := pstFile.GetHeader()

	if !pstFile.IsValidSignature(fileHeader) {
		log.Fatalf("Invalid file signature!")
	}

	fileContentType := pstFile.GetContentType(fileHeader)

	if fileContentType == ContentTypePST {
		log.Println("Identified content type as Personal Storage Table (PST).")
	} else if fileContentType == ContentTypeOST {
		log.Println("Identified content type as Offline Storage Table (OST).")
	} else if fileContentType == ContentTypePAB {
		log.Println("Identified content type as Public Address Book (PAB).")
	} else {
		log.Fatalf("Failed to identify content type.")
	}

	fileFormatType := pstFile.GetFormatType(fileHeader)

	if fileFormatType == FormatType64 {
		log.Println("Identified format type as 64-bit (Unicode).")
	} else if fileFormatType == FormatType32 {
		log.Println("Identified format type as 32-bit (ANSI).")
	} else {
		log.Fatalf("Failed to identify format type.")
	}

	fileHeaderData := pstFile.GetHeaderData(fileFormatType)
	fileEncryptionType := pstFile.GetEncryptionType(fileHeaderData)

	if fileEncryptionType == EncryptionTypeNone {
		log.Println("Identified encryption type as none.")
	} else if fileEncryptionType == EncryptionTypePermute {
		log.Println("Identified encryption type as permute.")
	} else if fileEncryptionType == EncryptionTypeCyclic {
		log.Println("Identified encryption type as cyclic.")
	} else {
		log.Fatal("Failed to identify encryption type.")
	}

	fileBTreeStartOffset := pstFile.GetBTreeStartOffset(fileHeaderData)

	log.Printf("Walking b-tree at start offset: %d...", fileBTreeStartOffset)
}