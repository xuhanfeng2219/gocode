package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

func AddOrder(order *model.Order) error {
	sql := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}

func GetOrders() ([]*model.Order, error) {
	sql := "select id,create_time,total_count,total_amount,state,user_id from orders"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		e := rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		if e != nil {
			return nil, e
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func GetMyOrders(userID int) ([]*model.Order, error) {
	sql := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id = ?"
	rows, err := utils.Db.Query(sql, userID)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		e := rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		if e != nil {
			return nil, e
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func UpdateOrderState(orderID string, state int64) error {
	sql := "update orders set state = ? where id = ?"
	_, err := utils.Db.Exec(sql, state, orderID)
	if err != nil {
		return err
	}
	return nil
}
