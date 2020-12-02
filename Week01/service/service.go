package service

import (
	"Week01/model"
	"Week01/dao"
	"database/sql"
	"github.com/pkg/errors"
)

type Service struct {
	Users []model.User
}

func (s *Service)TagUser(userName string, tag int) error {
	d := dao.dao{}
	user err := d.GetUserInfo(userName)
	if err != nil {
		if errors.Cause(err, sql.ErrNoRows) {
			return nil
		}
		return errors.Wrap(err, "failed")
	}
}