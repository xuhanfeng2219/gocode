package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag {
		bookID := r.FormValue("bookId")
		book, _ := dao.GetBookByID(bookID)
		userID := session.UserID
		cart, _ := dao.GetCartByUserID(userID)
		if cart != nil {
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			if cartItem != nil {
				items := cart.CartItems
				for _, v := range items {
					if v.Book.ID == cartItem.Book.ID {
						v.Count = v.Count + 1
						_ = dao.UpdateBookCount(v)
					}
				}
			} else {
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				cart.CartItems = append(cart.CartItems, cartItem)
				_ = dao.AddCartItem(cartItem)
			}
			_ = dao.UpdateCart(cart)
		} else {
			uuid := utils.CreateUUID()
			cart := &model.Cart{
				CartID: uuid,
				UserID: userID,
			}
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: uuid,
			}
			cartItems = append(cartItems, cartItem)
			cart.CartItems = cartItems
			_ = dao.AddCart(cart)
		}
	} else {
		w.Write([]byte("first login"))
	}
}

func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		session.Cart = cart
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		_ = t.Execute(w, session)
	} else {
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		_ = t.Execute(w, session)
	}
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartId")
	_ = dao.DeleteCartByCartID(cartID)
	GetCartInfo(w, r)
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	items := cart.CartItems
	for k, v := range items {
		if v.CartItemID == iCartItemID {
			items = append(items[k:], items[k+1:]...)
			cart.CartItems = items
			_ = dao.DeleteCartItemByID(cartItemID)
		}
	}
	_ = dao.UpdateCart(cart)
	GetCartInfo(w, r)
}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	items := cart.CartItems
	for _, v := range items {
		if v.CartItemID == iCartItemID {
			v.Count = iBookCount
			_ = dao.UpdateBookCount(v)
		}
	}

	_ = dao.UpdateCart(cart)
	cart, _ = dao.GetCartByUserID(userID)
	totalCount := cart.TotalCount
	totalAmount := cart.TotalAmount
	var amount float64
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if iCartItemID == v.CartItemID {
			amount = v.Amount
		}
	}

	data := model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}

	json, _ := json.Marshal(data)
	w.Write(json)
}
