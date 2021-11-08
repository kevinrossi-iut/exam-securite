package connection

import (
	"dataacces"
)

type notSecuredQuery struct {
	info ConnetionInfo
}
type securedQuery struct {
	info ConnetionInfo
}

func getUser(id int) *ConnectedUser {
	rows, _ := dataacces.ExecQuery("SELECT ID, USER, PASSWORD FROM USER WHERE ID = ?", id)
	if !rows.Next() {
		return nil
	}
	var user ConnectedUser
	_ = rows.Scan(&user.Id, &user.User, &user.Password)
	return &user
}
