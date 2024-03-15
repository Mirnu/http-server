package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "5H8yN@example.com",
		password: "password123",
	}
}
