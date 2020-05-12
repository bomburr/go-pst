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

	log.Printf("File signature validated...")
}