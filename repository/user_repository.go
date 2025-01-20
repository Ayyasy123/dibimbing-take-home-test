package repository

import (
	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
	FindUserByID(id int) (*entity.User, error)
	FindAllUsers() ([]entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
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

func (r *userRepository) UpdateUser(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(user *entity.User) error {
	return r.db.Delete(user).Error
}

func (r *userRepository) IsEmailExists(email string) bool {
	var count int64
	r.db.Model(&entity.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
