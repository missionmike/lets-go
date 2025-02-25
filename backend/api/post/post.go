package post

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"

	// Thanks to: https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
	_ "github.com/lib/pq"
)

// Post represents a blog post. If using Prisma ORM, this would be a model, and
// it would be generated automatically based on the schema.prisma definitions.
type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	CreatedAt string `json:"createdAt"` // In hindsight, snake_case may have been better.
	UpdatedAt string `json:"updatedAt"`
	Published bool   `json:"published"`
}

func GetAllPosts(c context.Context) ([]byte, error) {
	// sslmode=disable is needed on localhost because the database is not configured to use SSL.
	// This is not secure for production, but it's fine for local development.
	// If this were a project set up for deployment, I'd conditionally set the sslmode based on the environment.
	// For example, if the environment is "production", then sslmode would be set to "require".
	database_url := os.Getenv("DATABASE_URL") + "?sslmode=disable"

	db, err := sql.Open("postgres", database_url)
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
	rows, err := db.Query("SELECT id, title, \"desc\", \"createdAt\", \"updatedAt\", published FROM \"Post\"")
	if err != nil {
		log.Printf("Failed to query database: %v", err)
	}

	// Iterate over the rows and create a Post for each one.
	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Desc, &post.CreatedAt, &post.UpdatedAt, &post.Published); err != nil {
			log.Printf("Failed to scan row: %v", err)

			// TODO: Return a more user-friendly error message, and track the error in Sentry.
			return nil, err
		}
		posts = append(posts, post)
	}

	// Convert Post to JSON and return as JSON string.
	var postJSON []byte
	if postJSON, err = json.Marshal(posts); err != nil {
		log.Printf("Failed to marshal posts: %v", err)

		// TODO: Return a more user-friendly error message, and track the error in Sentry.
		return nil, err
	}

	return postJSON, nil
}
