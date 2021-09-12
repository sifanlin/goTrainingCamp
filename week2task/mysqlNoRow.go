package mysqlNoRow

import (
	"database/sql"
	"github.com/pkg/errors"
)

var db *sql.DB

func init() {
	DSN := "user:password@/dbname"
	db, _ = sql.Open("mysql", DSN)
}

type peopleInfo struct {
	Id   int64
	Name string
}

func selects() ([]*peopleInfo, error) {

	obj := make([]*peopleInfo, 0)
	age := "27"
	querySql := "selct id name from info where age :=?"
	rows, err := db.Query(querySql, age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return obj, nil
		}
		return obj, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			continue
		}
		data := &peopleInfo{
			Id:   id,
			Name: name,
		}
		obj = append(obj, data)
	}
	return obj, nil
}
