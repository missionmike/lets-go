package post

import (
	"context"
	"database/sql"
	"lets-go/api"
	"log"

	_ "github.com/lib/pq"
)

func DeletePost(c context.Context, id string) error {
	db, err := sql.Open("postgres", api.DATABASE_URL)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	// TODO: Add a conditional check to ensure the post exists before deleting it.

	// First, delete all related PostMeta values.
	_, err = db.Exec(`
		DELETE FROM "PostMeta"
		WHERE "postId" = $1
	`, id)
	if err != nil {
		log.Printf("Failed to delete post meta: %v", err)
		return err
	}

	_, err = db.Exec(`
		DELETE FROM "Post"
		WHERE id = $1
	`, id)
	if err != nil {
		log.Printf("Failed to delete post: %v", err)
		return err
	}

	return nil
}
