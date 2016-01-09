package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func prettyPrint(value interface{}) {
	bytes, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
