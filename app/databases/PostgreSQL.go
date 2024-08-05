package databases

import (
	"be-empower-hr/app/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func InitDBpostgre(cfg *config.AppConfig) *gorm.DB {
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
// 		cfg.HOST, cfg.USER, cfg.PASSWORD, cfg.DBNAME, cfg.PORT)
// 	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		NamingStrategy: schema.NamingStrategy{
// 			TablePrefix: "finalproject.",
// 		},
// 	})
// 	if err != nil {
// 		panic("failed to connect to database: " + err.Error())
// 	}
// 	return DB
// }

func InitDBpostgre(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cfg.HOST, cfg.USER, cfg.PASSWORD, cfg.DBNAME, cfg.PORT)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	return DB
}
