package dao

import (
	"Week01/model"
	"database/sql"

	"github.com/pkg/errors"
)

type Dao struct {
	db *sql.DB
}

func (d *Dao) GetUserInfo(name string) *model.User, error {
	var user = model.User{Name: name}
	stmt, err := d.db.Prepare("select age from user where name = ?")
	if err != nil {
		return nil, errors.Wrap(err, "Prepare failed")
	}

	err = stmt.QueryRow(name).Scan(&user.Age)
	if err != nil && errors.Cause(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, "query failed")
	}
	return &user, nil
}
