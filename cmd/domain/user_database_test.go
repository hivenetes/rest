package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "user not found")
	assert.NotNil(t, err)
	assert.EqualValues(t, "user 0 not found", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
}

func TestGetUserFound(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Ab", user.FirstName)
	assert.EqualValues(t, "S", user.LastName)
	assert.EqualValues(t, "abs@mail.com", user.Email)
}

func BenchmarkGetUserFound(b *testing.B){
	for i := 0; i < b.N; i++ {
		GetUser(123)
	}
}