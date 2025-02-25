package post

import (
	"context"
	"database/sql"
	"lets-go/api"
	"log"

	"github.com/google/uuid"

	// Thanks to: https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
	_ "github.com/lib/pq"
)

func CreatePost(c context.Context, title, desc string) (string, error) {
	db, err := sql.Open("postgres", api.DATABASE_URL)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return "", err
	}
	var id string

	err = db.QueryRow(`
		INSERT INTO "Post" (id, title, "desc", published)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, uuid.New().String(), title, desc, true).Scan(&id)
	if err != nil {
		log.Printf("Failed to insert post: %v", err)
		return "", err
	}

	return id, nil
}
