package core

import (
	"github.com/buger/jsonparser"
	"github.com/divinerapier/divinespider/download"
)

// NewDownload 创建 Phantom 下载器
func NewDownload(args []byte) download.Download {
	// 1. 解析 URL
	u, _ := jsonparser.GetString(args, "url")
	if len(u) > 7 && u[:7] != "http://" && u[:8] != "https://" {
		u = "http://" + u
	}

	header, err := jsonparser.GetString(args, "option", "header")
	if err != nil {
		header = ""
	}
	cookie, err := jsonparser.GetString(args, "option", "cookie")
	if err != nil {
		header = ""
	}
	customCookie := ""
	if cookie != "" {
		customCookie = setCookie(cookie)
	}

	phantomPath, _ := jsonparser.GetString(args, "phantomjs", "path")
	phantomOpts, _ := jsonparser.GetString(args, "phantomjs", "args")
	phantomJS, _ := jsonparser.GetString(args, "phantomjs", "script")

	return &download.Phantom{
		Header: "",
		Cookie: "",
		Script: "",
		Option: "",
		Path:   "",
	}
}
