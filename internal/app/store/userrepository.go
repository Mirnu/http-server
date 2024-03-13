package store

import "github.com/gopherschool/http-rest-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users(email, encrypted_password) VALUES($1, $2) RETURNING id, email, encrypted_password",
		user.Email, user.EncryptedPassword).
		Scan(&user.ID); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
