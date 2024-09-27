package service

type ProfileResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile"`
	Sex       string `json:"sex"`
	Status    string `json:"status"`
	Image     string `json:"image"`
	UserId    int    `json:"user_id"`
}

type ProfileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile"`
	Sex       string `json:"sex"`
}

type ProfileService interface {
	GetProfileByUserId(int) (*ProfileResponse, error)
	CreateUserProfile(ProfileRequest, int, bool) error
}
