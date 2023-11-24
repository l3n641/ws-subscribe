package event

import (
	"encoding/json"
	"github.com/olahol/melody"
)

type ChannelReq struct {
	Channels []string
}

func SetChannel(s *melody.Session, data []byte) (interface{}, error) {
	var channels ChannelReq

	err := json.Unmarshal(data, &channels)
	if err != nil {
		return false, err
	}
	s.Set("channels", channels.Channels)

	return true, nil

}
