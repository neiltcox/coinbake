package service

import (
	"fmt"
)

type ExchangeIdentifier string

const (
	ExchangeIdentifierMocked      ExchangeIdentifier = "mock"
	ExchangeIdentifierKraken      ExchangeIdentifier = "kraken"
	ExchangeIdentifierCoinbasePro ExchangeIdentifier = "coinbasepro"
	ExchangeIdentifierBinance     ExchangeIdentifier = "binance"
)

var exchanges map[ExchangeIdentifier]Exchange = make(map[ExchangeIdentifier]Exchange)

func init() {
	exchanges[ExchangeIdentifierMocked] = &ExchangeMocked{}
	exchanges[ExchangeIdentifierKraken] = &ExchangeKraken{}
}

type SupportedAsset struct {
	Asset Asset
}

// An interface representing a generic exchange.
type Exchange interface {
	CreateOrder(*Portfolio, string, float32) (CreatedOrder, error)
	Holdings(*Portfolio) (map[string]Holding, error)
	SupportedAssets(*Portfolio) (map[string]bool, error)
	SupportsAsset(*Portfolio, Asset) bool
}

type MockSupportedAssets struct {
}

type CreatedOrder struct {
	OrderIdentifier string
}

type Holding struct {
	Asset   Asset
	Balance float64
}

type ExchangeKraken struct {
}

type ExchangeMocked struct {
	MockSupportedAssets
}

func (mockSupportedAssets *MockSupportedAssets) SupportedAssets(exchangeConnection *Portfolio) (map[string]bool, error) {
	return map[string]bool{
		"BTC": true,
		"ETH": true,
		"XMR": true,
	}, nil
}

func (exchangeKraken *ExchangeKraken) SupportedAssets(exchangeConnection *Portfolio) (map[string]bool, error) {
	// TODO: implement
	return map[string]bool{}, nil
}

// Gets the Exchange object for a given Exchange Connection, which is where the API call logic is.
func (exchangeConnection *Portfolio) Exchange() (Exchange, error) {
	exchange, exists := exchanges[exchangeConnection.ExchangeIdentifier]
	if !exists {
		return nil, fmt.Errorf("exchange %q is not implemented", exchangeConnection.ExchangeIdentifier)
	}

	return exchange, nil
}

func (exchangeMocked *ExchangeMocked) CreateOrder(exchangeConnection *Portfolio, asset string, amount float32) (CreatedOrder, error) {
	return CreatedOrder{
		OrderIdentifier: "123456",
	}, nil
}

func (exchangeMocked *ExchangeMocked) Holdings(exchangeConnection *Portfolio) (map[string]Holding, error) {
	return map[string]Holding{
		"BTC": {Asset: FindAssetBySymbol("BTC"), Balance: 0.23},
		"ETH": {Asset: FindAssetBySymbol("ETH"), Balance: 2.3},
		"XMR": {Asset: FindAssetBySymbol("XMR"), Balance: 43.145},
		"BNB": {Asset: FindAssetBySymbol("BNB"), Balance: 0.033},
		"ADA": {Asset: FindAssetBySymbol("ADA"), Balance: 50.2},
	}, nil
}

func (exchangeMocked *ExchangeMocked) SupportsAsset(exchangeConnection *Portfolio, asset Asset) bool {
	return true
}

func (exchangeKraken *ExchangeKraken) CreateOrder(exchangeConnection *Portfolio, asset string, amount float32) (CreatedOrder, error) {
	return CreatedOrder{
		OrderIdentifier: "123456",
	}, nil
}

func (exchangeKraken *ExchangeKraken) Holdings(exchangeConnection *Portfolio) (map[string]Holding, error) {
	return map[string]Holding{}, nil
}

func (exchangeKraken *ExchangeKraken) SupportsAsset(exchangeConnection *Portfolio, asset Asset) bool {
	return true
}
