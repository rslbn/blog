package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	db "github.com/rslbn/blog/postgres"
)

type JwtService interface {
	GenerateToken(u *db.User) (string, error)
	ValidateToken(t string) error
	GetUsernameFromToken(t string) (string, error)
}

type jwtService struct {
}

var signingKey []byte = secretKey()

func NewJwtService() JwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(u *db.User) (string, error) {
	var (
		key    []byte = signingKey
		t      *jwt.Token
		signed string
	)
	issAt := time.Now()
	expAt := issAt.Add(24 * time.Hour)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    "this-blog",
			Subject:   u.Username,
			ExpiresAt: jwt.NewNumericDate(expAt),
			IssuedAt:  jwt.NewNumericDate(issAt),
		},
	)
	signed, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	return signed, nil
}

func (s *jwtService) ValidateToken(t string) error {
	token, err := s.parseToken(t)
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func (s *jwtService) parseToken(t string) (*jwt.Token, error) {
	parser := jwt.NewParser()
	token, err := parser.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *jwtService) GetUsernameFromToken(t string) (string, error) {

	token, err := s.parseToken(t)
	if err != nil {
		return "", err
	}
	sub, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}
	return sub, nil
}

func secretKey() []byte {
	return []byte(os.Getenv("BLOG_SECRET"))
}
