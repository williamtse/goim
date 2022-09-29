package auth

import (
	"encoding/json"
)

type Auth struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (a *Auth) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func (a *Auth) UnmarshalBinary(data []byte) (err error) {
	return json.Unmarshal(data, a)
}

func CheckToken(tokenStr string) (authUser *Auth, err error) {

}
