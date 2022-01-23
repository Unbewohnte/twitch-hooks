package twitchhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Keys struct {
	ClientID     string
	ClientSecret string
}

// access token response struct
type TokenResponse struct {
	AcessToken string `json:"access_token"`
	ExpiresIn  uint   `json:"expires_in"`
	TokenType  string `json:"token_type"`
}

type RequestOptions struct {
	ApplicationKeys Keys
	AccessToken     TokenResponse
}

// Retrieves access token from Twitch
func GetToken(keys *Keys) (*TokenResponse, error) {
	getTokenUrl := fmt.Sprintf("https://id.twitch.tv/oauth2/token?client_id=%s&client_secret=%s&grant_type=client_credentials",
		keys.ClientID, keys.ClientSecret)

	resp, err := http.Post(getTokenUrl, "", bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, fmt.Errorf("could not make a post request: %s", err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(content, &tokenResp)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal token response: %s", err)
	}

	return &tokenResp, nil
}

// gets data about user from api endpoint
func GetUser(displayname string, options *RequestOptions) (string, error) {
	requestUrl := fmt.Sprintf("https://api.twitch.tv/helix/users?login=%s", displayname)

	httpClient := http.Client{}

	request, err := http.NewRequest("GET", requestUrl, new(bytes.Buffer))
	if err != nil {
		return "", fmt.Errorf("could not create a new twitch request: %s", err)
	}
	defer request.Body.Close()

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", options.AccessToken.AcessToken))
	request.Header.Add("Client-id", options.ApplicationKeys.ClientID)

	response, err := httpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("could not make a request to twitch api: %s", err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Checks if the user streaming right now
func IsLive(displayname string, options *RequestOptions) (bool, error) {
	requestUrl := fmt.Sprintf("https://api.twitch.tv/helix/streams?user_login=%s", displayname)

	httpClient := http.Client{}

	request, err := http.NewRequest("GET", requestUrl, new(bytes.Buffer))
	if err != nil {
		return false, fmt.Errorf("could not create a new twitch request: %s", err)
	}
	defer request.Body.Close()

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", options.AccessToken.AcessToken))
	request.Header.Add("Client-id", options.ApplicationKeys.ClientID)

	response, err := httpClient.Do(request)
	if err != nil {
		return false, fmt.Errorf("could not make a request to twitch api: %s", err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	// check if got an empty response -> offline
	if len(data) <= 28 {
		return false, nil
	}

	return true, nil
}

func GetStream() {

}
