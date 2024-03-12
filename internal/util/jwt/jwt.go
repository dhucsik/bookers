package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/samber/lo"
)

const (
	refreshClaim = "refresh"
	expClaim     = "exp"
)

var internalClaims = []string{refreshClaim, expClaim}

func GenerateJWT(session *models.Session, ttl time.Duration, isRefresh bool) (string, error) {
	claims := make(map[string]any)

	claims["user_id"] = session.UserID
	claims["role"] = session.Role
	claims["exp"] = time.Now().Add(ttl).Unix()
	claims["refresh"] = isRefresh

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims(claims))
	out, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return out, nil
}

func ParseJWT(tokenString string) (*models.Session, bool, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})
	if err != nil {
		return nil, false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, false, errors.ErrInvalidJWTToken
	}

	return getSession(claims)
}

func getSession(claims jwt.MapClaims) (*models.Session, bool, error) {
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return nil, false, errors.ErrInvalidJWTToken
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, false, errors.ErrInvalidJWTToken
	}

	_, ok = claims["exp"].(float64)
	fmt.Println(ok)
	if !ok {
		return nil, false, errors.ErrInvalidJWTToken
	}

	isRefresh, ok := claims["refresh"].(bool)
	if !ok {
		return nil, false, errors.ErrInvalidJWTToken
	}

	sesClaims := make(map[string]interface{})
	for k, v := range claims {
		if lo.Contains(internalClaims, k) {
			continue
		}

		sesClaims[k] = v
	}

	return &models.Session{
		UserID: int(userID),
		Role:   role,
		Claims: sesClaims,
	}, isRefresh, nil
}
