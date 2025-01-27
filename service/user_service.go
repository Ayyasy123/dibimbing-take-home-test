package service

import (
	"errors"
	"time"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"github.com/Ayyasy123/dibimbing-take-home-test/repository"
	"github.com/Ayyasy123/dibimbing-take-home-test/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(req *entity.RegisterReq) (*entity.UserRes, error)
	LoginUser(req *entity.LoginReq) (*entity.UserRes, error)
	FindUserByID(id int) (*entity.UserRes, error)
	FindAllUsers() ([]entity.UserRes, error)
	UpdateUser(id int, req *entity.UpdateUserReq) error
	DeleteUser(id int) error
	RegisterAsAdmin(req *entity.RegisterReq) (*entity.UserRes, error)
	GetUserReport(startDate, endDate time.Time) (*entity.UserReport, error)
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

	err = s.userRepository.CreateUser(user)
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

	// Generate token JWT
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	userRes := &entity.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token, // Tambahkan field Token ke UserRes
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

func (s *userService) UpdateUser(id int, req *entity.UpdateUserReq) error {
	existingUser, err := s.userRepository.FindUserByID(id)
	if err != nil {
		return err
	}

	if req.Name != "" {
		existingUser.Name = req.Name
	}

	if req.Email != "" {
		existingUser.Email = req.Email
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		existingUser.Password = string(hashedPassword)
	}

	if req.Role != "" {
		existingUser.Role = req.Role
	}

	return s.userRepository.UpdateUser(id, existingUser)

}

func (s *userService) DeleteUser(id int) error {
	return s.userRepository.DeleteUser(id)
}

func (s *userService) RegisterAsAdmin(req *entity.RegisterReq) (*entity.UserRes, error) {
	exist, err := s.userRepository.IsEmailExists(req.Email)
	if err != nil {
		return nil, err
	}

	if exist {
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
		Role:     "admin",
	}

	err = s.userRepository.CreateUser(user)
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

func (s *userService) GetUserReport(startDate, endDate time.Time) (*entity.UserReport, error) {
	// Hitung total user
	totalUser, err := s.userRepository.GetTotalUsers(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Daftar role user yang ingin dihitung
	roles := []string{"admin", "user"}

	// Slice untuk menyimpan distribusi role user
	var roleDistribution []entity.UserRoleDistribution

	// Loop melalui setiap role dan hitung distribusinya
	for _, role := range roles {
		totalRoleUser, err := s.userRepository.GetUserRoleDistribution(role, startDate, endDate)
		if err != nil {
			return nil, err
		}
		roleDistribution = append(roleDistribution, entity.UserRoleDistribution{
			Role:      role,
			TotalUser: int(totalRoleUser),
		})
	}

	return &entity.UserReport{
		TotalUser:            int(totalUser),
		UserRoleDistribution: roleDistribution,
	}, nil
}
