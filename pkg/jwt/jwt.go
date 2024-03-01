package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenereateToken(claims *UserClaims) (*JWTRes, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Issuer:    "nagabelajar",
	}

	accesToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accesTokenString, err := accesToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	res := &JWTRes{
		AccessToken: accesTokenString,
	}

	return res, nil
}

func ValidateJWT(token string) (*UserClaims, error) {
	claims := &UserClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
