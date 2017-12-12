package chanstorage

import (
	"sync"

	"github.com/divinerapier/divinespider/storage"
)

// ChanStorage channel 作为存储数据结构
type ChanStorage struct {
	tempResultChan   chan []byte
	crawledResultMap sync.Map
}

// New create a storage adaptor
func New() storage.Storage {
	return new(ChanStorage)
}

// Peek 选择一个
func (c *ChanStorage) Peek() <-chan []byte {
	return c.tempResultChan
}

// SaveTemporarilyURL 添加URL
func (c *ChanStorage) SaveTemporarilyURL(url string) {
	c.tempResultChan <- []byte(url)
}

// SaveTemporarilyResults 添加爬虫结果
func (c *ChanStorage) SaveTemporarilyResults(results [][]byte) {
	for i := range results {
		c.tempResultChan <- results[i]
	}
}

// PersistURL 保存 URL
func (c *ChanStorage) PersistURL(url string) {
	c.crawledResultMap.Store(url, struct{}{})
}
