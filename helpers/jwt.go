package helpers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type JWTClaim struct {
	jwt.StandardClaims
	Sub   uuid.UUID `json:"sub"`
	Email string    `json:"email"`
}

func GenerateTokenJWT(authID uuid.UUID, email string) (string, error) {
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

func ValidateToken(signedToken string, ctx echo.Context) (string, error) {
    // Load public key
    publicKeyFile := "./config/public_key.pem"
    publicKeyBytes, err := os.ReadFile(publicKeyFile)
    if err != nil {
        return "", err
    }

    publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
    if err != nil {
        return "", err
    }

    token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return publicKey, nil
    })
    if err != nil {
        return "", err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return "", errors.New("invalid token")
    }

    authId, ok := claims["sub"].(string)
    if !ok {
        return "", errors.New("authId not found in token claims")
    }

    // Set authId in context
    ctx.Set("authId", authId)

    return authId, nil
}
