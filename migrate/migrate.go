package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sns-barko/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, sec := config.InitConfig()
	db := initDatabse(cfg, sec)
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	path, _ := filepath.Abs("./migrate/" + "migrate.sql")
	data, err := os.ReadFile(path)
	if err != nil {
		sqlDb.Close()
		log.Fatal(err)
	}

	if err := db.Exec(string(data)).Error; err != nil {
		log.Println(err)
		return
	}
	log.Println("SuccessFully Migrate User Data")
	sqlDb.Close()
}

func initDatabse(cfg *config.Config, secret *config.Secret) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		secret.Database.User,
		secret.Database.Password,
		secret.Database.Host,
		secret.Database.Port,
		cfg.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NowFunc: func() time.Time { return time.Now().Local() }})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
