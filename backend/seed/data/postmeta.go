package data

import (
	"context"
	"lets-go/db"
)

func SeedPostMeta() {

	client := db.NewClient()
	client.Connect()

	posts, err := client.Post.FindMany().Exec(context.Background())
	if err != nil {
		panic(err)
	}

	for _, post := range posts {
		print("Creating post meta for post: ", post.Title, "\n")

		client.PostMeta.CreateOne(
			db.PostMeta.Post.Link(db.Post.ID.Equals(post.ID)),
			db.PostMeta.Key.Set("author"),
			db.PostMeta.Value.Set("John Doe"),
		).Exec(context.Background())
	}

	client.Disconnect()
}
