package download

// Downloader 下载器
type Downloader interface {
	Crawl(string) [][]byte
}
