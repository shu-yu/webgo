package webgo

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

// Config is used for reading app's configuration from json file
type Config struct {
	// Host is the host on which the server is listening
	Host string `json:"host,omitempty"`
	// Port is the port number where the server has to listen for the HTTP requests
	Port string `json:"port,omitempty"`

	// CertFile is the TLS/SSL certificate file path, required for HTTPS
	CertFile string `json:"certFile,omitempty"`
	// KeyFile is the filepath of private key of the certificate
	KeyFile string `json:"keyFile,omitempty"`
	// HTTPSPort is the port number where the server has to listen for the HTTP requests
	HTTPSPort string `json:"httpsPort,omitempty"`

	// readTimeout is the maximum duration for which the server would read a request
	ReadTimeout time.Duration `json:"readTimeout,omitempty"`
	// writeTimeout is the maximum duration for which the server would try to respond
	WriteTimeout time.Duration `json:"writeTimeout,omitempty"`

	// InsecureSkipVerify is the HTTP certificate verification
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
}

// Load config file from the provided filepath and validate
func (cfg *Config) Load(filepath string) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		errLogger.Fatal(err)
	}

	err = json.Unmarshal(file, cfg)
	if err != nil {
		errLogger.Fatal(err)
	}

	cfg.Validate()
}

// Validate the config parsed into the Config struct
func (cfg *Config) Validate() {
	i, err := strconv.Atoi(cfg.Port)
	if err != nil {
		errLogger.Fatal(ErrInvalidPort)
	}
	if i <= 0 || i > 65535 {
		errLogger.Fatal(ErrInvalidPort)
	}
}
