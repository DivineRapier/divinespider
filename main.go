package main

import (
	"github.com/divinerapier/divinespider/chanstorage"
	"github.com/divinerapier/divinespider/core/spider"
	"github.com/divinerapier/divinespider/phantom"
)

func main() {
	spider.New().WithDownloader(phantom.New()).WithStorage(chanstorage.New()).Start()
}
