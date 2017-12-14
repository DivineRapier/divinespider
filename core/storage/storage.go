package storage

// Storage 结果存储
type Storage interface {
	// Peek 选择一个
	Peek() <-chan []byte

	// SaveTemporarilyURL 添加URL
	SaveTemporarilyURL(string)

	// SaveTemporarilyResults 添加爬虫结果
	SaveTemporarilyResults([][]byte)

	// PersistURL 持久化 URL
	PersistURL(string)
}
