package connection

import (
	"dataacces"
	"database/sql"
	"log"
)

func (query notSecuredQuery) GetResult() *sql.Rows {
	rows, err := dataacces.ExecQuery(getUserNotSecureAsString(query.info))
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func getUserNotSecureAsString(info ConnetionInfo) string {
	query := `SELECT ID, USER, PASSWORD
            FROM USER
            WHERE USER   = '` + info.User + `'
            AND PASSWORD = '` + info.Password + `'`
	log.Println(query)
	return query
}
