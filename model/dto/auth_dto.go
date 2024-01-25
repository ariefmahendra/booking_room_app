package dto

type AuthRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
