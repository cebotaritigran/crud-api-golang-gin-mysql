package models

import (
	"crud-api/utils/token"
	"errors"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `gorm:"not null;primaryKey;autoIncrement"`
	FirstName string    `gorm:"not null;type:varchar(115);default:''" json:"firstname"`
	LastName  string    `gorm:"not null;type:varchar(115);default:''" json:"lastname"`
	Email     string    `gorm:"not null;type:varchar(115);default:'';unique" json:"email"`
	Password  string    `gorm:"not null;type:varchar(115);default:''" json:"password"`
	Authlevel uint      `gorm:"not null;type:int(11);default:0" json:"authlevel"`
	IsActive  bool      `gorm:"not null;type:tinyint(4);default:1;" json:"isactive"`
	CreatedAt time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

func GetUserByID(uid uint) (User, error) {

	var user User

	if err := DB.First(&user, uid).Error; err != nil {
		return user, errors.New("User not found!")
	}

	user.PrepareGive()

	return user, nil

}

func (user *User) PrepareGive() {
	user.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {

	var err error

	user := User{}

	err = DB.Model(User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (user *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	//remove spaces in username
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))

	return nil

}
