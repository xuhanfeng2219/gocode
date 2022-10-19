package session

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func SaveSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["foo"] = "bar"
	session.Values[42] = 43

	if err = session.Save(r, w); err != nil {
		fmt.Println("save wrong!")
	}
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	foo := session.Values["foo"]
	fmt.Println(foo)
}
