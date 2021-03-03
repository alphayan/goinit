package main

type User struct {
	ID int64
}

func (u *User) create() error {
	_, err := connectionDB.Insert(u)
	return err
}

func (u *User) update() error {
	_, err := connectionDB.ID(u.ID).Update(u)
	return err
}

func (u *User) get() error {
	_, err := connectionDB.ID(u.ID).Get(u)
	return err
}

func (u *User) gets(limit, page int) (interface{}, int64) {
	us, count := []User{}, int64(0)
	connectionDB.Where("").Find(&us)
	connectionDB.Where("").Count(&count)
	return us, count
}

func (u *User) delete() error {
	_, err := connectionDB.ID(u.ID).Delete(u)
	return err
}
