package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Id     uint
	Name   string
	Family string
}

type GORMHandler struct {
	db *gorm.DB
}

func CreateConnection() (*GORMHandler, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &GORMHandler{
		db: db,
	}, nil
}

func (handler *GORMHandler) GetPeople() ([]User, error) {
	var users []User
	res, _ := handler.db.Find(&users).Rows()
	for res.Next() {
		err := handler.db.ScanRows(res, &users)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return users, nil
}
