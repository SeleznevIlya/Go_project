package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/SeleznevIlya/Go_project"
	"github.com/SeleznevIlya/Go_project/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	repo repository.Autorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user Go_project.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		fmt.Println("metropolit")
		return "", err
	}

	fmt.Println(user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.Id,
		},
	)

	return token.SignedString([]byte(os.Getenv("signingKey")))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("salt"))))
}
