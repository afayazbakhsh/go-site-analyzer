package response

type PageDataIndexResponse struct {
	ID          int       `json:"id"`
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	WordCount   int       `json:"word_count"`
	Links       PageLinks `json:"links" gorm:"-"`
	StatusCode  int       `json:"status_code"`
	LoadTime    int       `json:"load_time"`
	Language    string    `json:"language"`
	// User   User `gorm:"foreignKey:UserID"`
	// db.Preload("Posts").First(&userWithPosts, 1)
}

type PageLinks struct {
	Internal int `json:"internal"`
	External int `json:"external"`
}
