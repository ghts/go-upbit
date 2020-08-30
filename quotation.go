package upbit

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

func (s *MarketService) All(ctx context.Context) ([]*MarketCode, *http.Response, error) {
	u := fmt.Sprintf("v1/market/all?isDetail=true")
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var markets []*MarketCode

	resp, err := s.client.Do(ctx, req, &markets)
	if err != nil {
		return nil, resp, err
	}

	return markets, nil, nil
}

type CandleListOptions struct {
	To                  string `url:"to,omitempty"`
	Count               int    `url:"count,omitempty"`
	ConvertingPriceUnit string `url:"convertingPriceUnit,omitempty"`
}

func (s *CandleService) candle(ctx context.Context, path string, market string, opts *CandleListOptions) ([]*Candle, *http.Response, error) {
	qv, err := query.Values(opts)
	if err != nil {
		return nil, nil, err
	}

	qv.Add("market", market)

	u := fmt.Sprintf("v1/candles/%s?%s", path, qv.Encode())
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var candles []*Candle

	resp, err := s.client.Do(ctx, req, &candles)
	if err != nil {
		return nil, resp, err
	}

	return candles, nil, nil
}

func (s *CandleService) CandleMinutes(ctx context.Context, market string, unit int, opts *CandleListOptions) ([]*Candle, *http.Response, error) {
	return s.candle(ctx, fmt.Sprintf("minutes/%d", unit), market, opts)
}

func (s *CandleService) CandleDays(ctx context.Context, market string, opts *CandleListOptions) ([]*Candle, *http.Response, error) {
	return s.candle(ctx, "days", market, opts)
}

func (s *CandleService) CandleWeeks(ctx context.Context, market string, opts *CandleListOptions) ([]*Candle, *http.Response, error) {
	return s.candle(ctx, "weeks", market, opts)
}

func (s *CandleService) CandleMonths(ctx context.Context, market string, opts *CandleListOptions) ([]*Candle, *http.Response, error) {
	return s.candle(ctx, "months", market, opts)
}

func (s *CandleService) Ticker(ctx context.Context, markets []string) ([]*Ticker, *http.Response, error) {
	if len(markets) == 0 {
		return nil, nil, ErrInvalidArguments
	}

	u := fmt.Sprintf("v1/ticker?markets=%s", strings.Join(markets, ","))
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var tickers []*Ticker

	resp, err := s.client.Do(ctx, req, &tickers)
	if err != nil {
		return nil, resp, err
	}

	return tickers, nil, nil
}

func (s *CandleService) TickerMarket(ctx context.Context, market string) (*Ticker, *http.Response, error) {
	lst, resp, err := s.Ticker(ctx, []string{market})
	if err != nil {
		return nil, resp, err
	}

	return lst[0], resp, err
}
