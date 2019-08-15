package models

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	u "github.com/d3z41k/rest-boilerplate/utils"
)

// Token is JWT claims struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// Account is a struct to rep user account
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

// Validate incoming user details
func (a *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(a.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(a.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	tempAccount := &Account{}

	// Check for errors and duplicate emails
	err := GetDB().Table("accounts").Where("email = ?", a.Email).First(tempAccount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if tempAccount.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

// Create new user with token
func (a *Account) Create() map[string]interface{} {
	if resp, ok := a.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	a.Password = string(hashedPassword)

	GetDB().Create(a)

	if a.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	// Create new JWT token for the newly registered account
	tk := &Token{UserID: a.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
	a.Token = tokenString

	a.Password = ""

	response := u.Message(true, "Account has been created")
	response["account"] = a
	return response
}

// Login - user authorization
func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	account.Password = ""

	// Create JWT token
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

// GetUser return a user by ID
func GetUser(id uint) *Account {
	account := &Account{}

	GetDB().Table("accounts").Where("id = ?", id).First(account)
	if account.Email == "" {
		return nil
	}

	account.Password = ""
	return account
}
