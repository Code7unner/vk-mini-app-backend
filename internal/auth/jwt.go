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
	NewCookie(name, value string) *http.Cookie
}

type auth struct {
	SecretKey string
	expTime   time.Time
}

type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func New(secret string) Auth {
	return &auth{SecretKey: secret, expTime: time.Now().Add(200 * time.Hour)}
}

func (a auth) GetUserID(c *http.Cookie) (int, error) {
	userClaims := new(UserClaims)
	_, err := jwt.ParseWithClaims(c.Value, userClaims, func(token *jwt.Token) (interface{}, error) {
		return a.SecretKey, nil
	})
	if err != nil {
		return 0, err
	}

	userID, _ := strconv.Atoi(userClaims.UserID)

	return userID, nil
}

func (a auth) SetUserToken(id int) (string, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &UserClaims{
		UserID: strconv.Itoa(id),
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: a.expTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
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
