package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"html/template"
	"net/http"
)

func LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		value := cookie.Value
		_ = dao.DeleteSession(value)
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	GetPageBooksByPrice(w, r)
}

func Login(w http.ResponseWriter, r *http.Request) {
	flag, _ := dao.IsLogin(r)
	if flag {
		GetPageBooksByPrice(w, r)
	} else {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		user, _ := dao.CheckUserNameAndPassword(username, password)
		if user.ID > 0 {
			uuid := utils.CreateUUID()
			session := &model.Session{
				SessionID: uuid,
				Username:  user.Username,
				UserID:    user.ID,
			}
			_ = dao.AddSession(session)
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			_ = t.Execute(w, user)
		} else {
			//用户名或密码不正确
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			_ = t.Execute(w, "用户名或密码不正确！")
		}
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		_ = t.Execute(w, "用户名已存在！")
	} else {
		_ = dao.SaveUser(username, password, email)
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		_ = t.Execute(w, "")
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		w.Write([]byte("用户名已存在！"))
	} else {
		//用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}
