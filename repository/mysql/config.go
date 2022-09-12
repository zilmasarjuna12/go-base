package repository_mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string `envconfig:"MYSQL_HOST" default:"localhost"`
	Port     string `envconfig:"MYSQL_PORT" default:"3307"`
	User     string `envconfig:"MYSQL_USER" default:"payment"`
	Password string `envconfig:"MYSQL_PASSWORD" default:"payment123"`
	Database string `envconfig:"MYSQL_DATABASE" default:"payment"`
}

var (
	cfg *Config = &Config{}
)

func init() {
	err := envconfig.Process("myapp", cfg)
	if err != nil {
		log.Fatalf("Failed to get myapp env %v", err)
	}
}

func NewRepositories() (*gorm.DB, error) {
	// init connection mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
	})
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	log.Printf("INFO: Connected to DB")

	return db, nil
}
