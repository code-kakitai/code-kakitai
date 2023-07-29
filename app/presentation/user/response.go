package user

type getUserResponse struct {
	User userResponseModel `json:"users"`
}

type userResponseModel struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	Address     string `json:"address"`
}
