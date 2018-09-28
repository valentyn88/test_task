package client

import (
	"test_task/reader/user"
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
)

// RequestConsumer - request consumer service
func RequestConsumer(u user.User) error {
	url := "http://localhost:9090"

	jsonStr, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("Error while encoding user data %s", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while making request to consumer service %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Consumer service is not accesible")
	}

	return nil
}
