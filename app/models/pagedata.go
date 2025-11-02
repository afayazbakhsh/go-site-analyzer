package models

type PageLinks struct {
	Internal int
	External int
}

type PageData struct {
	ID          uint `gorm:"primaryKey"`
	URL         string
	Title       string
	Description string
	WordCount   int
	Links       PageLinks `gorm:"embedded"`
	StatusCode  int
	LoadTime    float64
	Language    string
}
