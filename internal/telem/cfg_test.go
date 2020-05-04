package telem

import (
	env "git.dsg.tuwien.ac.at/mc2/go-telemc/internal/env"
	"testing"
)

func TestNewDefaultApplicationConfig(t *testing.T) {
	cfg := NewDefaultApplicationConfig()

	if cfg.Redis.URL != "redis://localhost" {
		t.Error("Unexpected default redis URL")
	}
}

func TestApplicationConfig_ReadFromEnvironment(t *testing.T) {
	cfg := NewDefaultApplicationConfig()
	e := env.OsEnv

	e.Set("telemc_redis_host", "192.168.99.1")
	e.Set("telemc_redis_port", "1234")

	cfg.LoadFromEnvironment(env.OsEnv)

	if cfg.Redis.URL != "redis://192.168.99.1:1234" {
		t.Error("Expected url to be redis://192.168.99.1:1234, but was", cfg.Redis.URL)
	}
}
