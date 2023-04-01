package service

import (
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	SALT       = "m_YnFxOtnL9j_gW)6sV0(IBH"
	LETTERS    = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_-+="
	tokenTTL   = 12 * time.Hour
	singingKey = "=$sRsdvs!P@K$77IGj!F8DJm"
)

type AuthService struct {
	repo repository.IAuthorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.IAuthorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user cargodelivery.User) (int, error) {
	salt, err := generatePasswordSalt(24)
	if err != nil {
		return 0, err
	}
	hash, err := generatePasswordHash(user.Password, salt)
	if err != nil {
		return 0, err
	}
	user.Password = hash
	user.Salt = salt
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUserById(userId int) (cargodelivery.User, error) {
	return s.repo.GetUserById(userId)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+user.Salt+SALT))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(singingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(singingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt+SALT), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func generatePasswordSalt(n int) (string, error) {
	ret := make([]byte, n)
	letLen := int64(len(LETTERS))
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(letLen))
		if err != nil {
			return "", err
		}
		ret[i] = LETTERS[num.Int64()]
	}

	return string(ret), nil
}
