package postgre

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  type ConfigDB struct {
	DB_Host string
	DB_User string
	DB_Password     string
	DB_Name    string
	DB_Port string
}
  func (config *ConfigDB) InitialDB() *gorm.DB{
  dsn :=fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
config.DB_Host,
config.DB_User,
config.DB_Password,
config.DB_Name,
config.DB_Port)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil{
	  log.Fatal(err)
  }
  return db
  }