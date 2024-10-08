package mysql

import (
	"fmt"
	"github.com/RSRFV/online-storage/lib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	//"online-storage/lib"
)

var DB *gorm.DB

func InitDB(conf lib.ServerConfig) {
	var err error
	fmt.Printf("conf.User:", conf.User, "\n")
	fmt.Printf("conf.Password:", conf.Password, "\n")
	fmt.Printf("conf.Host:", conf.Host, "\n")
	fmt.Printf("conf.DbName:", conf.DbName, "\n")
	dbParams := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
	)
	DB, err = gorm.Open("mysql", dbParams)
	if err != nil {
		log.Fatal(2, err)
	}

	// 全局禁用表名复数
	DB.SingularTable(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	fmt.Println("database init on port ", conf.Host)
}
