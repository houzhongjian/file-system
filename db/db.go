package db

import (
	mig "github.com/houzhongjian/postgresql-migration"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `json:"host"`
	Dbname   string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
	Sslmode  string `json:"ssl"`
}

var Db *sql.DB

func InitDbs() {
	cf := Config{
		Host:     "127.0.0.1",
		Dbname:   "file",
		User:     "postgres",
		Password: "Cp0422$()",
		Sslmode:  "disable",
	}
	if err := ConnentDb(&cf); err != nil {
		log.Printf("%+v\n", err)
		return
	}

	srv := mig.MigServ{
		Db: Db,
	}
	srv.Migration(
		&User{},
		&Account{},
	)
}

//连接数据库.
func ConnentDb(cf *Config) error {
	args := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", cf.Host, cf.User, cf.Password, cf.Dbname, cf.Sslmode)
	db, err := sql.Open("postgres", args)
	if err != nil {
		log.Printf("%+v\n", err)
		return err
	}
	Db = db
	return nil
}
