package metricsandprofiling

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "net/http/pprof"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

var notificationsCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "notifications_request_count",
		Help: "No of request handled by Notifications handler",
	},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
}

func main() {
	prometheus.MustRegister(pingCounter)
	fmt.Println("http server started")

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/notifications", notifications)
	http.HandleFunc("/mail", mail)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8094", nil)
}

func mail(w http.ResponseWriter, req *http.Request) {
	go sendMail()
	fmt.Fprintf(w, "mail")
}

func notifications(w http.ResponseWriter, req *http.Request) {
	notificationsCounter.Inc()

	go listenAndServeNotifications()

	fmt.Fprintf(w, "notifications")
}

func listenAndServeNotifications() {
	for {
		time.Sleep(100 * time.Millisecond)
	}
}

func sendMail() {
	done := make(chan bool)
	<-done
}
