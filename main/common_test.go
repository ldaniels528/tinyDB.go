package main

import (
	"log"
	"strconv"
	"testing"
)

func TestMapF(t *testing.T) {
	// define the input collection
	items := []string{"100", "200", "300"}
	log.Printf("items: %s\n", AsJSON(items))

	// transform the input collection
	newItems, err := MapF(items, func(input string) (int, error) {
		return strconv.Atoi(input)
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("newItems: %s\n", AsJSON(newItems))

	// validate the results
	if len(newItems) != len(items) {
		log.Fatalf("input/output size mismatch (input: %d, output: %d)\n", len(items), len(newItems))
	}
	for i, item := range items {
		value, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err.Error())
		}
		if value != newItems[i] {
			log.Fatalf("item[%d] mismatch: %d was not %d\n", i, value, newItems[i])
		}
	}
}

func TestToByte(t *testing.T) {
	if ToByte(true) != 1 {
		log.Fatal("ToByte(true) did not produce 1")
	}
	if ToByte(false) != 0 {
		log.Fatal("ToByte(false) did not produce 0")
	}
}
