package main

import (
	"fmt"
	"log"

	"../internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Can't load Config: %s", err)
	}
	fmt.Printf("config: %+v", cfg)
}
