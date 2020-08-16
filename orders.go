package upbit

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

func (s *OrderService) cancelOrder(ctx context.Context, queryString string) (*Order, *http.Response, error) {
	u := fmt.Sprintf("v1/order?%s", queryString)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.generateToken(req, queryString)
	if err != nil {
		return nil, nil, err
	}

	order := &Order{}
	resp, err := s.client.Do(ctx, req, &order)
	if err != nil {
		return nil, resp, err
	}

	return order, resp, nil
}

func (s *OrderService) CancelOrderByUUID(ctx context.Context, uuid string) (*Order, *http.Response, error) {
	params := url.Values{}
	params.Add("uuid", uuid)
	qs := params.Encode()

	return s.cancelOrder(ctx, qs)
}

func (s *OrderService) CancelOrderByIdentifier(ctx context.Context, identifier string) (*Order, *http.Response, error) {
	params := url.Values{}
	params.Add("identifier", identifier)
	qs := params.Encode()

	return s.cancelOrder(ctx, qs)
}

func (s *OrderService) getOrder(ctx context.Context, queryString string) (*Order, *http.Response, error) {
	u := fmt.Sprintf("v1/order?%s", queryString)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.generateToken(req, queryString)
	if err != nil {
		return nil, nil, err
	}

	order := &Order{}
	resp, err := s.client.Do(ctx, req, &order)
	if err != nil {
		return nil, resp, err
	}

	return order, resp, nil
}

func (s *OrderService) GetOrderByIdentifier(ctx context.Context, identifier string) (*Order, *http.Response, error) {
	params := url.Values{}
	params.Add("identifier", identifier)
	qs := params.Encode()

	return s.getOrder(ctx, qs)
}

func (s *OrderService) GetOrderByUUID(ctx context.Context, uuid string) (*Order, *http.Response, error) {
	params := url.Values{}
	params.Add("uuid", uuid)
	qs := params.Encode()

	return s.getOrder(ctx, qs)
}

func (s *OrderService) ListOrders(ctx context.Context, listOpt *OrderListOptions) ([]*Order, *http.Response, error) {
	qv, err := query.Values(listOpt)
	if err != nil {
		return nil, nil, err
	}

	queryString := qv.Encode()
	u := fmt.Sprintf("v1/orders?%s", queryString)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.generateToken(req, queryString)
	if err != nil {
		return nil, nil, err
	}

	orders := []*Order{}
	resp, err := s.client.Do(ctx, req, &orders)
	if err != nil {
		return nil, resp, err
	}

	return orders, resp, nil
}

func (s *OrderService) Order(ctx context.Context, orderReq *OrderRequest) (*Order, *http.Response, error) {
	qv, err := query.Values(orderReq)
	if err != nil {
		return nil, nil, err
	}
	qs := qv.Encode()

	u := fmt.Sprintf("v1/orders?%v", qs)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.generateToken(req, qs)
	if err != nil {
		return nil, nil, err
	}

	order := &Order{}
	resp, err := s.client.Do(ctx, req, &order)
	if err != nil {
		return nil, resp, err
	}

	return order, resp, nil
}

func (s *OrderService) Chances(ctx context.Context, market string) (*Chance, *http.Response, error) {
	params := url.Values{}
	params.Add("market", market)
	qs := params.Encode()

	u := fmt.Sprintf("v1/orders/chance?%v", qs)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.generateToken(req, qs)
	if err != nil {
		return nil, nil, err
	}

	chance := new(Chance)

	resp, err := s.client.Do(ctx, req, chance)
	if err != nil {
		return nil, resp, err
	}

	return chance, nil, nil
}

