package data

import (
	"context"
	"lets-go/db"
	"log"
)

func SeedPostMeta() {

	client := db.NewClient()
	if err := client.Connect(); err != nil {
		panic(err)
	}

	posts, err := client.Post.FindMany().Exec(context.Background())
	if err != nil {
		panic(err)
	}

	for _, post := range posts {
		log.Printf("Creating post meta for post: %s\n", post.Title)

		if _, err := client.PostMeta.CreateOne(
			db.PostMeta.Post.Link(db.Post.ID.Equals(post.ID)),
			db.PostMeta.Key.Set("author"),
			db.PostMeta.Value.Set("John Doe"),
		).Exec(context.Background()); err != nil {
			log.Printf("Error creating post: %v\n", err)
			panic(err)
		}
	}

	if err := client.Disconnect(); err != nil {
		panic(err)
	}
}
