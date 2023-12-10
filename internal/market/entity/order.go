package entity

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PeddindShares int
	price         float64
	OrderType     string
	status        string
	transactions  []*Transaction
}

func NewOrder(id string, investor *Investor, asset *Asset, shares int, price float64, orderType string) *Order {
	return &Order{
		ID:            id,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PeddindShares: shares,
		price:         price,
		OrderType:     orderType,
		status:        "OPEN",
		transactions:  []*Transaction{},
	}
}
