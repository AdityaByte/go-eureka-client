package main

import (
	"fmt"
	"log"
	"net/http"
)

func Disconnect(appName string, instanceId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8761/eureka/apps/%s/%s", appName, instanceId), nil)
	if err != nil {
		return fmt.Errorf("ERROR: Failed to create the request: %w", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ERROR: Failed to make the request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("ERROR: Invalid response status: %s", resp.Status)
	}

	log.Println("Successfully disconnected from the eureka server")
	return nil
}
