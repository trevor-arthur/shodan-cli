package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/trevor-arthur/shodan/shodan"
)

func main() {
	// Load in API key from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("[!] Error loading .env file")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")

	// Anything longer than 2 words ('shodan <searchterm>') will return usage help
	if len(os.Args) != 2 {
		log.Fatalln("[*] Usage: shodan <searchterm>")
	}

	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits: %d\n\n",
		info.QueryCredits,
		info.ScanCredits,
	)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
