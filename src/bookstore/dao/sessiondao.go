package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"net/http"
)

func AddSession(s *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, s.SessionID, s.Username, s.UserID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSession(sessionID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr, sessionID)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(sessionID string) (*model.Session, error) {
	sqlStr := "select session_id,username,user_id from sessions where session_id = ?"
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	row := inStmt.QueryRow(sessionID)
	s := &model.Session{}
	e := row.Scan(&s.SessionID, &s.Username, &s.UserID)
	if e != nil {
		return nil, e
	}
	return s, nil
}

func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		value := cookie.Value
		session, _ := GetSession(value)
		if session.UserID > 0 {
			return true, session
		}
	}
	return false, nil
}
