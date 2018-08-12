package config_test

import (
	"testing"

	"github.com/soulskit/config"
	// "github.com/wxcsdb88/gin-quick/config"
)

func Test_LoadConfig(t *testing.T) {
	testfile := "./app_test.toml"

	var conf map[string]interface{}

	err := config.LoadConfig(testfile, &conf)
	if err != nil {
		t.Error(err)
	}
	// conf := config.LoadConfig(testfile)
	t.Logf("config is: %#v", conf)
}
