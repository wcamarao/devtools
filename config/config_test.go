package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcamarao/devtools/config"
)

var c = config.GetConfig()

func TestNewConfig(t *testing.T) {
	assert.Equal(t, c.DB.Host, "localhost")
	assert.Equal(t, c.DB.User, "codelab")
	assert.Equal(t, c.DB.Pass, "codelab")
	assert.Equal(t, c.DB.Name, "devtools")
	assert.Equal(t, c.DB.SSLMode, "disable")
}
