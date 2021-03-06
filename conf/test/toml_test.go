package test

import (
	"fmt"
	"github.com/typeck/frame/conf"
	"log"
	"testing"
	"time"
)

type db struct {
	UserName 	string		`toml:"user_name"`
	Password 	string
}

func TestToml(t *testing.T) {
	config, err := conf.NewFromFile("C:\\golib\\frame\\conf\\test\\conf.toml",true)
	if err != nil {
		log.Fatal(err)
	}
	var db = db{}
	err = config.Unmarshal("postgres.db1", &db)
	fmt.Println(db, err)
	fmt.Println(config.GetStr("postgres.db1"))
	fmt.Println(config.GetInt("postgres.db1"))
	fmt.Println(config.GetStr("postgres.db1.num"))
	fmt.Println(config.GetInt("postgres.db1.num"))

	for {
		time.Sleep(time.Second)
		fmt.Println(config.Get("postgres.db1.user_name"))
	}
}
