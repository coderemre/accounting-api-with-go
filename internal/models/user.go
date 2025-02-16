package models

import (
	"accounting-api-with-go/internal/utils"
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_secret_key")

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username,omitempty"`
	Email        string `json:"email"`
	Password string `json:"password,omitempty"`
	Role         string `json:"role,omitempty"`
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func (u *User) Validate(forLogin bool) utils.Message {
	if forLogin {
		if u.Email == "" {
			return utils.ErrEmailRequired
		}
		if !isValidEmail(u.Email) {
			return utils.ErrInvalidEmailFormat
		}
		if len(u.Password) < 6 {
			return utils.ErrPasswordTooShort
		}
		return ""
	}

	if u.Username == "" {
		return utils.ErrUsernameRequired
	}
	if u.Email == "" {
		return utils.ErrEmailRequired
	}
	if !isValidEmail(u.Email) {
		return utils.ErrInvalidEmailFormat
	}
	if len(u.Password) < 6 {
		return utils.ErrPasswordTooShort
	}
	return ""
}

func (u *User) GenerateToken() (string, utils.Message) {
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"email":   u.Email,
		"role":    u.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", utils.ErrTokenGeneration
	}
	return tokenString, ""
}

func (u *User) ValidateToken(tokenString string) (*User, utils.Message) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(string(utils.ErrInvalidToken))
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, utils.ErrInvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return u.FromTokenClaims(claims)
	}
	return nil, utils.ErrInvalidToken
}

func (u *User) FromTokenClaims(claims jwt.MapClaims) (*User, utils.Message) {
	user := &User{}
	if id, ok := claims["user_id"].(float64); ok {
		user.ID = int64(id)
	} else {
		return nil, utils.ErrInvalidToken
	}

	if email, ok := claims["email"].(string); ok {
		user.Email = email
	} else {
		return nil, utils.ErrInvalidToken
	}

	if role, ok := claims["role"].(string); ok {
		user.Role = role
	} else {
		return nil, utils.ErrInvalidToken
	}

	return user, ""
}