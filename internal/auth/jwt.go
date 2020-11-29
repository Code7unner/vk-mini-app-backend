package auth

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"time"
)

type Auth interface {
	SetUserToken(id int) (string, error)
	GetUserID(cookie *http.Cookie) (int, error)
	GetUserIDFromHeader(token string) string

	NewCookie(name, value string) *http.Cookie
}

type auth struct {
	SecretKey string
	expTime   time.Time
}

func New(secret string) Auth {
	return &auth{SecretKey: secret, expTime: time.Now().Add(200 * time.Hour)}
}

func (a auth) GetUserIDFromHeader(token string) string {
	userClaims := new(UserClaims)
	_, err := jwt.ParseWithClaims(token, userClaims, func(token *jwt.Token) (interface{}, error) {
		return a.SecretKey, nil
	})
	if err != nil {
		return ""
	}

	return userClaims.UserID
}

func (a auth) GetUserID(c *http.Cookie) (int, error) {
	userClaims := new(UserClaims)
	_, err := jwt.ParseWithClaims(c.Value, userClaims, func(token *jwt.Token) (interface{}, error) {
		return a.SecretKey, nil
	})
	if err != nil {
		return 0, err
	}

	id, _ := strconv.Atoi(userClaims.UserID)

	return id, nil
}

func (a auth) SetUserToken(id int) (string, error) {
	claims := &UserClaims{
		UserID: strconv.Itoa(id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: a.expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a auth) NewCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: a.expTime,
	}
}
