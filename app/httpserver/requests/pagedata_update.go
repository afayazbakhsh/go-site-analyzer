package requests

type PageDataUpdateRequest struct {
	URL   *string `json:"url" binding:"min=1,max=1000"`
	Title *string `json:"title" binding:"min=1,max=255"`
}
