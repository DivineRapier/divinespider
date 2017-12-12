/*
 * @Author: divinerapier
 * @Date: 2017-11-18 22:36:19
 * @Last Modified by: divinerapier
 * @Last Modified time: 2017-11-18 23:04:52
 */

package download

// Downloader 下载器
type Downloader interface {
	Crawl(string) [][]byte
}
