package main

import (
	"fmt"

	"github.com/AnnonaOrg/pkg/fifomap"
)

func main() {
	fifoMap := fifomap.NewFIFOMap()

	fifoMap.Set("key1", "value1")
	fifoMap.Set("key2", "value2")

	fmt.Println("Before removing oldest:")
	for _, key := range fifoMap.keys {
		fmt.Printf("Key: %s, Value: %v\n", key, fifoMap.items[key])
	}

	fifoMap.RemoveOldest()

	fmt.Println("\nAfter removing oldest:")
	for _, key := range fifoMap.keys {
		fmt.Printf("Key: %s, Value: %v\n", key, fifoMap.items[key])
	}
}
