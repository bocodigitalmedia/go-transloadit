package transloadit_api

import "time"

const ApiAuthExpiresFormat = "2006/01/02 15:04:05+00:00"

type Auth struct {
	Key     string `json:"key"`
	Expires string `json:"expires"`
}

func NewAuth(key, expires string) *Auth {
	return &Auth{
		Key:     key,
		Expires: expires,
	}
}

func NewAuthExpiresIn(key string, t time.Duration) *Auth {
	expiresAt := time.Now().Add(t).UTC()
	return &Auth{
		Key:     key,
		Expires: expiresAt.Format(ApiAuthExpiresFormat),
	}
}
