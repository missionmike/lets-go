package api

import "os"

// sslmode=disable is needed on localhost because the database is not configured to use SSL.
// This is not secure for production, but it's fine for local development.
// If this were a project set up for deployment, I'd conditionally set the sslmode based on the environment.
// For example, if the environment is "production", then sslmode would be set to "require".
var DATABASE_URL = os.Getenv("DATABASE_URL") + "?sslmode=disable"

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
