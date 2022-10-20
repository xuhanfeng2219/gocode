package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}

	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		page, _ = dao.GetBooksByPrice(pageNo, minPrice, maxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}

	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = session.Username
	}

	t := template.Must(template.ParseFiles("views/index.html"))
	err := t.Execute(w, page)
	if err != nil {
		log.Printf("run err:%s", err.Error())
	}
}

func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	_ = t.Execute(w, page)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	_ = dao.DeleteBook(bookID)
	GetPageBooks(w, r)
}

func ToUpdatePageBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	book, _ := dao.GetBookByID(bookID)
	if book.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		_ = t.Execute(w, book)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		_ = t.Execute(w, "")
	}
}

func AddOrUpdateBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	iBookID, _ := strconv.ParseInt(bookID, 10, 0)
	book := &model.Book{
		ID:      int(iBookID),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	if book.ID > 0 {
		_ = dao.UpdateBook(book)
	} else {
		_ = dao.AddBook(book)
	}
	GetPageBooks(w, r)
}
