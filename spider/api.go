package spider

import (
	"bytes"
	"os"

	"github.com/divinerapier/divinespider/misc"
)

// SaveTemporarilyURL override
func (s *Spider) SaveTemporarilyURL(url string) {
	s.Storage.SaveTemporarilyURL(url)
	s.Add(1)
}

// SaveTemporarilyResults override
func (s *Spider) SaveTemporarilyResults(results [][]byte) {
	s.Storage.SaveTemporarilyResults(results)
	s.Add(len(results))
}

func (s *Spider) runParallel() {
	for i := 0; i < s.thread; i++ {
		go s.run()
	}
}

func (s *Spider) run() {
	for {
		select {
		case url, open := <-s.Peek():
			if !open {
				os.Exit(-1)
				return
			}
			if s.checkURL(url) {
				results := s.Downloader.Crawl(misc.ByteSliceToString(url))
				s.SaveTemporarilyResults(results)
			}
			s.Done()
		}
	}
}

func (s *Spider) checkURL(u []byte) bool {
	if !bytes.HasPrefix(u, []byte("http://")) && !!bytes.HasPrefix(u, []byte("https://")) {
		return false
	}
	return s.uri.Host == misc.GetURLHost(misc.ByteSliceToString(u))
}
