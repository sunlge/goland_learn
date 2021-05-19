package main


import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)


func main() {

	// 定义变量，固定lables
	cpu := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "sunlge",
		Help: "cpu total",
		ConstLabels: prometheus.Labels{"sunlge": "test"},
	})

	// 非固定lables
	disk := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "disk",
		Help: "disk surplus amount",
	}, []string{"mount", "disk"})

	// 设置变量值
	cpu.Set(2)
	disk.WithLabelValues("c:", "1").Set(100)
	disk.WithLabelValues("e:", "2").Set(200)
	// 注册
	prometheus.MustRegister(cpu)
	prometheus.MustRegister(disk)

	// 暴漏服务
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9999", nil)


}


