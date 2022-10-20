package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"html/template"
	"net/http"
	"time"
)

func CheckOut(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	orderID := utils.CreateUUID()
	timeStr := time.Now().Format("2006-01-02 15:04:05")

	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userID),
	}

	_ = dao.AddOrder(order)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		book := v.Book
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   book.Title,
			Author:  book.Author,
			ImgPath: book.ImgPath,
			OrderID: orderID,
		}
		_ = dao.AddOrderItem(orderItem)
		book.Sales = book.Sales + int(v.Count)
		book.Stock = book.Stock - int(v.Count)
		_ = dao.UpdateBook(book)
	}
	_ = dao.DeleteCartByCartID(cart.CartID)
	session.OrderID = orderID
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	_ = t.Execute(w, session)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, _ := dao.GetOrders()
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	_ = t.Execute(w, orders)
}

func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	orders, _ := dao.GetMyOrders(userID)

	session.Orders = orders
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	//执行
	_ = t.Execute(w, session)
}

func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	orders, _ := dao.GetMyOrders(userID)
	session.Orders = orders
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	//执行
	_ = t.Execute(w, session)
}

func SendOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderId")
	_ = dao.UpdateOrderState(orderID, 1)
	GetOrders(w, r)
}

func TakeOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderId")
	_ = dao.UpdateOrderState(orderID, 2)
	GetMyOrders(w, r)
}
