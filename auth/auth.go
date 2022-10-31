package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

// Bearer tokens
type Bearer_auth struct {
	Token string
}

func (a Bearer_auth) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

// Access Token
type Access_auth struct {
	Access_Token  string
	Access_Secret string
}

// Client
type Client_auth struct {
	Client_Token  string
	Client_Secret string
}

func (a Client_auth) Add(req *http.Request) {
	token := a.Client_Token + ":" + a.Client_Secret
	token = base64.StdEncoding.EncodeToString([]byte(token))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
}
