package main

import (
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <host:port>", os.Args[0])
	}

	conn, err := tls.Dial("tcp", os.Args[1], &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	cstate := conn.ConnectionState()
	conn.Close()

	switch cstate.Version {
	case tls.VersionTLS12:
		fmt.Print("TLS1.2")
	case tls.VersionTLS11:
		fmt.Print("TLS1.1")
	case tls.VersionTLS10:
		fmt.Print("TLS1.0")
	case tls.VersionSSL30:
		fmt.Print("SSL3.0")
	default:
		panic("what tls version")
	}
	fmt.Print("/")

	switch cstate.CipherSuite {
	case tls.TLS_RSA_WITH_RC4_128_SHA:
		fmt.Println("TLS_RSA_WITH_RC4_128_SHA")
	case tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA:
		fmt.Println("TLS_RSA_WITH_3DES_EDE_CBC_SHA")
	case tls.TLS_RSA_WITH_AES_128_CBC_SHA:
		fmt.Println("TLS_RSA_WITH_AES_128_CBC_SHA")
	case tls.TLS_RSA_WITH_AES_256_CBC_SHA:
		fmt.Println("TLS_RSA_WITH_AES_256_CBC_SHA")
	case tls.TLS_RSA_WITH_AES_128_CBC_SHA256:
		fmt.Println("TLS_RSA_WITH_AES_128_CBC_SHA256")
	case tls.TLS_RSA_WITH_AES_128_GCM_SHA256:
		fmt.Println("TLS_RSA_WITH_AES_128_GCM_SHA256")
	case tls.TLS_RSA_WITH_AES_256_GCM_SHA384:
		fmt.Println("TLS_RSA_WITH_AES_256_GCM_SHA384")
	case tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA:
		fmt.Println("TLS_ECDHE_ECDSA_WITH_RC4_128_SHA")
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:
		fmt.Println("TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA")
	case tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:
		fmt.Println("TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA")
	case tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA:
		fmt.Println("TLS_ECDHE_RSA_WITH_RC4_128_SHA")
	case tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA:
		fmt.Println("TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA")
	case tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:
		fmt.Println("TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA")
	case tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:
		fmt.Println("TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA")
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256:
		fmt.Println("TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256")
	case tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256:
		fmt.Println("TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256")
	case tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:
		fmt.Println("TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256")
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:
		fmt.Println("TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256")
	case tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:
		fmt.Println("TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384")
	case tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:
		fmt.Println("TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384")
	case tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305:
		fmt.Println("TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305")
	case tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305:
		fmt.Println("TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305")
	}

	fmt.Println()

	for _, v := range cstate.PeerCertificates {
		fmt.Printf("Subject: %s\n", v.Subject.CommonName)
		fmt.Printf("Issuer: %s\n", v.Issuer.CommonName)
		fmt.Printf("SubjectKeyID: %s\n", strings.ToUpper(hex.EncodeToString(v.SubjectKeyId)))
		fmt.Printf("AuthorityKeyID: %s\n", strings.ToUpper(hex.EncodeToString(v.AuthorityKeyId)))
		fmt.Printf("DNS Names: %+v\n", v.DNSNames)
		fmt.Printf("IP Addresses: %+v\n", v.IPAddresses)
		fmt.Println()
	}
}
