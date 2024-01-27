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
	for _, key := range fifoMap.Keys() {
		kk, ok := fifoMap.Get(key)
		fmt.Printf("Key: %s, Value(%v): %v \n", key, ok, kk)
	}

	fifoMap.RemoveOldest()

	fmt.Println("\nAfter removing oldest:")
	for _, key := range fifoMap.Keys() {
		kk, ok := fifoMap.Get(key)
		fmt.Printf("Key: %s, Value(%v): %v \n", key, ok, kk)
	}
}
