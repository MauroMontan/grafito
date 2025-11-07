package auth

import "net/http"

type Authenticator interface {
	Apply(req *http.Request)
}

type BearerAuth struct {
	Token string
}

func (a BearerAuth) Apply(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+a.Token)
}
