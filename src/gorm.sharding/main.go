package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.sharding/model"
	"strconv"
	"time"
)

func LoadDb() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/user_log?parseTime=True&loc=Local&allowNativePasswords=true")
	if err != nil {
		panic(err.Error())
	} else {
		db.DB().SetConnMaxLifetime(time.Second * time.Duration(10))
		db.DB().SetMaxOpenConns(10)
		db.DB().SetMaxIdleConns(10)
		//db.SetLogger(logger.Logger)
		db.LogMode(true)
		db.Debug()
	}
	return
}

func main() {
	db, err := LoadDb()
	fmt.Println(db, err)
	// Generate a snowflake ID
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	//ModHash insert Test
	//for i := 0; i < 30; i++ {
	//	id := node.Generate()
	//	u := model.NewUser(id.Int64())
	//	u.Name = "n:" + strconv.FormatInt(u.Id, 10)
	//	fmt.Println(u.Db(),":", u.TableName())
	//	err := db.Create(u).Error
	//	fmt.Println(err)
	//}

	//ModHash Query Test
	//var id int64 = 1410147438529875968
	//u := model.NewUser(id)
	//db.Where("id", id)
	//err = db.First(u).Error
	//fmt.Println(u)

	//MonthSharding insert Test
	logId := node.Generate()
	ul := model.NewUserLog(logId.Int64())
	ul.Log = "l:" + strconv.FormatInt(ul.Id, 10)
	fmt.Println(ul.Db(), ":", ul.TableName())
	err = db.Create(ul).Error
	fmt.Println(err)

}
