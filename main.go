package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Config represents the application configuration
type Config struct {
	Port int `json:"port"`
}

// loadConfig reads configuration from config.json file
func loadConfig() Config {
	config := Config{Port: 80} // default port

	file, err := os.Open("config.json")
	if err != nil {
		log.Println("Config file not found, using default port 80")
		return config
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Println("Failed to parse config file, using default port 80")
		return config
	}

	log.Printf("Loaded configuration: port=%d\n", config.Port)
	return config
}

// handler returns 204 No Content for all requests
func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	config := loadConfig()

	http.HandleFunc("/", handler)

	addr := fmt.Sprintf(":%d", config.Port)
	log.Printf("Starting HTTP server on port %d\n", config.Port)
	log.Println("Server will respond with 204 No Content to all requests")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}