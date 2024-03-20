package database

import (
	"testing"

	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Brunin", "b@mail.com", "12345")
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
