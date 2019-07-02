package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2/jwt"
)

var scopes = []string{"https://www.googleapis.com/auth/spreadsheets"}

type serviceAccountFile struct {
	Type        string `json:"type"`
	TokenURI    string `json:"token_uri"`
	PrivateKey  string `json:"private_key"`
	ClientEmail string `json:"client_email"`
}

func NewFromServiceAccountFile(path string) (*http.Client, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	sa := serviceAccountFile{}
	err = json.Unmarshal(b, &sa)
	if err != nil {
		return nil, err
	}

	if sa.Type != "service_account" {
		return nil, fmt.Errorf("not service_account type")
	}

	config := &jwt.Config{
		Email:      sa.ClientEmail,
		PrivateKey: []byte(sa.PrivateKey),
		TokenURL:   sa.TokenURI,
		Scopes:     scopes,
	}

	client := config.Client(context.Background())
	return client, nil
}
