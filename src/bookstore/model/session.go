package model

type Session struct {
	SessionID string
	Username  string
	UserID    int
	Cart      *Cart
	OrderID   string
	Orders    []*Order
}
