package stack

import (
	"context"
	"github.com/sitehostnz/gosh/pkg/utils"
)

func (s *Client) Stop(ctx context.Context, request *StopStartRequest) (*StartStopResponse, error) {

	u := "cloud/stack/stop.json"
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}
	keys := []string{
		"apikey",
		"client_id",
		"server",
		"name",
	}
	// leave out the containers for now...

	v := req.URL.Query()
	v.Add("server", request.ServerName)
	v.Add("name", request.Name)

	req.URL.RawQuery = utils.Encode(v, keys)

	response := new(StartStopResponse)
	err = s.client.Do(ctx, req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
