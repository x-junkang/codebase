package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var dsn string

func init() {
	flag.StringVar(&dsn, "dsn", "", "mysql dsn")
	flag.Parse()
	fmt.Println(dsn)
}

func createEngine(dsn string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, err
	}

	engine.SetMaxIdleConns(20)
	engine.SetMaxOpenConns(200)
	engine.SetConnMaxLifetime(time.Minute * 30)

	// engine.SetLogger()
	engine.ShowSQL(true)
	if true {
		engine.ShowSQL(true)
	}
	return engine, nil
}

func main() {
	engine, err := createEngine(dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	datas := make([]Data, 0, 1000)
	engine.Where("date = ?", "2022-04-30 00:00:00").Find(&datas)
	dataBytes, err := json.Marshal(datas)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("tmp2.json", dataBytes, 0666)
	if err != nil {
		panic(err)
	}
}
