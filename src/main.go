package main

import (
	"Proxy/proxyHandlers"
	"Proxy/repeatReqHandlers"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
}

func main() {
	proxyPort := "8080"
	if port := os.Getenv("PROXY_PORT"); len(port) != 0 {
		proxyPort = port
	}
	dashboardPort := "80"
	if port := os.Getenv("DASHBOARD_PORT"); len(port) != 0 {
		dashboardPort = port
	}

	server := &http.Server{
		Addr:    ":" + proxyPort,
		Handler: http.HandlerFunc(proxyHandlers.ServeHttp),
	}

	http.HandleFunc("/", repeatReqHandlers.SendRequestList)
	http.HandleFunc("/req", repeatReqHandlers.ExecRepReq)
	http.HandleFunc("/scan", repeatReqHandlers.TryVulnerabilities)

	logrus.Info("Proxy server started on port ", proxyPort)
	logrus.Info("Repeat server started on port ", dashboardPort)

	go http.ListenAndServe(":"+dashboardPort, nil)
	logrus.Fatal(server.ListenAndServe())
}
