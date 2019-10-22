package config

import (
	"golang-odai/adapter/http/render"
	"golang-odai/adapter/http/session"
	"golang-odai/external/firebase"
)

type Config struct {
	Domain  string `required:"true"`
	Session *session.Config `required:"true"`
	Render  *render.Config `required:"true"`
	Firebase *firebase.Config `required:"true"`
	Jaeger *Jaeger
}

type Jaeger struct {
	CollectorEndpoint string `yaml:"collector_endpoint"`
	AgentEndpoint string `yaml:"trace:6831"`
	ServiceName string `yaml:"service_name"`
}