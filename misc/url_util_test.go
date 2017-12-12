package misc

import (
	"extra/testify/assert"
	"testing"
)

func TestGetURLHost(t *testing.T) {
	url, host := "http://www.baidu.com", "www.baidu.com"
	assert.Equal(t, host, GetURLHost(url), "")
	url, host = "https://www.baidu.com", "www.baidu.com"
	assert.Equal(t, host, GetURLHost(url), "")
	url, host = "http://www.baidu.com/", "www.baidu.com"
	assert.Equal(t, host, GetURLHost(url), "")
	url, host = "https://www.baidu.com/", "www.baidu.com"
	assert.Equal(t, host, GetURLHost(url), "")
	url, host = "http://www.baidu.com/search?", "www.baidu.com"
	assert.Equal(t, host, GetURLHost(url), "")
	url, host = "https://www.baidu.com/search?", "www.baidu.com"
	assert.Equal(t, host, GetURLHost(url), "")
	url, host = "http://www.baidu.com/search?a=golang", "www.baidu.com"
	assert.Equal(t, host, GetURLHost(url), "")
	url, host = "https://www.baidu.com/search?a=golang", "www.baidu.com"
	assert.Equal(t, host, GetURLHost(url), "")
}
