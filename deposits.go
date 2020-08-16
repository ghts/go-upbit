package upbit

import (
	"context"
	"net/http"
)

func (s *DepositService) ListCoinAddresses(ctx context.Context) ([]*CoinAddress, *http.Response, error) {
	u := "v1/deposits/coin_addresses"
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.generateToken(req, "")
	if err != nil {
		return nil, nil, err
	}

	addrs := []*CoinAddress{}
	resp, err := s.client.Do(ctx, req, &addrs)
	if err != nil {
		return nil, resp, err
	}

	return addrs, nil, nil
}
