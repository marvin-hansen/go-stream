package messages

import (
	"github.com/shopspring/decimal"
	"time"
)

type MessageType string

const (
	TRADE     MessageType = "trade"
	QUOTE     MessageType = "quote"
	BOOK      MessageType = "book"
	BOOK5     MessageType = "book5"
	BOOK20    MessageType = "book20"
	BOOK50    MessageType = "book50"
	OHLCV     MessageType = "ohlcv"
	VOLUME    MessageType = "volume"
	HEARTBEAT MessageType = "hearbeat"
	HELLO     MessageType = "hello"
	ERROR     MessageType = "error"
	EXRATE    MessageType = "exrate"
	RECONNECT MessageType = "reconnect"
)

type Ask struct {
	Type  MessageType     `json:"type"`
	Price decimal.Decimal `json:"price"`
	Size  decimal.Decimal `json:"size"`
}

type Bid struct {
	Type  MessageType     `json:"type"`
	Price decimal.Decimal `json:"price"`
	Size  decimal.Decimal `json:"size"`
}

type ErrorMessage struct {
	Type    MessageType `json:"type"`
	Message string      `json:"message"`
}

type ExchangeRate struct {
	// https://docs.coinapi.io/#exchange-rate-in
	Type           MessageType     `json:"type"`
	Asset_id_base  string          `json:"asset_id_base"`
	Asset_id_quote string          `json:"asset_id_quote"`
	Time           time.Time       `json:"time"`
	Rate           decimal.Decimal `json:"rate"`
}

type Orderbook struct {
	// https://docs.coinapi.io/#orderbook-l2-full-in
	Type          MessageType `json:"type"`
	Symbol_id     string      `json:"symbol_id"`
	Sequence      uint32      `json:"sequence"`
	Time_exchange time.Time   `json:"time_exchange"`
	Time_coinapi  time.Time   `json:"time_coinapi"`
	Asks          []Bid       `json:"asks"`
	Bids          []Bid       `json:"bids"`
}

type Ohlcv struct {
	// https://docs.coinapi.io/#ohlcv-in
	Type              MessageType     `json:"type"`
	Symbol_id         string          `json:"symbol_id"`
	Sequence          uint32          `json:"sequence"`
	PeriodID          string          `json:"period_id"`
	Time_period_start time.Time       `json:"time_period_start"`
	Time_period_end   time.Time       `json:"time_period_end"`
	Time_open         time.Time       `json:"time_open"`
	Time_close        time.Time       `json:"time_close"`
	Price_open        decimal.Decimal `json:"price_open"`
	Price_high        decimal.Decimal `json:"price_high"`
	Price_low         decimal.Decimal `json:"price_low"`
	Price_close       decimal.Decimal `json:"price_close"`
	Volume_traded     decimal.Decimal `json:"volume_traded"`
	Trades_count      uint32          `json:"trades_count"`
}

type Quote struct {
	// https://docs.coinapi.io/#quotes-in
	Type          MessageType     `json:"type"`
	Symbol_id     string          `json:"symbol_id"`
	Sequence      uint32          `json:"sequence"`
	Time_exchange time.Time       `json:"time_exchange"`
	Time_coinapi  time.Time       `json:"time_coinapi"`
	Ask_price     decimal.Decimal `json:"ask_price"`
	Ask_size      decimal.Decimal `json:"ask_size"`
	Bid_price     decimal.Decimal `json:"bid_price"`
	Bid_size      decimal.Decimal `json:"bid_size"`
}
type Reconnect struct {
	// https://docs.coinapi.io/#reconnect-in
	Type           MessageType `json:"type"`
	Within_seconds uint32      `json:"within_seconds"`
	Before_time    time.Time   `json:"before_time"`
}

type Hearbeat struct {
	// https://docs.coinapi.io/#heartbeat-in
	Type MessageType `json:"type"`
}

type Trade struct {
	// https://docs.coinapi.io/#trades-in
	Type          MessageType     `json:"type"`
	Symbol_id     string          `json:"symbol_id"`
	Sequence      uint32          `json:"sequence"`
	Time_exchange time.Time       `json:"time_exchange"`
	Time_coinapi  time.Time       `json:"time_coinapi"`
	Uuid          string          `json:"uuid"`
	Price         decimal.Decimal `json:"price"`
	Size          decimal.Decimal `json:"size"`
	Taker_side    string          `json:"taker_side"`
}

type Volume struct {
	// https://docs.coinapi.io/#volume-in
	Type             MessageType      `json:"type"`
	Period_id        string           `json:"period_id"`
	Time_coinapi     time.Time        `json:"time_coinapi"`
	Volume_by_symbol []VolumeBySymbol `json:"volume_by_symbol"`
}

type VolumeBySymbol struct {
	// https://docs.coinapi.io/#volume-in
	Symbol_id      string          `json:"symbol_id"`
	Asset_id_base  string          `json:"asset_id_base"`
	Asset_id_quote string          `json:"asset_id_quote"`
	Volume_base    decimal.Decimal `json:"volume_base"`
	Volume_quote   decimal.Decimal `json:"volume_quote"`
}
