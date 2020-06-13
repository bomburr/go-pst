// This file is part of go-pst (https://www.go-pst.org/)
// Copyright (C) 2020 Marten Mooij (https://www.mooijtech.com/)
package main

import (
	"flag"
	"github.com/mooijtech/go-pst/pst"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	inputFile := flag.String("parse", "", "File path to parse")
	logLevel := flag.String("log", "debug", "Set the logging level (info, warn, fatal, error, debug or trace)")

	flag.Parse()

	// Input validation
	if *inputFile == "" {
		log.Fatal("Please specify a file path with the \"-parse\" flag.")
	} else if _, err := os.Stat(*inputFile); os.IsNotExist(err) {
		log.Fatal("The specified file path does not exist.")
	}

	// Set logging level
	switch *logLevel {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}

	log.Infof("Starting go-pst v%s...", pst.Version)
	log.Infof("Using logging level: %s...", log.GetLevel().String())

	pst.ParseFile(*inputFile)
}
