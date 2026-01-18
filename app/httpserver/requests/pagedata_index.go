package requests

type PageDataIndexRequest struct {
	URL   *string `form:"url" binding:"omitempty,min=3,max=1000"`
	Title *string `form:"title" binding:"omitempty,min=3,max=1000"`
}
