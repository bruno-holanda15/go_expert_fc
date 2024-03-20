package database

import "github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/entity"

type UserDBInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}