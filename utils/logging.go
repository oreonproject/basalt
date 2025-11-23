package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func LogInit(fileName string) *log.Logger {
	logRoot := "./logs/"
	// Make the directory if it doesn't already exist
	err := os.MkdirAll(filepath.Dir(logRoot), 0777)
	if err != nil {
		log.Fatal("Error While Making Directory: ", err)
	}

	logPath := logRoot + fileName
	fmt.Println(logPath)
	// We attempt the open the log file with the path, In write only mode if it doesn't already exist
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}

	// We now set the log to output to the created logfile and our error output
	logger := log.New(io.MultiWriter(logFile, os.Stderr), "log: ", log.Lmsgprefix|log.Lshortfile|log.Ltime|log.Ldate)
	return logger
}
