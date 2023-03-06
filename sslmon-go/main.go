package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SSLInfo struct {
	Version            uint16
	HandshakeComplete  bool
	DidResume         bool
	CipherSuite       uint16
	PeerCertificates  []CertificateInfo
}

type CertificateInfo struct {
	Subject     string
	Issuer      string
	NotBefore   time.Time
	NotAfter    time.Time
	SignatureAlgorithm  string
	PublicKeyAlgorithm string
}

func main() {
	// URL of the SSL website to monitor
	url := "https://www.google.com"

	// Create a new client with a timeout of 5 seconds
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Make a GET request to the URL
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Create a new instance of SSLInfo
	sslInfo := SSLInfo{
		Version: resp.TLS.Version,
		HandshakeComplete: resp.TLS.HandshakeComplete,
		DidResume: resp.TLS.DidResume,
		CipherSuite: resp.TLS.CipherSuite,
	}

	// Retrieve information about the peer certificates
	for _, cert := range resp.TLS.PeerCertificates {
		peerCertificate := CertificateInfo{
			Subject: cert.Subject.String(),
			Issuer: cert.Issuer.String(),
			NotBefore: cert.NotBefore,
			NotAfter: cert.NotAfter,
			SignatureAlgorithm: cert.SignatureAlgorithm.String(),
			PublicKeyAlgorithm: cert.PublicKeyAlgorithm.String(),
		}
		sslInfo.PeerCertificates = append(sslInfo.PeerCertificates, peerCertificate)
	}

	// Marshal the SSLInfo struct to JSON
	sslInfoJSON, err := json.Marshal(sslInfo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON representation of SSLInfo
	fmt.Println(string(sslInfoJSON))
}


