package controllers

import "github.com/poseidon/app/models"

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (in *UserLoginInput) IsValid() bool {
	return in.Email != "" && in.Password != ""
}

type UserLoginOutput struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func NewUserLoginOutput(user *models.UserModel) *UserLoginOutput {
	return &UserLoginOutput{
		ID:    user.ID.Hex(),
		Email: user.Email,
	}
}
