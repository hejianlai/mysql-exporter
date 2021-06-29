package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"mysql-exporter/collectors"
	"net/http"
)

func main() {
	addr := ":9999"
	dsn := "root:123456@tcp(localhost:3306)/mysql?charset=utf8mb4&loc=PRC&parseTime=true"

	//
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logrus.Fatal(err)
	}
	prometheus.MustRegister(collectors.NewUpCollector(db))
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println()
	http.ListenAndServe(addr, nil)
}
