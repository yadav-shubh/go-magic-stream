package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yadav-shubh/go-magic-stream/config"
	"github.com/yadav-shubh/go-magic-stream/models"
	"go.uber.org/zap"
)

type KindeUtility struct {
}

func NewKindeUtility() *KindeUtility {
	return &KindeUtility{}
}

func (k *KindeUtility) VerifyAuthCode(code string) (*models.AuthResponse, error) {
	ssoConfig := config.Get().SSO

	endpoint := ssoConfig.Issuer + ssoConfig.TokenUrl

	data := url.Values{}
	data.Add("client_id", ssoConfig.ClientId)
	data.Add("client_secret", ssoConfig.ClientSecret)
	data.Add("grant_type", ssoConfig.GrantType)
	data.Add("redirect_uri", ssoConfig.RedirectUri)
	data.Add("code", code)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Log.Error("failed to close response body", zap.Error(err))
		}
	}(resp.Body)

	var apiResponse *models.AuthResponse

	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &apiResponse)
	if err != nil || apiResponse.AccessToken == "" {
		Log.Error("invalid or expired code")
		return nil, fmt.Errorf("invalid or expired code")
	}

	return apiResponse, nil
}

func (k *KindeUtility) GetUserEmail(accessToken string) string {
	token, _, err := jwt.NewParser().ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return ""
	}
	return token.Claims.(jwt.MapClaims)["email"].(string)
}
