package service

import (
	"WhyAi/models"
	"WhyAi/pkg/repository"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

type UserService struct {
	repo repository.User
}
type ResetClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}
func (s *UserService) GeneratePasswordResetToken(email, signingKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &ResetClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL / 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		email,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *UserService) ResetPassword(resetModel models.UserReset) error {
	token, err := jwt.ParseWithClaims(resetModel.Token, &ResetClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*ResetClaims)
	fmt.Println(claims)
	if !ok || !token.Valid {
		fmt.Println(token)
		return errors.New("token claims are not of type jwt.MapClaims or token invalid")

	}
	email := claims.Username
	if email == "" {
		return errors.New("empty email")
	}
	return s.repo.ResetPassword(email, generatePasswordHash(resetModel.NewPass))
}

func (s *UserService) ResetPasswordRequest(email models.ResetRequest) error {
	token, err := s.GeneratePasswordResetToken(email.Login, signingKey)
	log.Default().Println("end gen reset token")
	if err != nil {
		return err
	}
	resetLink := fmt.Sprintf("%s/reset/?t=%s", os.Getenv("FRONTEND_URL"), token)
	fmt.Println(resetLink)
	/*
		from := os.Getenv("DB_HOST")
		password := os.Getenv("DB_HOST")
		fmt.Println(resetLink)
		// Информация о получателе
		to := []string{
			email.Login,
		}

		// smtp сервер конфигурация
		smtpHost := os.Getenv("DB_HOST")
		smtpPort := os.Getenv("DB_HOST")

		// Сообщение.
		message := []byte("<h1>Сброс пароля</h1>\n" +
			"<p>Перейдите по ссылке, чтобы сбросить пароль</p>\n" +
			"<a href=\"" + resetLink + "\">Сброс</a>\n" +
			"<p>Если вы не запрашивали сброс, не переходите. Время действия ссылки один час</p>")
		log.Default().Println("mail gen end")
		// Авторизация.
		auth := smtp.PlainAuth("", from, password, smtpHost)

		// Отправка почты.
		err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
		if err != nil {
			return err

		}
		fmt.Println("Почта отправлена!") */
	return nil

}
