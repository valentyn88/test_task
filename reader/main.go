package main

import (
	"os"
	"bufio"
	"strings"
	"test_task/reader/user"
	"fmt"
	"strconv"
	"github.com/golang/protobuf/proto"
)

func main()  {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough arguments to run reader service")
		os.Exit(1)
	}

	fPath := args[1]
	csvFile, err := os.Open(fPath)
	defer csvFile.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "csv file can't be opened by reader service: %s", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(csvFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:2] == "id" {
			continue
		}

		row := strings.Split(line, ",")
		if len(row) < 4 {
			fmt.Fprintln(os.Stderr, "csv file row contains less than 4 params")
			continue
		}

		id, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not convert row id to int64: %s", err)
			continue
		}

		u := &user.User{
			Id:id,
			Name:row[1],
			Email:row[2],
			MobileNumber:row[3],
		}

		fmt.Print(proto.MarshalTextString(u))

		//if err := client.RequestConsumer(u); err != nil {
		//	log.Printf("An error occured while sending data to consumer service %s", err)
		//}
	}

	//if err := scanner.Err(); err != nil {
	//	log.Fatalf("An error occured while reading csv file by reader service %s", err)
	//}

	fmt.Printf("File %s was successfully imported\n", csvFile.Name())
}
