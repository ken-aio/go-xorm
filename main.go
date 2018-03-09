package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/k0kubun/pp"
	"github.com/ken-aio/go-xorm/models"
)

var engine *xorm.Engine

func main() {
	con, err := xorm.NewEngine("mysql", "sampledb:sampledb@tcp([localhost]:3306)/sampledb?parseTime=true&loc=Asia%2FTokyo")
	con.ShowSQL(true)
	if err != nil {
		panic(err)
	}

	user1 := &models.Users{}
	user1.Name = "akiho"
	user1.Gender = "male"
	_, err = con.Insert(user1) // 最初の戻りは成功row数
	if err != nil {
		panic(err)
	}

	user2 := &models.Users{}
	user2.Name = "funaki"
	user2.Gender = "male"
	_, err = con.Insert(user2)
	if err != nil {
		panic(err)
	}

	var users []models.Users
	if err := con.Asc("id").Find(&users); err != nil {
		panic(err)
	}
	pp.Println("-----1. Show all users")
	for i := 0; i < len(users); i++ {
		pp.Println(users[i])
	}

	var akiho models.Users
	if _, err = con.Where("name = ?", "akiho").Get(&akiho); err != nil {
		panic(err)
	}
	akiho.Gender = "female"
	con.ID(akiho.Id).Update(&akiho)
	pp.Println("-----2. Show reload akiho user")
	pp.Println(akiho)

	if _, err = con.ID(akiho.Id).Delete(&akiho); err != nil {
		panic(err)
	}
	ok, err := con.Where("name = ?", "akiho").Get(&akiho)
	if err != nil {
		panic(err)
	}
	if !ok {
		pp.Println("-----3. Show reload not found user")
	}
	pp.Println("-----4. Show akiho user")
	pp.Println(akiho)
}
