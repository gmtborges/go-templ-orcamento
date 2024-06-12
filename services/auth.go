package services

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/argon2"

	"github.com/gustavomtborges/orcamento-auto/models"
	"github.com/gustavomtborges/orcamento-auto/repositories"
)

type argon2Params struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

type AuthService struct {
	argon2Params
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthService {
	params := &argon2Params{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}

	return &AuthService{argon2Params: *params, userRepo: userRepo}
}

func (s *AuthService) GetRolesByUserID(userID int) ([]string, error) {
	return []string{"admin", "coop_admin"}, nil
}

func (s *AuthService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) GeneratePasswordHash(password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		s.argon2Params.time,
		s.argon2Params.memory,
		s.argon2Params.threads,
		s.argon2Params.keyLen,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	finalHash := fmt.Sprintf("%s:%s", b64Salt, b64Hash)
	return finalHash, nil
}

func (s *AuthService) VerifyPasswordHash(password, hash string) (bool, error) {
	parts := strings.Split(hash, ":")
	if len(parts) != 2 {
		return false, errors.New("invalid hash format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	hashToVerify := argon2.IDKey(
		[]byte(password),
		salt,
		s.argon2Params.time,
		s.argon2Params.memory,
		s.argon2Params.threads,
		s.argon2Params.keyLen,
	)

	if subtle.ConstantTimeCompare(decodedHash, hashToVerify) == 1 {
		return true, nil
	}

	return false, nil
}

func (s *AuthService) SetAuthSession(c echo.Context, userID int) error {
	session, err := session.Get("auth-session", c)
	if err != nil {
		return err
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	session.Values["authenticated"] = true
	session.Values["user_id"] = userID
	return session.Save(c.Request(), c.Response())
}

func (s *AuthService) IsAuthenticated(c echo.Context) (bool, error) {
	session, err := session.Get("auth-session", c)
	if err != nil {
		return false, err
	}
	auth, ok := session.Values["authenticated"].(bool)
	return auth && ok, nil
}

func (s *AuthService) RemoveAuthSession(c echo.Context) error {
	session, err := session.Get("auth-session", c)
	if err != nil {
		return err
	}
	session.Values["authenticated"] = false
	session.Values["user_id"] = nil
	session.Options.MaxAge = -1
	return session.Save(c.Request(), c.Response())
}
