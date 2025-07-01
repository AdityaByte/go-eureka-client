package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func registerService(serviceName string, ipAddr string, host string, port int) error {
	register_payload := RegisterPayload{
		HostName:   host,
		App:        serviceName,
		IpAddr:     ipAddr,
		VipAddress: serviceName,
		Status:     "UP",
		Port: Port{
			Port: port,
			Enabled: "true",
		},
		DataCenterInfo: DataCenterInfo{
			Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
			Name: "MyOwn",
		},
	}

	instance := InstancePayload{
		Instance: register_payload,
	}

	jsonData, err := json.Marshal(instance)
	if err != nil {
		return fmt.Errorf("ERROR: Failed to encode the data; more likely an internal server error %w", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8761/eureka/apps/%s", serviceName), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("ERROR: Failed to make the http request")
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	log.Println("Response status: %s", resp.Status)
	return nil
}
