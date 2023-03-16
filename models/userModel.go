package models

import (
	"crud-api/utils/token"
	"errors"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// user struct to define table columns like id, name and email
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

// takes user id as parameter and returns User struct and an error
func GetUserByID(uid uint) (User, error) {

	var user User

	if err := DB.First(&user, uid).Error; err != nil {
		return user, errors.New("User not found!")
	}

	user.PrepareGive()
	// because I need to return an error and an user struct If the if statement above
	//doesnt find an error we will return  nil (null)
	return user, nil

}

// when serving user to frontend we can't pass the password as it is, so here
// I just set it a empty string
// Struct Method that takes user struct and calls preparegive method on it
// a function that acts on user object
func (user *User) PrepareGive() {
	user.Password = ""
}

// this functions takes 2 parameters(strings) and returns an error if none it will return nil
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// it will take email and password as parameters and will return a string and an error
// the returned string will be jwt token and error if all went well will be nil
func LoginCheck(email string, password string) (string, error) {

	var err error
	// assigning user to our USER model so that we can pass information
	user := User{}

	//Take return a record that match given conditions, the order will depend on the database implementation
	err = DB.Model(User{}).Where("email = ?", email).Take(&user).Error

	// checking if there were errors assigning information to user
	if err != nil {
		return "", err
	}

	// verifying password
	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	//generating token with our user.ID
	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	//returning token as string and nil as error
	return token, nil
}

// method to register user in database which we will call in our
// controller function
func (user *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// method to hash password before save
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
