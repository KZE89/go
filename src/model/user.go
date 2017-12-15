package model

import ("fmt")

type User struct {
	ID         int64  `gorm:primary key;not_nil`
	Login      string `gorm:not_nil`
	Pass       string `gorm:not_nil`
	Worknumber int `gorm:column:worknumber;not_nil`
}

func Get(u *User, login string, pass string) error {
	
	query := fmt.Sprintf("login = '%s' and pass = '%s'", login, pass)
	GormInit()
	err := DBConn.Where(query).First(&u).Error
	GormClose()
	return err
	
}

func Save(u *User, newPass string) error {
	query := fmt.Sprintf("login = '%s' and pass = '%s'", u.Login, u.Pass)
	GormInit()
	err := DBConn.Where(query).First(&u).Error

	if err == nil {
		u.Pass = newPass
		err = DBConn.Save(u).Error
		GormClose()
		return err
	}
	
	GormClose()
	return err
}
