package host

import "encoding/json"

type UserConnectedHost struct {
	IP   string `json:"ip" redis:"ip"`
	Port int32  `json:"port" redis:"port"`
}

func (a *UserConnectedHost) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func (a *UserConnectedHost) UnmarshalBinary(data []byte) (err error) {
	return json.Unmarshal(data, a)
}
func (a *UserConnectedHost) BinaryUnmarshaler(data []byte) error {
	return json.Unmarshal(data, a)
}
