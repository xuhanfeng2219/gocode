package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

func AddCartItem(item *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, item.Count, item.GetAmount(), item.Book.ID, item.CartID)
	if err != nil {
		return err
	}
	return nil
}

func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)

	item := &model.CartItem{}
	err := row.Scan(&item.CartItemID, &item.Count, &item.Amount, &item.CartID)
	if err != nil {
		return nil, err
	}
	book, _ := GetBookByID(bookID)
	item.Book = book
	return item, nil
}

func UpdateBookCount(item *model.CartItem) error {
	sql := "update cart_items set count = ? , amount = ? where book_id = ? and cart_id = ?"
	_, err := utils.Db.Exec(sql, item.Count, item.GetAmount(), item.Book.ID, item.CartID)
	if err != nil {
		return err
	}
	return nil
}

func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var items []*model.CartItem
	for rows.Next() {
		var bookID string
		item := &model.CartItem{}
		e := rows.Scan(&item.CartItemID, &item.Count, &item.Amount, &bookID, &item.CartID)
		if e != nil {
			return nil, e
		}
		book, _ := GetBookByID(bookID)
		item.Book = book
		items = append(items, item)
	}
	return items, nil
}

func DeleteCartItemsByCartID(cartID string) error {
	sql := "delete from cart_items where cart_id = ?"
	_, err := utils.Db.Exec(sql, cartID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCartItemByID(cartItemID string) error {
	sql := "delete from cart_items where id = ?"
	_, err := utils.Db.Exec(sql, cartItemID)
	if err != nil {
		return err
	}
	return nil
}
