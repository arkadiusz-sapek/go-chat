package auth

import (
	"fmt"
	"time"

	"github.com/ArkadiuszSa/go-chat/config"
	"github.com/dgrijalva/jwt-go"
)

//JWTService - interface for jwt service
type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	ID string `json:"id"`
	ExpirationDate int64 `json:expirationDate`
	IssueDate int64 `json:"issueDate"`
	jwt.StandardClaims
}


 func GenerateToken(email string, userId string) string {
	claims := &authCustomClaims{	
		userId,
		time.Now().Add(time.Hour * 48).Unix(),
		time.Now().Unix(),	
		jwt.StandardClaims{
			Id: userId,
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(config.SecretKey), nil
	})

}

