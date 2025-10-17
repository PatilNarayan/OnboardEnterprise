package db

import (
	"core/db/dao"
	"core/models"
	"fmt"
	"log"
	"os"

	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	once      sync.Once
	DbHandler *dao.Query
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func NewDBConfig() (*DBConfig, error) {

	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}, nil
}

func (oc *DBConfig) DSN() string {
	return "host=" + oc.Host + " port=" + oc.Port + " user=" + oc.User + " password=" + oc.Password + " dbname=" + oc.DbName + " sslmode=disable TimeZone=Asia/Shanghai"
}

func ConnectDB() error {
	cfg, err := NewDBConfig()
	if err != nil {
		return err
	}

	// Open a new database connection with prepared statements disabled
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.DSN(),
		PreferSimpleProtocol: true, // Use simple protocol for PostgreSQL
	}), &gorm.Config{
		PrepareStmt: false, // Disable prepared statements
	})

	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}
	dao.SetDefault(DB)
	DbHandler = dao.Q
	log.Println("Connected successfully to the database")
	return nil
}

func MigrateSchema() {
	if err := DB.AutoMigrate(
		&models.Organization{},
		&models.Team{},
		&models.User{},
		&models.TeamMember{},
		&models.DataSource{},
		&models.FileToMigrate{},
		&models.Document{},
		&models.MigrateJobs{},
	); err != nil {
		log.Fatal(err)
	}
}
