package model

type CartItem struct {
	CartItemID int64
	Book       *Book
	Count      int64
	Amount     float64
	CartID     string
}

func (cartItem *CartItem) GetAmount() float64 {
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}
