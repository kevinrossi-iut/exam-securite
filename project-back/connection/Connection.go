package connection

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ConnetionInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func Connect(w http.ResponseWriter, r *http.Request) {
	var info ConnetionInfo
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("connection...")
	connectedUser := info.c	onnect()

	if connectedUser == nil {
		log.Println("connection error")
		http.Error(w, "wrong user/password", http.StatusForbidden)
		return
	}
	log.Println("connection success")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<div id="formContent">
        <div>
            <p id="welcome" class="success"></p>
        </div>
        <div>
            <span>You are now connected.</span>
            <span>Welcome `+connectedUser.User+`</span>
        </div>
    </div>`)
}

type ConnectedUser struct {
	Id       int    `json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
	Session  string `json:"session"`
}

func (info ConnetionInfo) connect() *ConnectedUser {
	row := notSecuredQuery{info}.GetResult()

	if !row.Next() {
		return nil
	}
	var user ConnectedUser
	_ = row.Scan(&user.Id, &user.User, &user.Password)
	return &user
}
