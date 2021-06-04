package web

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var store *sessions.CookieStore

func InitSession() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 24,
		HttpOnly: true,
	}
}

func CreateCookie(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	session.Values["user"] = "name;instance"
	_ = session.Save(r, w)
}

func GetCookie(r *http.Request) {
	session, _ := store.Get(r, "user")
	log.Println(session.Values)
}