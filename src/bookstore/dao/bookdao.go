package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"strconv"
)

func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func AddBook(b *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(bookID string) error {
	//写sql语句
	sqlStr := "delete from books where id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

func GetBookByID(bookID string) (*model.Book, error) {
	sql := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	row := utils.Db.QueryRow(sql, bookID)
	book := &model.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(b *model.Book) error {
	sql := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	_, err := utils.Db.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetPageBooks(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	sql := "select count(*) from books"
	var totalRecord int64
	row := utils.Db.QueryRow(sql)
	err := row.Scan(&totalRecord)
	if err != nil {
		return nil, err
	}
	var pageSize int64 = 4
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	sql2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	rows, err := utils.Db.Query(sql2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

func GetBooksByPrice(pageNo string, minPrice string, maxPrice string) (*model.Page, error) {
	iPageNo, err := strconv.ParseInt(pageNo, 10, 64)
	sql := "select count(*) from books where price between ? and ?"
	var totalRecord int64
	row := utils.Db.QueryRow(sql, minPrice, maxPrice)
	err = row.Scan(&totalRecord)
	if err != nil {
		return nil, err
	}
	var pageSize int64 = 4
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	sql2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"
	rows, err := utils.Db.Query(sql2, minPrice, maxPrice, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}
