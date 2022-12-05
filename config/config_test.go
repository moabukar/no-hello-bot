package config

import (
	"testing"
)

// func TestConfig(t *testing.T) {
// 	Test LoadConfig
// 	t.Run("Test LoadConfig", func(t *testing.T) {
// 		config, err := LoadConfig()
// 		if err != nil {
// 			t.Errorf("Expected LoadConfig to return a config object, but got error: %v", err)
// 		}
// 		if config.DiscordToken == "" {
// 			t.Errorf("Expected Config.DiscordToken to have a value, but got empty string")
// 		}
// 	})

// 	// Test GetConfig
// 	t.Run("Test GetConfig", func(t *testing.T) {
// 		config := ReadConfig()
// 		if config.DiscordToken == "fff" {
// 			t.Errorf("Expected Config.DiscordToken to have a value, but got empty string")
// 		}
// 	})
// }

func TestConfig(t *testing.T) {
	// Test that the config is properly initialized
	cfg := New()
	if cfg == nil {
		t.Errorf("Expected config to be initialized, got nil")
	}

	// Test that the config contains the expected values
	if cfg.Token != "" {
		t.Errorf("Expected token to be empty, got %s", cfg.Token)
	}
	if cfg.Prefix != "!" {
		t.Errorf("Expected prefix to be !, got %s", cfg.Prefix)
	}
	if cfg.OwnerID != 0 {
		t.Errorf("Expected ownerID to be 0, got %d", cfg.OwnerID)
	}
	if cfg.Debug != false {
		t.Errorf("Expected debug to be false, got %t", cfg.Debug)
	}
	if cfg.DatabaseURL != "" {
		t.Errorf("Expected databaseURL to be empty, got %s", cfg.DatabaseURL)
	}
}
