package main

import (
	"awesomeProject/buggy"
	"awesomeProject/fixed"
	"fmt"
)

func main() {
	// Buggy config - loading won't fail even if PRIVATE_KEY env variable is missing
	buggyConfig, err := buggy.LoadConfig()

	if err == nil {
		fmt.Printf("Server URL: %s\n", buggyConfig.GetServerURL())
		fmt.Printf("Port: %d\n", buggyConfig.GetAppPort())
		fmt.Printf("PrivateKey: %s\n", buggyConfig.GetPrivateKey())
	}

	// Fixed config - will panic if PRIVATE_KEY env variable is missing
	config, err := fixed.LoadConfig()

	if err == nil {
		fmt.Printf("Server URL: %s\n", config.GetServerURL())
		fmt.Printf("Port: %d\n", config.GetAppPort())
		fmt.Printf("PrivateKey: %s\n", config.GetPrivateKey())
	}
}
