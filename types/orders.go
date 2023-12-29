package types

type Order struct {
	Name        string `json:"name"`
	OrderID     string `json:"orderID"`
	Sku         string `json:"sku"`
	Description string `json:"description"`
	Address     string `json:"address"`
}
