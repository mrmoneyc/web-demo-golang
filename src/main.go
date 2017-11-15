package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
)

const listenPort = ":9000"

func getIPAndHostname(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()
	ipAddr, _ := net.InterfaceAddrs()

	fmt.Fprintf(w, "Hostname: %v\nIP: %v", hostName, ipAddr)
}

func getNvidiaGPUInfo(w http.ResponseWriter, r *http.Request) {
	binary, lookErr := exec.LookPath("nvidia-smi")
	if lookErr != nil {
		fmt.Fprintf(w, "LookPath Error: %s", lookErr)
	}

	out, err := exec.Command("sh", "-c", binary).Output()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "%s", out)
}

func main() {
	http.HandleFunc("/", getIPAndHostname)
	http.HandleFunc("/gpuinfo", getNvidiaGPUInfo)

	fmt.Printf("Web server is listening on %v...", listenPort)
	err := http.ListenAndServe(listenPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
