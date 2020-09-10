package upbit

import (
	"context"
	"fmt"
	"net/http"
)

var ErrCurrencyNotFound = fmt.Errorf("upbit: currency not found")

func (s *AccountService) AccountCurrency(ctx context.Context, currency string) (*Account, *http.Response, error) {
	accounts, resp, err := s.Accounts(ctx)
	if err != nil {
		return nil, resp, err
	}

	for _, account := range accounts {
		if account.Currency == currency {
			return account, resp, nil
		}
	}

	return nil, resp, ErrCurrencyNotFound
}

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
