package config

import (
	"fmt"
	"os"
)

func GetServePort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	return fmt.Sprintf(":%s", port)
}
