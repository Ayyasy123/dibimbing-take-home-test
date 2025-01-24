package repository

import (
	"time"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
	FindUserByID(id int) (*entity.User, error)
	FindAllUsers() ([]entity.User, error)
	UpdateUser(id int, user *entity.User) error
	DeleteUser(id int) error
	IsEmailExists(email string) (bool, error)
	GetTotalUsers(startDate, endDate time.Time) (int64, error)
	GetUserRoleDistribution(role string, startDate, endDate time.Time) (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindUserByID(id int) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *userRepository) FindAllUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUser(id int, user *entity.User) error {
	result := r.db.Model(&entity.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) DeleteUser(id int) error {
	return r.db.Delete(&entity.User{}, id).Error
}

func (r *userRepository) IsEmailExists(email string) (bool, error) {
	var count int64
	r.db.Model(&entity.User{}).Where("email = ?", email).Count(&count)
	return count > 0, nil
}

func (r *userRepository) GetTotalUsers(startDate, endDate time.Time) (int64, error) {
	var totalUser int64
	query := r.db.Model(&entity.User{})

	// Tambahkan filter tanggal jika startDate atau endDate tidak kosong
	if !startDate.IsZero() {
		query = query.Where("created_at >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("created_at <= ?", endDate)
	}

	err := query.Count(&totalUser).Error
	return totalUser, err
}

func (r *userRepository) GetUserRoleDistribution(role string, startDate, endDate time.Time) (int64, error) {
	var totalUser int64
	query := r.db.Model(&entity.User{}).Where("role = ?", role)

	// Tambahkan filter tanggal jika startDate atau endDate tidak kosong
	if !startDate.IsZero() {
		query = query.Where("created_at >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("created_at <= ?", endDate)
	}

	err := query.Count(&totalUser).Error
	return totalUser, err
}
