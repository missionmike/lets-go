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
