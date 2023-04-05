package openpeoplesearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResponse struct {
	Username       string    `json:"username"`
	Token          string    `json:"token"`
	Message        string    `json:"message"`
	TokenExpiryUtc time.Time `json:"token_expiry_utc"`
}

func NewSession(username, password string) (Session, error) {
	authReq := User{username, password}
	requestBody, err := json.Marshal(authReq)
	if err != nil {
		return Session{}, err
	}

	req, err := http.NewRequest("POST", "https://api.openpeoplesearch.com/api/v1/User/authenticate", bytes.NewBuffer(requestBody))
	if err != nil {
		return Session{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "*/*")
	req.Header.Set("User-Agent", USERAGENT)

	resp, err := opsClient.Do(req)
	if err != nil {
		return Session{}, err
	}

	defer resp.Body.Close()
	var authResp authResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return Session{}, err
	}

	if authResp.Message != "OK" {
		return Session{}, errors.New(authResp.Message)
	}

	session := Session{
		Username:    authResp.Username,
		Token:       authResp.Token,
		TokenExpiry: authResp.TokenExpiryUtc,
	}

	return session, nil
}
