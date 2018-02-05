package models

import (
	"time"
	"io"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"crypto/rand"
	"app-backend/models/db"
)

type AppUser struct {
	ID       int       `orm:"pk;column(id);auto" json:"id,omitempty"`
	Phone    string    `orm:"column(phone);unique" json:"phone,omitempty"`
	Name     string    `orm:"column(name)"     json:"name,omitempty"`
	Password string    `orm:"column(password)" json:"password,omitempty"`
	Salt     string    `orm:"column(salt)"     json:"salt,omitempty"`
	RegDate  time.Time `orm:"column(reg_date)" json:"reg_date,omitempty"`
}

const pwHashBytes = 64

func generateSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf), nil
}

func generatePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h), nil
}

func NewUser(r *RegisterForm) (u *AppUser, err error) {
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}
	hash, err := generatePassHash(r.Password, salt)
	if err != nil {
		return nil, err
	}
	user := AppUser{
		Phone:    r.Phone,
		Name:     r.Name,
		Salt:     salt,
		Password: hash,
		RegDate:  time.Now()}
	return &user, nil
}

func (u *AppUser) Insert() (int, error) {
	o := db.GetOrmer()
	ret, err := o.Insert(u)
	if err != nil {
		return int(ret), err
	}
	return int(ret), nil
}

func (u *AppUser) FindByPhone(phone string) (int, error) {
	o := db.GetOrmer()
	user := new(AppUser)
	err := o.QueryTable(user).Filter("phone", phone).One(u)
	if err != nil {
		return ErrDatabase, err
	}
	return SuccessDB, nil
}
