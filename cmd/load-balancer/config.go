package main

import (
	"errors"
	"net/url"
	"time"
)

// Config for the proxy
type Config struct {
	Name             string            `default:"load-balancer" desc:"Name of Network Service Endpoint"`
	ServiceName      string            `default:"load-balancer" desc:"Name of providing service" split_words:"true"`
	Labels           map[string]string `default:"" desc:"Endpoint labels"`
	ConnectTo        url.URL           `default:"unix:///var/lib/networkservicemesh/nsm.io.sock" desc:"url to connect to NSM" split_words:"true"`
	DialTimeout      time.Duration     `default:"5s" desc:"timeout to dial NSMgr" split_words:"true"`
	RequestTimeout   time.Duration     `default:"15s" desc:"timeout to request NSE" split_words:"true"`
	MaxTokenLifetime time.Duration     `default:"24h" desc:"maximum lifetime of tokens" split_words:"true"`
	NSPService       string            `default:"nsp-service:7778" desc:"IP (or domain) and port of the NSP Service" split_words:"true"`
	Namespace        string            `default:"default" desc:"Namespace the pod is running on" split_words:"true"`
	ConfigMapName    string            `default:"meridio-configuration" desc:"Name of the ConfigMap containing the configuration" split_words:"true"`
}

// IsValid checks if the configuration is valid
func (c *Config) IsValid() error {
	if c.ConnectTo.String() == "" {
		return errors.New("no NSMGr ConnectTO URL are specified")
	}
	return nil
}
