package models

type Currency string

const (
	ARS  Currency = "ARS"
	USD  Currency = "USD"
	USDT Currency = "USDT"
	BTC  Currency = "BTC"
	ETH  Currency = "ETH"
)

var mapCurrencies = map[string]Currency{
	string(ARS):  ARS,
	string(USD):  USD,
	string(USDT): USDT,
	string(BTC):  BTC,
	string(ETH):  ETH,
}

func CurrencyFromString(value string) Currency {
	return mapCurrencies[value]
}

// Method to get the string of the currency type.
func (c *Currency) String() string {
	return string(*c)
}

func GetCurrencies() map[string]Currency {
	return mapCurrencies
}
