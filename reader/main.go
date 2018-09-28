package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"test_task/reader/user"
	"test_task/reader/client"
)

func main()  {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Not enough arguments to run reader service")
	}

	fPath := args[1]
	csvFile, err := os.Open(fPath)
	defer csvFile.Close()
	if err != nil {
		log.Fatalf("Csv file can't be opened by reader service")
	}

	scanner := bufio.NewScanner(csvFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:2] == "id" {
			continue
		}
		row := strings.Split(line, ",")
		u := user.NewUser(row)
		if err := client.RequestConsumer(u); err != nil {
			log.Printf("An error occured while sending data to consumer service %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("An error occured while reading csv file by reader service %s", err)
	}
}
