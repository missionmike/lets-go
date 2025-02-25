package post

import (
	"context"
	"database/sql"
	"encoding/json"
	"lets-go/api"
	"log"

	// Thanks to: https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
	_ "github.com/lib/pq"
)

func GetAllPosts(c context.Context) ([]byte, error) {
	db, err := sql.Open("postgres", api.DATABASE_URL)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)

		// TODO: Return a more user-friendly error message, and track the error in Sentry.
		return nil, err
	}
	defer db.Close()

	// Query the database for all posts. There's probably a nicer way to format
	// this query in Go, but I'm not sure what it is yet...
	// In hindsight, table names, columns, etc. should have been snake_case with no capitalization.
	// This would make it easier to query directly with SQL. This differs from using Prisma with TypeScript in
	// that the Prisma client abstracts away the SQL queries, so the naming convention doesn't matter as much.
	rows, err := db.Query(`
		SELECT id, title, "desc", "createdAt", "updatedAt", published 
		FROM "Post"
	`)
	if err != nil {
		log.Printf("Failed to query database: %v", err)
	}

	// Iterate over the rows and create a Post for each one.
	var posts []api.Post
	for rows.Next() {
		var post api.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Desc,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Published,
		); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	var postJSON []byte
	if postJSON, err = json.Marshal(posts); err != nil {
		log.Printf("Failed to marshal posts: %v", err)
		return nil, err
	}

	return postJSON, nil
}

func GetPostByID(c context.Context, id string) ([]byte, error) {
	db, err := sql.Open("postgres", api.DATABASE_URL)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	row, err := db.Query(`
		SELECT id, title, "desc", "createdAt", "updatedAt", published 
		FROM "Post" 
		WHERE id = $1
	`, id)
	if err != nil {
		log.Printf("Failed to query database: %v", err)
	}

	// If row is empty, return an error.
	if !row.Next() {
		log.Printf("Post not found")
		return nil, err
	}

	var post api.Post
	if row.Next() {
		if err := row.Scan(
			&post.ID,
			&post.Title,
			&post.Desc,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Published,
		); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, err
		}
	}

	var postJSON []byte
	if postJSON, err = json.Marshal(post); err != nil {
		log.Printf("Failed to marshal post: %v", err)
		return nil, err
	}

	return postJSON, nil
}
