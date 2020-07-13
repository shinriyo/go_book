import "time"

func authenticate(w http.TrsponseWrite, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostForValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
} 

cookie := http:Cookie {
	Name: "__cookie",
	Value: session.Uuid,
	HttpOnly: true,
}