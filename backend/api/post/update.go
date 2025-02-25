package post

import (
	"context"
	"database/sql"
	"lets-go/api"
	"log"

	_ "github.com/lib/pq"
)

func UpdatePost(c context.Context, id, title, desc string) error {
	db, err := sql.Open("postgres", api.DATABASE_URL)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	// First get the existing post
	var existingTitle, existingDesc string
	err = db.QueryRow(`
		SELECT title, "desc"
		FROM "Post"
		WHERE id = $1
	`, id).Scan(&existingTitle, &existingDesc)
	if err != nil {
		log.Printf("Failed to get existing post: %v", err)
		return err
	}

	// Use existing values if new values are empty
	if title == "" {
		title = existingTitle
	}
	if desc == "" {
		desc = existingDesc
	}

	_, err = db.Exec(`
		UPDATE "Post"
		SET title = $1, "desc" = $2, "updatedAt" = NOW()
		WHERE id = $3
	`, title, desc, id)
	if err != nil {
		log.Printf("Failed to update post: %v", err)
		return err
	}

	return nil
}
