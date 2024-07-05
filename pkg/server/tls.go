package server

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func CreateCaPool(caCertPath string) (pool *x509.CertPool, err error)  {
	pool = x509.NewCertPool()
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		return nil, err
	}
	pool.AppendCertsFromPEM(caCrt)
	return
}

func CreateTlsConfig(serverKeyPath, serverCertPath string, pool *x509.CertPool) (tlsConfig *tls.Config, err error) {
	tlsConfig = &tls.Config{
		ClientCAs:  pool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		KeyLogWriter: nil,
	}
	if tlsCertKey, err := tls.LoadX509KeyPair(serverCertPath, serverKeyPath); err != nil {
		return nil, err
	} else {
		tlsConfig.Certificates = []tls.Certificate{tlsCertKey}
	}
	return
}
