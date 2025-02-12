package model

type UserPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductCreateResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
