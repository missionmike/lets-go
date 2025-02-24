package main

import (
	"lets-go/seed/data"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error: %s", r)
		}
	}()

	data.SeedPosts()
	data.SeedPostMeta()
}
