package config

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/StefanWellhoner/task-manager-api/internal/mode"
	"github.com/stretchr/testify/assert"
)

// TestGet is a unit test function that tests the Get function in the config package.
// It verifies the correctness of the configuration values returned by the Get function.
func TestGet(t *testing.T) {
	// Store the current directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	// Determine the project root directory
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), "../../")

	mode.SetEnv(mode.Test)

	// Change to project root directory
	err = os.Chdir(dir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

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

	// Change back to the original directory
	err = os.Chdir(originalDir)
	if err != nil {
		t.Fatalf("Failed to change back to original directory: %v", err)
	}
}

// TestConfigFiles tests the behavior of the configFiles function.
// It verifies that the correct configuration file is returned based on the provided mode.
func TestConfigFiles(t *testing.T) {
	testCases := []struct {
		mode        string
		expectedCfg string
	}{
		{mode.Prod, "config.yml"},
		{mode.Dev, "config.dev.yml"},
		{mode.Test, "config.test.yml"},
		{"unknown", "config.dev.yml"},
	}

	for _, tc := range testCases {
		mode.SetEnv(tc.mode)
		cfgFile := configFiles()
		assert.Equal(t, tc.expectedCfg, cfgFile, "Config file should be %s", tc.expectedCfg)
	}
}
