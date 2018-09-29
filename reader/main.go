package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"test_task/user"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "path to file as second param should be provided to run reader service")
		os.Exit(1)
	}

	fPath := args[1]
	csvFile, err := os.Open(fPath)
	defer csvFile.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open file by reader service: %s %s", fPath, err)
		os.Exit(1)
	}

	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to consumer service %s", err)
		os.Exit(1)
	}
	client := user.NewUsersClient(conn)
	ctx := context.Background()

	scanner := bufio.NewScanner(csvFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:2] == "id" {
			continue
		}

		u, err := user.ConvStr2User(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		usr, err := client.Add(ctx, u)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not add user: %s", err)
			continue
		}

		fmt.Fprintf(os.Stdin, "user with id: %d was successfully added\n", usr.Id)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occured while reading csv file by reader service %s %s", csvFile.Name(), err)
	}

	fmt.Printf("file '%s' was successfully imported\n", csvFile.Name())
}
