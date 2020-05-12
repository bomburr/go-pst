// This file is part of go-pst (https://github.com/mooijtech/go-pst)
// Copyright (C) 2020 Marten Mooij (https://www.mooijtech.com/)
package main

import "log"

func main() {
	version := "0.0.1"
	pstFile := PSTFile{"/home/bot/Documents/test.pst"}

	log.Printf("Starting go-pst v%s...", version)
	log.Printf("Using file: %s...", pstFile.Path)

	fileHeader := ReadFileHeader(pstFile)

	if !IsValidSignature(fileHeader) {
		log.Fatalf("Invalid file signature!")
	}

	fileContentType := ReadContentType(fileHeader)

	if fileContentType == ContentTypePST {
		log.Println("Identified as Personal Storage Table (PST).")
	} else if fileContentType == ContentTypeOST {
		log.Println("Identified as Offline Storage Table (OST).")
	} else if fileContentType == ContentTypePAB {
		log.Println("Identified as Public Address Book (PAB).")
	} else {
		log.Fatalf("Failed to identify content type.")
	}
}