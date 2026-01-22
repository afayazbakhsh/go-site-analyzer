package requests

type UrlCheckRequest struct {
	Url *string `form:"url" binding:"required,min=5,max=500"`
}
