package entity

type Post struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title"`
	Body  string `json:"body"`
	// ContributedBy uint      `json:"contributed_by"`
	// ContributedAt time.Time `json:"contributed_at"`
	// Deleted       bool      `json:"deleted"`
}
