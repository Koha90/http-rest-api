package main

import (
	"fmt"
	"log"

	"github.com/koha90/http-rest-api/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); if err != nil {
		log.Fatal(err)
	}
}
