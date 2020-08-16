package upbit_test

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/investing-kr/go-upbit"
)

var testOrder = flag.Bool("order", false, "test order")

func TestClient_Accounts(t *testing.T) {
	accounts, _, err := c.Accounts.Accounts(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, account := range accounts {
		t.Logf("%+v", account)
	}
}

func TestMarketCodes(t *testing.T) {
	ctx := context.Background()
	markets, _, err := c.Markets.All(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(markets) == 0 {
		t.Fatal("len(markets) == 0")
	}

	for _, market := range markets {
		t.Logf("%+v", market)
	}
}

func TestOrder(t *testing.T) {
	if *testOrder == false {
		t.Skip("skipiing TestOrder")
	}

	ctx := context.Background()
	order, _, err := c.Orders.Order(ctx, &upbit.OrderRequest{
		Market:  upbit.KRW_HBAR,
		Side:    upbit.SideSell,
		Volume:  "1",
		Price:   "1000",
		OrdType: upbit.OrdTypeLimit,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", order)

	order, _, err = c.Orders.GetOrderByUUID(ctx, order.UUID)
	if err != nil {
		t.Fatal(err)
	}

	order, _, err = c.Orders.CancelOrderByUUID(ctx, order.UUID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", order)
}

func TestOrderList(t *testing.T) {
	ctx := context.Background()

	orders, _, err := c.Orders.ListOrders(ctx, &upbit.OrderListOptions{
		// Market: "KRW-HBAR",
		State: upbit.OrderStateWait,
		// States: []upbit.OrderState{upbit.OrderStateWait},
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, o := range orders {
		t.Logf("%+v", o)
	}
}

func TestChances(t *testing.T) {
	ctx := context.Background()

	chance, _, err := c.Orders.Chances(ctx, upbit.KRW_BTC)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(chance)
}

var c *upbit.Client

func TestMain(m *testing.M) {
	var err error
	c, err = upbit.NewClient(nil, upbit.ClientOptionsFromEnv())
	if err != nil {
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestCandles(t *testing.T) {
	ctx := context.Background()

	candles, _, err := c.Candles.CandleMinutes(ctx, upbit.KRW_BTC, 1, &upbit.CandleListOptions{
		Count: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	candles, _, err = c.Candles.CandleMonths(ctx, upbit.KRW_BTC, nil)
	if err != nil {
		t.Fatal(err)
	}

	candles, _, err = c.Candles.CandleDays(ctx, upbit.KRW_BTC, nil)
	if err != nil {
		t.Fatal(err)
	}

	candles, _, err = c.Candles.CandleWeeks(ctx, upbit.KRW_BTC, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(candles)
}

func TestTicker(t *testing.T) {
	ctx := context.Background()

	tickers, _, err := c.Candles.Ticker(ctx, []string{upbit.BTC_ADA})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(*(tickers[0]))
}

func TestCoinAddresses(t *testing.T) {
	ctx := context.Background()

	addresses, _, err := c.Deposits.ListCoinAddresses(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(addresses)
}
