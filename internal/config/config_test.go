package config

import (
	"testing"

	"github.com/StefanWellhoner/task-manager-api/internal/mode"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	mode.SetEnv(mode.Test)

	conf := Get()
	// Test Server Settings
	assert.Equal(t, 8080, conf.Server.Port, "Port should be 8080 (default)")
	assert.Equal(t, "localhost", conf.Server.ListenAddr, "ListenAddr should be localhost")
	assert.Equal(t, 0, conf.Server.KeepAlivePeriodSecond, "KeepAlivePeriodSecond should be 0")

	// Test SSL Settings
	assert.Equal(t, false, conf.Server.SSL.Enabled, "SSL should be disabled")
	assert.Equal(t, false, conf.Server.SSL.RedirectToHTTPS, "RedirectToHTTPS should be disabled")
	assert.Equal(t, "", conf.Server.SSL.ListenAddr, "SSL ListenAddr should be empty")
	assert.Equal(t, 443, conf.Server.SSL.Port, "SSL Port should be 443")
	assert.Equal(t, "", conf.Server.SSL.CertFile, "CertFile should be empty")
	assert.Equal(t, "", conf.Server.SSL.KeyFile, "KeyFile should be empty")

	// Test LetsEncrypt Settings
	assert.Equal(t, false, conf.Server.SSL.LetsEncrypt.Enabled, "LetsEncrypt should be disabled")
	assert.Equal(t, []string{"test.com", "test.co.za"}, conf.Server.SSL.LetsEncrypt.Hosts, "LetsEncrypt Hosts should be test.com and test.co.za")
	assert.Equal(t, true, conf.Server.SSL.LetsEncrypt.AcceptTOS, "LetsEncrypt AcceptTOS should be false")
	assert.Equal(t, "data/certs", conf.Server.SSL.LetsEncrypt.CacheDir, "LetsEncrypt CacheDir should be data/certs")

	// Test Database Settings
	assert.Equal(t, "localhost", conf.Database.Host, "Database Host should be localhost")
	assert.Equal(t, 5432, conf.Database.Port, "Database Port should be 5432")
	assert.Equal(t, "test_user", conf.Database.User, "Database User should be postgres")
	assert.Equal(t, "test_password", conf.Database.Password, "Database Password should be postgres")
	assert.Equal(t, "test_db", conf.Database.Database, "Database Database should be postgres")
	assert.Equal(t, "disable", conf.Database.SSLMode, "Database SSLMode should be disable")
}
