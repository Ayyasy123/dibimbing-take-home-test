package entity

import "time"

type User struct {
	ID        int       `gorm:"primary_key,auto_increment" json:"id"`
	Name      string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tickets   []Ticket  `json:"tickets" gorm:"foreignKey:UserID"`
}

type RegisterReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UserRes struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"token,omitempty"`
}

type UserRoleDistribution struct {
	Role      string `json:"role"`       // Role user (admin, user)
	TotalUser int    `json:"total_user"` // Total user dengan role tersebut
}

type UserReport struct {
	TotalUser            int                    `json:"total_user"`             // Total user yang terdaftar
	UserRoleDistribution []UserRoleDistribution `json:"user_role_distribution"` // Distribusi role user
}
