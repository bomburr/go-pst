// This file is part of go-pst (https://github.com/mooijtech/go-pst)
// Copyright (C) 2020 Marten Mooij (https://www.mooijtech.com/)
package main

import "log"

func main() {
	version := "0.0.1"
	pstFile := PSTFile{"/home/bot/Documents/test.pst"}

	log.Printf("Starting go-pst v%s...", version)
	log.Printf("Using file: %s...", pstFile.Path)

	fileHeader := ReadHeader(pstFile)

	if !IsValidSignature(fileHeader) {
		log.Fatalf("Invalid file signature!")
	}

	fileContentType := ReadContentType(fileHeader)

	if fileContentType == ContentTypePST {
		log.Println("Identified content type as Personal Storage Table (PST).")
	} else if fileContentType == ContentTypeOST {
		log.Println("Identified content type as Offline Storage Table (OST).")
	} else if fileContentType == ContentTypePAB {
		log.Println("Identified content type as Public Address Book (PAB).")
	} else {
		log.Fatalf("Failed to identify content type.")
	}

	fileFormatType := ReadFormatType(fileHeader)

	if fileFormatType == FormatType64 {
		log.Println("Identified format type as 64-bit (Unicode).")
	} else if fileFormatType == FormatType32 {
		log.Println("Identified format type as 32-bit (ANSI).")
	} else {
		log.Fatalf("Failed to identify format type.")
	}

	fileHeaderData := ReadHeaderData(pstFile, fileFormatType)
	fileEncryptionType := ReadEncryptionType(fileHeaderData)

	if fileEncryptionType == EncryptionTypeNone {
		log.Println("Identified encryption type as none.")
	} else if fileEncryptionType == EncryptionTypePermute {
		log.Println("Identified encryption type as permute.")
	} else if fileEncryptionType == EncryptionTypeCyclic {
		log.Println("Identified encryption type as cyclic.")
	} else {
		log.Fatal("Failed to identify encryption type.")
	}
}