package data

import (
	"context"
	"lets-go/db"
	"lets-go/util/helpers"
	"log"
)

func SeedPosts() {

	posts := []db.InnerPost{
		{
			Title: "First Post",
			Desc:  helpers.Ptr("This is the content of the first post."),
		},
		{
			Title: "Second Post",
			Desc:  helpers.Ptr("This is the content of the second post."),
		},
	}

	client := db.NewClient()
	if err := client.Connect(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	// Delete existing posts.
	log.Printf("Finding existing posts.\n")
	postsToDelete, err := client.Post.FindMany().Exec(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Found %d existing posts.\n", len(postsToDelete))

	// Delete existing posts, if any.
	for _, post := range postsToDelete {
		log.Printf("Deleting post ID: %s, (%s)\n", post.ID, post.Title)

		// Delete related PostMeta records first.
		_, err := client.Prisma.ExecuteRaw("DELETE FROM \"PostMeta\" WHERE \"postId\" = $1", post.ID).Exec(ctx)
		if err != nil {
			log.Printf("Error deleting post meta: %v\n", err)
			panic(err)
		}

		// Then delete the post.
		_, err = client.Prisma.ExecuteRaw("DELETE FROM \"Post\" WHERE id = $1", post.ID).Exec(ctx)

		if err != nil {
			log.Printf("Error deleting post: %v\n", err)
			panic(err)
		}
	}

	for _, post := range posts {
		log.Printf("Creating post: %s\n", post.Title)

		if _, err := client.Post.CreateOne(
			db.Post.Title.Set(post.Title),
			db.Post.Published.Set(true),
			db.Post.Desc.Set(*post.Desc),
		).Exec(ctx); err != nil {
			log.Printf("Error creating post: %v\n", err)
			panic(err)
		}
	}

	if err := client.Disconnect(); err != nil {
		panic(err)
	}
}
