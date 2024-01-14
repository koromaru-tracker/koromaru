package types

type RegisterRequest struct {
	Username        string `form:"username" json:"username" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required"`
	Email           string `form:"email" json:"email" binding:"required"`
}
