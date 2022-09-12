// This is custom goose binary with sqlite3 support only.

package migrations

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/pressly/goose/v3"
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

func Migration() {
	var db *sql.DB
	// setup database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, _ = sql.Open("mysql", dsn)

	if err := db.Ping(); err != nil {
		log.Fatalln(string("\033[31m"), "error connection: ", err.Error())
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	s, _ := goose.GetDBVersion(db)
	fmt.Println("version of db", s)

	if err := goose.Run("up", db, "migrations"); err != nil {
		panic(err)
	}
}
