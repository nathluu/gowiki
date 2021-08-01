package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Hostname        string
	SourceIpAddress string
	IpAddress       string
	Time            time.Time
	Endpoint        string
}

func getSourceIpAddress(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func getInterfaceIpv4Addr(interfaceName string) (addr string, err error) {
	var (
		ief      *net.Interface
		addrs    []net.Addr
		ipv4Addr net.IP
	)

	if ief, err = net.InterfaceByName(interfaceName); err != nil { // get interface
		return
	}

	if addrs, err = ief.Addrs(); err != nil { // get addresses
		return
	}

	for _, addr := range addrs { // get ipv4 address
		if ipv4Addr = addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			break
		}
	}

	if ipv4Addr == nil {
		return "", fmt.Errorf("interface %s don't have an ipv4 address", interfaceName)
	}

	return ipv4Addr.String(), nil
}

func getHostname(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	addr, err := getInterfaceIpv4Addr("eth0")
	if err != nil {
		panic(err)
	}

	fmt.Println("Handling!!!")
	m := Message{hostname, getSourceIpAddress(r), addr, time.Now(), "/"}
	json.NewEncoder(w).Encode(m)
}

func getHostnameOidc(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	addr, err := getInterfaceIpv4Addr("eth0")
	if err != nil {
		panic(err)
	}

	fmt.Println("Handling!!!")
	m := Message{hostname, getSourceIpAddress(r), addr, time.Now(), "/oidc"}
	json.NewEncoder(w).Encode(m)
}

func handleRequests() {
	http.HandleFunc("/", getHostname)
	http.HandleFunc("/oidc", getHostnameOidc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
