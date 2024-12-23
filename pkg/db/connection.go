package db

import (
	"ecommerce_clean_architecture/pkg/config"
	"ecommerce_clean_architecture/pkg/domain"
	"ecommerce_clean_architecture/pkg/utils/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate(&domain.AdminDetails{})
	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Products{})
	db.AutoMigrate(&models.UserSignUp{})
	db.AutoMigrate(&models.UserLogin{})
	db.AutoMigrate(&models.OTP{})
	db.AutoMigrate(&models.TempUser{})

	return db, dbErr

}
