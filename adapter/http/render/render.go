package render

import (
	"github.com/unrolled/render"
)

type Config struct {
	IsDevelopment bool `yaml:"is_development"`
}

var defaultCfg = &Config{
	IsDevelopment: false,
}

func New(cfg *Config) *render.Render {
	c := defaultCfg
	if cfg != nil {
		c = cfg
	}

	return render.New(render.Options{
		Charset:       "UTF-8",
		Extensions:    []string{".html"},
		IsDevelopment: c.IsDevelopment,
	})
}
