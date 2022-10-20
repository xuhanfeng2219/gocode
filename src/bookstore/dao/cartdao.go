package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

func AddCart(cart *model.Cart) error {
	sql := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sql, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	items := cart.CartItems
	for _, item := range items {
		print(item)
	}
	return nil
}

func GetCartByUserID(userID int) (*model.Cart, error) {
	sql := "select id,total_count,total_amount,user_id from carts where user_id = ?"
	row := utils.Db.QueryRow(sql, userID)
	cart := &model.Cart{}
	err := row.Scan(&cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	cart.CartItems = cartItems
	return cart, nil
}

func UpdateCart(cart *model.Cart) error {
	sql := "update carts set total_count = ? , total_amount = ? where id = ?"
	_, err := utils.Db.Exec(sql, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCartByCartID(cartID string) error {
	err := DeleteCartItemsByCartID(cartID)
	if err != nil {
		return err
	}
	sql := "delete from carts where id = ?"
	_, e := utils.Db.Exec(sql, cartID)
	if e != nil {
		return e
	}
	return nil
}
