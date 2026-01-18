package requests

import (
	"gocrawler/app/models"
	"unicode"

	"github.com/go-playground/validator/v10"
)

type PageDataCreateRequest struct {
	URL   *string           `json:"url" binding:"required,min=1,max=1000"`
	Links *models.PageLinks `json:"links" binding:"required"`
}


// نمونه: اعتبارسنجی پسورد قوی
func PasswordStrength(fl validator.FieldLevel) bool {
	pwd := fl.Field().String()
	// حداقل ۸ کاراکتر، حداقل یک حرف بزرگ، یک عدد و یک علامت
	var (
		hasUpper = false
		hasLower = false
		hasDigit = false
		
		hasSpec  = false
	)

	for _, r := range pwd {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpec = true
		}
	}

	return len(pwd) >= 8 && hasUpper && hasLower && hasDigit && hasSpec
}
