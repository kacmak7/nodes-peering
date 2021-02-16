package main

import (
	"log"
	"net/http"
)

func getIP(r *http.Request) (string, error) {
    //Get IP from the X-REAL-IP header
    ip := r.Header.Get("X-REAL-IP")
    netIP := net.ParseIP(ip)
    if netIP != nil {
        return ip, nil
    }

    //Get IP from X-FORWARDED-FOR header
    ips := r.Header.Get("X-FORWARDED-FOR")
    splitIps := strings.Split(ips, ",")
    for _, ip := range splitIps {
        netIP := net.ParseIP(ip)
        if netIP != nil {
            return ip, nil
        }
    }

    //Get IP from RemoteAddr
    ip, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return "", err
    }
    netIP = net.ParseIP(ip)
    if netIP != nil {
        return ip, nil
    }
    return "", fmt.Errorf("No valid ip found")
}

func peer(w http.ResponseWriter, r *http.Request) {
	log.Println("Peering: ", getIP(r))
}

func main() {
	http.handleFunc("/", peer)
	if err := http.LisetnAndServe(":8080", nil), err != nil {
		log.Fatal(err)
	}
}

