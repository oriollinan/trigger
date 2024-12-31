package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	errTokenNotValid error = errors.New("token is not valid")
)

func Create(claims map[string]string, secret string) (string, error) {
	tokenClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	for k, v := range claims {
		tokenClaims[k] = v
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		tokenClaims,
	)
	rawToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rawToken, nil
}

func Expiry(rawToken string, secret string) (time.Time, error) {
	token, err := jwt.Parse(rawToken, func(_ *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return time.Time{}, err
	}
	exp, ok := token.Claims.(jwt.MapClaims)["exp"].(float64)
	if !ok {
		return time.Time{}, errTokenNotValid
	}
	return time.Unix(int64(exp), 0), nil
}

func Verify(rawToken string, secret string) error {
	token, err := jwt.Parse(rawToken, func(_ *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errTokenNotValid
	}
	return nil
}

func FromRequest(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("empty auth header")
	}

	if strings.HasPrefix(authHeader, "Bearer ") {
		parts := strings.Split(authHeader, " ")
		if len(parts) < 2 || parts[0] != "Bearer" {
			return "", errors.New("could not find token")
		}
		return parts[1], nil
	}
	return "", errors.New("invalid auth header")
}
