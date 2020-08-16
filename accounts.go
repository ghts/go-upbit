package upbit

import (
	"context"
	"fmt"
	"net/http"
)

func (s *AccountService) Accounts(ctx context.Context) ([]*Account, *http.Response, error) {
	u := fmt.Sprintf("v1/accounts")
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.generateToken(req, "")
	if err != nil {
		return nil, nil, err
	}

	var accounts []*Account
	resp, err := s.client.Do(ctx, req, &accounts)
	if err != nil {
		return nil, resp, err
	}

	return accounts, nil, nil
}
