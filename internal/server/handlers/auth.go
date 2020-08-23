package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
	"webapp1/internal/server/services/auth"
)

const (
	JwtKey               = "super_secret_jwt_key"
	AccessTokenDuration  = 15 * time.Minute
	RefreshTokenDuration = 24 * time.Hour
)

func PostAuthToken(c echo.Context) error {
	type authRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	req := authRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	user := auth.AuthenticateUser(req.Username, req.Password)
	if user == nil {
		return echo.ErrUnauthorized
	}

	accessToken, refreshToken, err := generateTokens(*user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  *accessToken,
		"refresh_token": *refreshToken,
	})
}

func PostAuthTokenRefresh(c echo.Context) error {
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	tokenReq := tokenReqBody{}
	if err := c.Bind(&tokenReq); err != nil {
		return err
	}

	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return err
	}
	if token == nil || !token.Valid {
		return echo.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.ErrUnauthorized
	}

	userId := claims["userid"].(int)
	user := auth.GetUser(userId)
	if user == nil {
		return echo.ErrUnauthorized
	}

	accessToken, refreshToken, err := generateTokens(*user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  *accessToken,
		"refresh_token": *refreshToken,
	})
}

func generateTokens(user auth.User) (*string, *string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	accessTokenClaims["exp"] = time.Now().Add(AccessTokenDuration).Unix()
	accessTokenClaims["userid"] = user.ID
	accessTokenClaims["username"] = user.Username
	accessTokenClaims["role"] = user.Role
	accessTokenClaims["department"] = user.Department
	accessTokenString, err := accessToken.SignedString([]byte(JwtKey))
	if err != nil {
		return nil, nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["userid"] = user.ID
	refreshTokenClaims["exp"] = time.Now().Add(RefreshTokenDuration).Unix()
	refreshTokenString, err := refreshToken.SignedString([]byte(JwtKey))
	if err != nil {
		return nil, nil, err
	}

	return &accessTokenString, &refreshTokenString, nil
}
