package service

import (
	"errors"
	"meow-meow/repository"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	userRepo repository.UsersRepository
}

func NewUserService(userRepo repository.UsersRepository) UserService {
	return userService{userRepo: userRepo}
}

func (u userService) CreateUser(userReq UserRequest) (string, error) {

	if len(strings.TrimSpace(userReq.Email)) == 0 || len(strings.TrimSpace(userReq.Password)) == 0 {
		return "", echo.NewHTTPError(http.StatusBadRequest, "email and password is require")
	}

	err := checkEmailIsDuplicate(userReq, u)

	if err != nil {
		return "", err
	}

	password := userReq.Password

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	user := repository.User{
		Email:    userReq.Email,
		Password: string(hashPassword),
	}

	result, err := u.userRepo.CreateUser(user)

	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return result, nil
}

func (u userService) UserLogin(userReq UserRequest) (*TokenResponse, error) {

	if len(strings.TrimSpace(userReq.Email)) == 0 || len(strings.TrimSpace(userReq.Password)) == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "email and password is require")
	}

	userData, err := u.userRepo.GetUserByName(userReq.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	user := repository.User{}

	for _, u := range userData {
		if userReq.Email == u.Email {
			user.Id = u.Id
			user.Email = u.Email
			user.Password = u.Password
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	token, err := createToken(user.Id, user.Email)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	tokenData := TokenResponse{
		Token: token,
	}

	return &tokenData, nil
}

func createToken(id int, username string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	secretKey := []byte(jwtSecret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":  id,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func checkEmailIsDuplicate(userReq UserRequest, u userService) error {
	emailData, err := u.userRepo.GetUserByName(userReq.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	if len(emailData) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Email is exist")
	}

	return nil
}
