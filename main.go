package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	file := flag.String("file", "", "origin file that need to be filtered")

	flag.Parse()
	line := flag.Args()
	if len(*file) == 0 {
		fmt.Println("Usage: log_filter.go -file")
		flag.PrintDefaults()
		os.Exit(1)
	}

	logToFilter := *file
	lineToFilter := line

	saveFileName := logToFilter + "-FILTERED-" + time.Now().Format(time.StampMilli)
	fmt.Println(lineToFilter)

	var scanner *bufio.Scanner

	openedFile, _ := os.Open("./" + logToFilter)
	scanner = bufio.NewScanner(openedFile)
	for scanner.Scan() {

		currentLine := scanner.Text()
		for _, lineToFilter := range line {
			if strings.Contains(currentLine, lineToFilter) {
				appendToLog(currentLine, saveFileName)
				fmt.Println(currentLine)
			}
		}

	}
}

// This func will make file if doesn't exist, if does it will append log data to it
func appendToLog(logLine, logFile string) {

	f, err := os.OpenFile("./"+logFile+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(logLine + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
