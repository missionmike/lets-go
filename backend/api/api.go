package api

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
