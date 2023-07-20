package user

type UserDto struct {
	Id             string `json:"id" form:"id"`
	UserName       string `json:"user_name" form:"user_name"`
	FullName       string `json:"full_name" form:"full_name"`
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	PhoneNumber    string `json:"phone_number" form:"phone_number"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	IsSeller       bool   `json:"is_seller" form:"is_seller"`
	IsAdmin        bool   `json:"is_admin" form:"is_admin"`
}

type UserNameLoginDto struct {
	UserName string `json:"user_name" form:"user_name" valid:"required~Email cannot be empty,email~Please provide a valid email"`
	Password string `json:"password" form:"password" valid:"required~Password cannot be empty, minstringlength(8)~Password cannot be less than 8 characters"`
}
