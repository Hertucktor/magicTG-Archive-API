package config

import "testing"

func TestGetConfig(t *testing.T) {
	c, err := GetConfig("internal/pkg/config/test_config.yml")
	if err != nil {
		t.Error("There should not be an error whilst parsing the yaml file")
	}

	if c.DBUser != "testUser" {
		t.Fatalf("Expected: %v, got: %v ", "vaultToken", c.DBUser)
	}
}