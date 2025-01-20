package service

import (
	"errors"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"github.com/Ayyasy123/dibimbing-take-home-test/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) RegisterUser(req *entity.RegisterReq) (*entity.UserRes, error) {
	// cek email apakah sudah ada di database
	exists, err := s.userRepository.IsEmailExists(req.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("email already registered")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	err = s.userRepository.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	userRes := &entity.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userRes, nil
}

func (s *userService) LoginUser(req *entity.LoginReq) (*entity.UserRes, error) {
	user, err := s.userRepository.FindUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	userRes := &entity.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userRes, nil
}

func (s *userService) FindUserByID(id int) (*entity.UserRes, error) {
	user, err := s.userRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	userRes := &entity.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userRes, nil
}

func (s *userService) FindAllUsers() ([]entity.UserRes, error) {
	users, err := s.userRepository.FindAllUsers()
	if err != nil {
		return nil, err
	}

	var userRes []entity.UserRes
	for _, user := range users {
		userRes = append(userRes, entity.UserRes{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return userRes, nil
}

func (s *userService) UpdateUser(user *entity.User) error {
	return s.userRepository.UpdateUser(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.userRepository.DeleteUser(id)
}
