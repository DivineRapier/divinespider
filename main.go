package main

import (
	"github.com/divinerapier/divinespider/download/phantom"
	"github.com/divinerapier/divinespider/spider"
	"github.com/divinerapier/divinespider/storage/chanstorage"
)

func main() {
	spider.New().WithDownloader(phantom.New()).WithStorage(chanstorage.New()).Start()
}
