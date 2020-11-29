package consumer

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"strings"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetString(string) string
	GetStringSlice(string) []string
	GetInt(string) int
	GetInt32(string) int32
	GetBool(string) bool
}

// ConsumerConfig ..
type ConsumerConfig struct {
	Version   string
	ClientID  string
	Topic     string
	Brokers   []string
	Partition int32

	SASL ConsumerSASLConfig

	TLS ConsumerTLSConfig

	CallbackError   ConsumerHandlerError
	CallbackSuccess ConsumerHandlerSuccess
}

// RawConsumerConfig ..
func RawConsumerConfig(confPrefix string, confHandler ConfigHandler) *ConsumerConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &ConsumerConfig{
		Version:   confHandler.GetString(confPrefix + ".version"),
		ClientID:  confHandler.GetString(confPrefix + ".client_id"),
		Topic:     confHandler.GetString(confPrefix + ".topic"),
		Partition: confHandler.GetInt32(confPrefix + ".partition"),
		Brokers:   confHandler.GetStringSlice(confPrefix + ".brokers"),
		SASL: ConsumerSASLConfig{
			Enable:   confHandler.GetBool(confPrefix + ".sasl_enable"),
			Username: confHandler.GetString(confPrefix + ".sasl_username"),
			Password: confHandler.GetString(confPrefix + ".sasl_password"),
		},
		TLS: ConsumerTLSConfig{
			Enable:   confHandler.GetBool(confPrefix + ".tls_enable"),
			Verify:   confHandler.GetBool(confPrefix + ".tls_verify"),
			CertFile: confHandler.GetString(confPrefix + ".tls_cert_file"),
			KeyFile:  confHandler.GetString(confPrefix + ".tls_key_file"),
			CaFile:   confHandler.GetString(confPrefix + ".tls_ca_file"),
		},
	}
}

// WithCallbackError set error callback.
//
// triggered every time when a error message is reveived.
func (t *ConsumerConfig) WithCallbackError(handler ConsumerHandlerError) *ConsumerConfig {
	t.CallbackError = handler
	return t
}

// WithCallbackSuccess set success callback.
//
// triggered every time when a success message is reveived.
func (t *ConsumerConfig) WithCallbackSuccess(handler ConsumerHandlerSuccess) *ConsumerConfig {
	t.CallbackSuccess = handler
	return t
}

// Build ..
func (t *ConsumerConfig) Build() (*Consumer, error) {
	return newConsumer(t)
}

type ConsumerSASLConfig struct {
	Enable   bool
	Username string
	Password string
}

type ConsumerTLSConfig struct {
	Enable   bool
	Verify   bool
	CertFile string
	KeyFile  string
	CaFile   string
}

// convert to tls.Config
func (t *ConsumerTLSConfig) tlsConfig() (*tls.Config, error) {
	conf := &tls.Config{
		InsecureSkipVerify: t.Verify,
	}

	if len(t.CertFile) > 0 && len(t.KeyFile) > 0 && len(t.CaFile) > 0 {
		cert, err := tls.LoadX509KeyPair(t.CertFile, t.KeyFile)
		if err != nil {
			return nil, err
		}

		caCert, err := ioutil.ReadFile(t.CaFile)
		if err != nil {
			return nil, err
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		conf = &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
			InsecureSkipVerify: t.Verify,
		}
	}

	return conf, nil
}
