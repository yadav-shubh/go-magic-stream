package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yadav-shubh/go-magic-stream/config"
	"github.com/yadav-shubh/go-magic-stream/models"
)

const (
	UserIDKey = "user_id"
	EmailKey  = "email"
	ClaimsKey = "claims"
	RoleKey   = "role"
)

var jwtSecret string

func init() {
	jwtSecret = config.Get().JWT.Secret
}

// Token types
const (
	TokenTypeAccess    = "access"
	TokenTypeRefresh   = "refresh"
	TokenTypeBearer    = "Bearer"
	AccessTokenExpiry  = 24 * time.Hour
	RefreshTokenExpiry = 7 * 24 * time.Hour
)

type Claims struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

// GenerateTokenPair creates both access and refresh tokens
func GenerateTokenPair(userID, email, role string) (*models.AuthResponse, error) {
	accessToken, accessExpiry, err := generateToken(userID, email, role, TokenTypeAccess, AccessTokenExpiry)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpiry, err := generateToken(userID, email, role, TokenTypeRefresh, RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessExpiry,
		RefreshTokenExpiresAt: refreshExpiry,
		TokenType:             TokenTypeBearer,
	}, nil
}

// GenerateToken creates a new JWT token for a user (backward compatibility)
func GenerateToken(userID, email string, role string) (string, int64, error) {
	return generateToken(userID, email, role, TokenTypeAccess, 24*time.Hour)
}

// generateToken is the internal function to create tokens
func generateToken(userID, email, role, tokenType string, duration time.Duration) (string, int64, error) {
	expiry := jwt.NewNumericDate(time.Now().Add(duration))
	claims := &Claims{
		UserID:    userID,
		Email:     email,
		Role:      role,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "magic-stream",
			ExpiresAt: expiry,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	if jwtSecret == "" {
		return "", 0, errors.New("jwt secret is missing")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", 0, err
	}
	return signedString, expiry.UnixMilli(), nil
}

// ValidateToken parses and validates a JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ValidateAccessToken validates and ensures it's an access token
func ValidateAccessToken(tokenString string) (*Claims, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != TokenTypeAccess {
		return nil, errors.New("invalid token type: expected access token")
	}

	return claims, nil
}

// ValidateRefreshToken validates and ensures it's a refresh token
func ValidateRefreshToken(tokenString string) (*Claims, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != TokenTypeRefresh {
		return nil, errors.New("invalid token type: expected refresh token")
	}

	return claims, nil
}

// RefreshAccessToken generates a new access token using a valid refresh token
func RefreshAccessToken(refreshTokenString string) (*models.AuthResponse, error) {
	// Validate the refresh token
	claims, err := ValidateRefreshToken(refreshTokenString)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// If so, issue a new refresh token as well
	timeUntilExpiry := time.Until(claims.ExpiresAt.Time)

	if timeUntilExpiry < RefreshTokenExpiry {
		// Issue both new access and refresh tokens
		return GenerateTokenPair(claims.UserID, claims.Email, claims.Role)
	}

	// Only issue new access token
	accessToken, accessExpiry, err := generateToken(
		claims.UserID,
		claims.Email,
		claims.Role,
		TokenTypeAccess,
		AccessTokenExpiry,
	)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshTokenString,
		AccessTokenExpiresAt:  accessExpiry,
		RefreshTokenExpiresAt: claims.ExpiresAt.UnixMilli(),
		TokenType:             TokenTypeBearer,
	}, nil
}

// RefreshToken generates a new token with extended expiration (backward compatibility)
func RefreshToken(tokenString string) (string, int64, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", 0, err
	}

	// Generate new token with same claims but new expiration
	return GenerateToken(claims.UserID, claims.Email, claims.Role)
}

// IsTokenExpired checks if a token is expired without full validation
func IsTokenExpired(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return true
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims.ExpiresAt.Before(time.Now())
	}

	return true
}

// GetTokenExpiryTime returns the expiry time of a token
func GetTokenExpiryTime(tokenString string) (time.Time, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return time.Time{}, err
	}

	return claims.ExpiresAt.Time, nil
}

// GetUserIdFromRefreshToken returns the user id from the refresh token
func GetUserIdFromRefreshToken(refreshTokenString string) (string, error) {
	claims, err := ValidateRefreshToken(refreshTokenString)
	if err != nil {
		return "", err
	}
	return claims.UserID, nil
}
