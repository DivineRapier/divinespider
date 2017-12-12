package spider

import (
	"flag"
	"net/url"
	"sync"

	"github.com/divinerapier/divinespider/download"
	"github.com/divinerapier/divinespider/storage"
)

var defaultSpiderClient = &Spider{}

func init() {
	flag.Var(&defaultSpiderClient.uri, "url", "url to crawl")
	flag.IntVar(&defaultSpiderClient.timeout, "timeout", 3, "timeout")
	flag.IntVar(&defaultSpiderClient.thread, "thread", 10, "thread count")
	flag.Parse()
}

// URL rename for url.URL to use it freely
type URL struct {
	*url.URL
}

func (u *URL) String() string {
	return u.URL.String()
}

// Set set value
func (u *URL) Set(value string) error {
	uri, err := url.Parse(value)
	if err != nil {
		return err
	}
	u.URL = uri
	return nil
}

// Spider 爬虫由 下载器, 调度器, 组成
type Spider struct {
	// Downloader 下载器
	download.Downloader

	// Storage 存储器
	storage.Storage

	timeout int
	uri     URL
	thread  int
	sync.WaitGroup
}

// New create
func New() *Spider {
	return defaultSpiderClient
}

// WithDownloader set downloader
func (spider *Spider) WithDownloader(d download.Downloader) *Spider {
	spider.Downloader = d
	return spider
}

// WithStorage set storage
func (spider *Spider) WithStorage(s storage.Storage) *Spider {
	spider.Storage = s
	return spider
}

// Start start task
func (spider *Spider) Start() {
	if len(spider.uri.Host) == 0 {
		panic("url was not set")
	}

	if spider.Downloader == nil {
		panic("downloader was not set")
	}

	if spider.Storage == nil {
		panic("storage was not set")
	}

	spider.SaveTemporarilyURL(spider.uri.String())

	spider.runParallel()

	spider.Wait()
}
