package common

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func GetServerTlsConfig() *tls.Config {
	// creds, err := credentials.NewServerTLSFromFile("certificate/localhost.pem", "certificate/localhost.key")
	// if err != nil {
	// 	panic(err)
	// }
	// cert, _ := tls.LoadX509KeyPair("config/certificate/servera/localhost.pem", "config/certificate/servera/localhost.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("config/certificate/ca.pem")
	ok := certPool.AppendCertsFromPEM(ca)
	if !ok {
		panic(any("append ca failed"))
	}

	tlsConfig := &tls.Config{
		// Certificates: []tls.Certificate{cert},
		// RootCAs: certPool,
		ClientCAs:  certPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		// ClientAuth: tls.VerifyClientCertIfGiven,
	}
	return tlsConfig
}
