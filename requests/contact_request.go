package requests

type ContactRequest struct {
	Name    string `json:"name" binding:"required,max=100"`
	Email   string `json:"email" binding:"required,email,max=100"`
	Phone   string `json:"phone" binding:"required,max=20"`
	Message string `json:"message" binding:"required"`
}
