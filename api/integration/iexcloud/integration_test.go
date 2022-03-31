package iexcloud

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetStock(t *testing.T) {
	c := NewIexCloudClient(viper.GetString("IEXCLOUD_API_KEY"))

	res, err := c.GetStock("aapl")

	assert.NotNil(t, res.CompanyName, "expecting non-nil res")
	assert.NotNil(t, err, "expecting nil err")

	assert.Equal(t, "Apple Inc", res.CompanyName, "expecting company name to be 'Apple Inc'")

}
