package service

import (
	"FirstProject/internal/domain/entity"
	"FirstProject/pkg/logging"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage interface {
	Create(user *entity.User) (*entity.User, error)
}

type userService struct {
	storage UserStorage
}

func NewUserService(storage UserStorage) *userService {
	return &userService{storage: storage}
}

func (s *userService) CreateUser(user *CreateUserDTO) (*entity.User, error) {
	logger := logging.GetLogger()
	logger.Info("Create user in service")
	logger.Info(user)
	var err error
	//TODO validate parameters
	err = validate(user)
	if err != nil {
		return nil, err
	}
	//TODO encrypt password and DTO--->entity
	encryptedPassword, err := encryptString(user.Password)
	if err != nil {
		return nil, err
	}
	userEntity := &entity.User{
		EncryptedPassword: encryptedPassword,
		Email:             user.Email,
	}

	return s.storage.Create(userEntity)
}

func validate(user *CreateUserDTO) error {
	return validation.ValidateStruct(user, validation.Field(&user.Email, validation.Required, is.Email), validation.Field(&user.Password, validation.Required, validation.Length(6, 20)))
}

func encryptString(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), nil
}
