package config_test

import (
	"testing"

	"github.com/wxcsdb88/gin-quick/config"
)

func Test_LoadConfig(t *testing.T) {
	testfile := "./app_test.toml"
	conf := config.LoadConfig(testfile)
	t.Logf("config is: %#v", conf)
}
