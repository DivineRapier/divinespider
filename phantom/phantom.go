package phantom

import (
	"bytes"
	"os/exec"

	"github.com/divinerapier/divinespider/core/download"
)

// Phantom phantomjs client
type Phantom struct {
	Header string `json:"header,omitempty"` // 自定义 Header
	Cookie string `json:"cookie,omitempty"` // 自定义 Cookie
	Path   string `json:"path,omitempty"`   // Phantom 执行路径
	Script string `json:"script,omitempty"` // Phantom 执行脚本
	Option string `json:"option,omitempty"` // Phantom 执行参数
}

func New() download.Downloader {
	return new(Phantom)
}

func (p *Phantom) Crawl(u string) [][]byte {
	var cmd *exec.Cmd

	cmd = exec.Command(p.Path, p.Option, p.Script, u,
		"-c", p.Cookie,
		"-o", p.Header)

	out, _ := cmd.Output()
	return bytes.Split(out, []byte("\n"))
}
