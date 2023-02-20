package db

import (
	"FirstProject/internal/domain/entity"
	"database/sql"
)

type userStorage struct {
	database *sql.DB
}

func NewUserStorage(db *sql.DB) *userStorage {
	return &userStorage{database: db}
}

func (s *userStorage) Create(user *entity.User) (*entity.User, error) {
	err := s.database.QueryRow("INSERT INTO users (email,encrypted_password) VALUES ($1,$2) RETURNING id", user.Email, user.EncryptedPassword).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userStorage) FindByEmail(email string) (*entity.User, error) {
	userEntity := &entity.User{}
	err := s.database.QueryRow("SELECT id,email,encrypted_password FROM users WHERE email = $1", email).Scan(&userEntity.ID, &userEntity.Email, &userEntity.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return userEntity, nil
}
func (s *userStorage) FindById(id int) (*entity.User, error) {
	userEntity := &entity.User{}
	err := s.database.QueryRow("SELECT id,email,encrypted_password FROM users WHERE id = $1", id).Scan(&userEntity.ID, &userEntity.Email, &userEntity.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return userEntity, nil
}
