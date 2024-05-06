package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	jwt.StandardClaims
	Sub   string `json:"sub"`
	Email string `json:"email"`
}

func GenerateTokenJWT(authID string, email string) (string, error) {
	// Load private key
	privateKeyFile := "./config/private_key.pem"
	privateKeyBytes, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", err
	}
	currentTime := time.Now()
	expirationTime := currentTime.Add(12 * time.Hour)

	claims := &JWTClaim{
		Sub:   authID,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			Issuer:    "https://tanam.id",
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	TokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return TokenString, err
}
