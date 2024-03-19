package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Brunin", "b@mail.com", "12345")
	assert.Nil(t, err)
	assert.NotEmpty(t, user)
	assert.Equal(t, "Brunin", user.Name)
	assert.Equal(t, "b@mail.com", user.Email)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("Brunin", "b@mail.com", "12345")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("12345"))
	assert.False(t, user.ValidatePassword("12346"))
	assert.NotEqual(t, "12345", user.Password)
}