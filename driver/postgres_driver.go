package driver

import (
	"TDKTServer/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
)

var configs = config.LoadConfigs()

func Connect() *gorm.DB {
	URL := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v", configs.Database.Host, configs.Database.Port, configs.Database.User, configs.Database.DbName, configs.Database.Pass, configs.Database.SslMode)
	db, err := gorm.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
func StartPosgresDbConnection() {
	con := Connect()
	defer con.Close()
	err := con.DB().Ping()
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	fmt.Println("Postgresql was connected!")
}
