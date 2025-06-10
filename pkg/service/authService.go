package service

import (
	"WhyAi/models"
	"WhyAi/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

const (
	salt       = "gdfgdf789fsd798ghdfh9d8f79d8fs"                            //абфускатор пароля "соль"
	signingKey = "js786b87^*bn98v79&(*jhkjhKj6kiu6iU^^u6iU^uk6tiuufv6biu^u6" //ключ подписи
	tokenTTL   = time.Hour * 12                                              //время действия токена
)

func (s *AuthService) GenerateToken(login models.AuthUser) (string, error) {
	//get user from db

	user, err := s.repo.GetUser(login.Email, generatePasswordHash(login.Password), true)

	if err != nil {

		// ("Failed Get User")
		return "fail", err
	}
	// ("auth.go: ", user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Name,
	})
	// ("EndGenToken")
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return models.User{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return models.User{}, errors.New("token claims are not of type *tokenClaims")
	}

	returnUser := models.User{
		Id:   claims.UserId,
		Name: claims.Username,
	}

	return returnUser, nil
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.SignUp(user)
}

// TODO Дописать получение
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
