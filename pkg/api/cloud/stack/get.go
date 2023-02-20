package stack

import (
	"context"

	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/sitehostnz/gosh/pkg/utils"
)

// Get fetches a cloud stack.
func (s *Client) Get(ctx context.Context, request GetRequest) (*models.Stack, error) {
	req, err := s.client.NewRequest("GET", "cloud/stack/get.json", "")
	if err != nil {
		return nil, err
	}
	keys := []string{
		"apikey",
		"client_id",
		"server",
		"name",
	}

	v := req.URL.Query()
	v.Add("server", request.ServerName)
	v.Add("name", request.Name)

	req.URL.RawQuery = utils.Encode(v, keys)

	response := new(GetResponse)
	err = s.client.Do(ctx, req, response)
	if err != nil {
		return nil, err
	}

	return response.Stack, nil
}
