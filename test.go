package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
var labels = map[string]string{
	"kubernetes_namespace": "myNamespace",
	"kubernetes_pod_name": "myPodName",
	"app_name":"test_app_ver1.0",
	"stage": "prd",
	//"__meta_kubernetes_pod_node_name": "test-work01",  //以 __ 作为前缀的标签，是系统保留的关键字，只能在系统内部使用
}
var (
	my_metric = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "ci",
		Subsystem: "zhirenyun",
		Name: "my_metric",
		Help: "help info.",
		ConstLabels: labels,
	})
)

func init() {
	prometheus.MustRegister(my_metric)
}

func main() {
	flag.Parse()
	go func() {
		for {
			my_metric.Add(5)
			time.Sleep(time.Second*1)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
