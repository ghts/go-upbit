package upbit

import (
	"encoding/json"
	"fmt"
	"time"
)

type OrderState string

const (
	OrderStateWait   OrderState = "wait"
	OrderStateDone   OrderState = "done"
	OrderStateCancel OrderState = "cancel"
)

type OrderKind string

const (
	OrderKindNormal OrderKind = "normal"
	OrderKindWatch  OrderKind = "watch"
)

type Account struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

type OrderRequest struct {
	Market     string    `url:"market,omitempty"`
	Side       OrderSide `url:"side,omitempty"`
	Volume     string    `url:"volume,omitempty"`
	Price      string    `url:"price,omitempty"`
	OrdType    OrdType   `url:"ord_type,omitempty"`
	Identifier string    `url:"identifier,omitempty"`
}

type OrderListOptions struct {
	Market      string       `url:"market,omitempty"`
	State       OrderState   `url:"state,omitempty"`
	States      []OrderState `url:"states,brackets"`
	UUIDs       []string     `url:"uuids,brackets"`
	Identifiers []string     `url:"identifiers,brackets"`
	Kind        OrderKind    `url:"kind,omitempty"`
	Page        int          `url:"page,omitempty"`
	Limit       int          `url:"limit,omitempty"`
	OrderBy     string       `url:"order_by,omitempty"`
}

type OrderSide string

const (
	SideBid  OrderSide = "bid" // Buying, 매수
	SideBuy  OrderSide = "bid" // Buying, 매수
	SideAsk  OrderSide = "ask" // Selling, 매도
	SideSell OrderSide = "ask" // Selling, 매도
)

type OrdType string

const (
	OrdTypeLimit  OrdType = "limit"  // Limit Order, 지정가
	OrdTypePrice  OrdType = "price"  // Market Price Order(Bid), 시장가(매수)
	OrdTypeMarket OrdType = "market" // Market Price Order(Ask), 시장가(매도)
)

type Order struct {
	UUID            string    `json:"uuid"`
	Side            string    `json:"side" header:"side"`
	OrdType         string    `json:"ord_type" header:"ord_type" `
	Price           string    `json:"price" header:"price"`
	AvgPrice        string    `json:"avg_price" `
	State           string    `json:"state" header:"state"`
	Market          string    `json:"market" header:"market"`
	CreatedAt       time.Time `json:"created_at" header:"created_at"`
	Volume          string    `json:"volume" header:"volume"`
	RemainingVolume string    `json:"remaining_volume" header:"remaining_volume"`
	ReservedFee     string    `json:"reserved_fee"`
	RemainingFee    string    `json:"remaining_fee"`
	PaidFee         string    `json:"paid_fee"`
	Locked          string    `json:"locked"`
	ExecutedVolume  string    `json:"executed_volume"`
	TradesCount     int       `json:"trades_count"`
}

type MarketCode struct {
	Market        string `json:"market"`
	KoreanName    string `json:"korean_name"`
	EnglishName   string `json:"english_name"`
	MarketWarning string `json:"market_warning"`
}

type Orderbook struct {
	Market         string          `json:"market"`
	Timestamp      int64           `json:"timestamp"`
	TotalAskSize   float64         `json:"total_ask_size"`
	TotalBidSize   float64         `json:"total_bid_size"`
	OrderbookUnits []OrderbookUnit `json:"orderbook_units"`

	Type string `json:"type,omitempty"` // for websocket response
	Code string `json:"code,omitempty"` // for websocket response
}

type OrderbookUnit struct {
	AskPrice float64 `json:"ask_price"`
	BidPrice float64 `json:"bid_price"`
	AskSize  float64 `json:"ask_size"`
	BidSize  float64 `json:"bid_size"`
}

type Ticker struct {
	Market              string  `json:"market"`
	TradeDate           string  `json:"trade_date"`
	TradeTime           string  `json:"trade_time"`
	TradeDateKst        string  `json:"trade_date_kst"`
	TradeTimeKst        string  `json:"trade_time_kst"`
	TradeTimestamp      int64   `json:"trade_timestamp"`
	OpeningPrice        float64 `json:"opening_price"`
	HighPrice           float64 `json:"high_price"`
	LowPrice            float64 `json:"low_price"`
	TradePrice          float64 `json:"trade_price"`
	PrevClosingPrice    float64 `json:"prev_closing_price"`
	Change              string  `json:"change"`
	ChangePrice         float64 `json:"change_price"`
	ChangeRate          float64 `json:"change_rate"`
	SignedChangePrice   float64 `json:"signed_change_price"`
	SignedChangeRate    float64 `json:"signed_change_rate"`
	TradeVolume         float64 `json:"trade_volume"`
	AccTradePrice       float64 `json:"acc_trade_price"`
	AccTradePrice24H    float64 `json:"acc_trade_price_24h"`
	AccTradeVolume      float64 `json:"acc_trade_volume"`
	AccTradeVolume24H   float64 `json:"acc_trade_volume_24h"`
	Highest52_WeekPrice float64 `json:"highest_52_week_price"`
	Highest52_WeekDate  string  `json:"highest_52_week_date"`
	Lowest52_WeekPrice  float64 `json:"lowest_52_week_price"`
	Lowest52_WeekDate   string  `json:"lowest_52_week_date"`
	Timestamp           int64   `json:"timestamp"`

	Type string `json:"type,omitempty"` // for websocket response
	Code string `json:"code,omitempty"` // for websocket response
}

type WebsocketRequest struct {
	Ticket string
	Type   []WebsocketRequestType
	// Format string
}

type WebsocketRequestType struct {
	Type           string   `json:"type,omitempty"`
	Codes          []string `json:"codes,omitempty"`
	IsOnlySnapShot bool     `json:"isOnlySnapshot,omitempty"`
	IsOnlyRealtime bool     `json:"isOnlyRealtime,omitempty"`
}

func (r *WebsocketRequest) MarshalJSON() ([]byte, error) {
	b := []byte{}
	b = append(b, []byte(fmt.Sprintf(`[{"ticket":"%s"}`, r.Ticket))...)

	for _, v := range r.Type {
		typ, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		b = append(b, []byte(",")...)
		b = append(b, typ...)
	}
	b = append(b, []byte("]")...)

	return b, nil
}

type Chance struct {
	BidFee      string `json:"bid_fee"`
	AskFee      string `json:"ask_fee"`
	MakerBidFee string `json:"maker_bid_fee"`
	MakerAskFee string `json:"maker_ask_fee"`
	Market      struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		OrderTypes []string `json:"order_types"`
		OrderSides []string `json:"order_sides"`
		Bid        struct {
			Currency  string  `json:"currency"`
			PriceUnit float64 `json:"price_unit"`
			MinTotal  float64 `json:"min_total"`
		} `json:"bid"`
		Ask struct {
			Currency  string  `json:"currency"`
			PriceUnit float64 `json:"price_unit"`
			MinTotal  float64 `json:"min_total"`
		} `json:"ask"`
		MaxTotal string `json:"max_total"`
		State    string `json:"state"`
	} `json:"market"`
	BidAccount struct {
		Currency            string `json:"currency"`
		Balance             string `json:"balance"`
		Locked              string `json:"locked"`
		AvgBuyPrice         string `json:"avg_buy_price"`
		AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
		UnitCurrency        string `json:"unit_currency"`
	} `json:"bid_account"`
	AskAccount struct {
		Currency            string `json:"currency"`
		Balance             string `json:"balance"`
		Locked              string `json:"locked"`
		AvgBuyPrice         string `json:"avg_buy_price"`
		AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
		UnitCurrency        string `json:"unit_currency"`
	} `json:"ask_account"`
}

type Candle struct {
	Market               string  `json:"market"`
	CandleDateTimeUtc    string  `json:"candle_date_time_utc"`
	CandleDateTimeKst    string  `json:"candle_date_time_kst"`
	OpeningPrice         float64 `json:"opening_price"`
	HighPrice            float64 `json:"high_price"`
	LowPrice             float64 `json:"low_price"`
	TradePrice           float64 `json:"trade_price"`
	Timestamp            int64   `json:"timestamp"`
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"`
	PrevClosingPrice     float64 `json:"prev_closing_price"`
	ChangePrice          float64 `json:"change_price"`
	ChangeRate           float64 `json:"change_rate"`
	FirstDayOfPeriod     string  `json:"first_day_of_period"`
}

type CoinAddress struct {
	Currency         string  `json:"currency"`
	DepositAddress   string  `json:"deposit_address"`
	SecondaryAddress string `json:"secondary_address"`
}
