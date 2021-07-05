package main

//noinspection GoUnresolvedReference
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
	mysqlAdder := "localhost:3306"
	addr := ":9999"
	dsn := "root:123456@tcp(localhost:3306)/mysql?charset=utf8mb4&loc=PRC&parseTime=true"

	//连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logrus.Fatal(err)
	}
	prometheus.MustRegister(collectors.NewUpCollector(db))
	prometheus.MustRegister(collectors.NewTrafficCollector(db))
	prometheus.MustRegister(prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name:        "mysql_info",
		Help:        "mysql info",
		ConstLabels: prometheus.Labels{"addr": mysqlAdder},
	}, func() float64 {
		return 1
	}))
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println()
	http.ListenAndServe(addr, nil)
}
