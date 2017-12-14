package core

import (
	"fmt"
	"net/url"
	"sort"
	"sync"

	"bytes"
	"os"
	"runtime"
	"strings"
	"time"

	"encoding/json"

	"github.com/DivineRapier/go-tools/urltool"
	"github.com/buger/jsonparser"
	"github.com/divinerapier/divinespider/download"
	"qiniupkg.com/x/log.v7"
)

// Spider type
type (
	Spider struct {
		download.Download
		result     *sync.Map // 抓取结果
		recordChan chan record
		hostInfo   string // host --> host:port -> uri.Host
		spiFlag    int    // 爬虫模式
		spiEntry   string // 爬虫入口
	}

	cookieType struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
)

// NewSpider 创建 Spider 对象
// 接受 json 作为参数
func NewSpider() *Spider {

	if !strings.HasPrefix(crawlURL, "http://") && !strings.HasPrefix(crawlURL, "https://") {
		crawlURL = "http://" + crawlURL
	}

	u, err := url.Parse(crawlURL)
	if err != nil {
		panic(err)
	}

	phantom := &download.Phantom{
		Header: header,
		Cookie: cookie,
		Script: path,
		Option: option,
		Path:   path,
	}
	spider := &Spider{
		Download:   phantom,
		result:     &sync.Map{},
		recordChan: make(chan record, 10240),
		hostInfo:   u.Host,
		spiFlag:    mode,
		spiEntry:   crawlURL,
	}

	spider.pushURL(crawlURL)

	return spider
}

func (s *Spider) pushURL(u string) {
	if s == nil || s.recordChan == nil {
		panic("please initialize spider")
	}
	s.recordChan <- NewRecord("request", "GET", u, "")
}

func (s *Spider) pushRecord(r record) {
	if s == nil || s.recordChan == nil {
		panic("please initialize spider")
	}
	s.recordChan <- r
}

// Exec 执行爬虫
func (s *Spider) Execute(args ...string) {
	// 执行 phantomjs 程序 获取链接
	go checkFinished()
	go s.queue_select()
	go func() {
		for {
			log.Printf("num of current go routines: %d\n", runtime.NumGoroutine())
			select {
			case rec, _ := <-execQueue:
				{
					for _, v := range rec {
						maxWaitRoutine <- true
						go func() {
							maxExecRoutine <- true
							s.exec(v.Data.URL)
							<-maxExecRoutine
							<-maxWaitRoutine
						}()
					}
				}
			case <-time.After(60 * time.Second):
				{
					log.Fatalf("超时!\n")
					os.Exit(-1)
				}
			}
		}
	}()

	<-finish
	fmt.Println("done!!!")
	fmt.Printf("%d URLs discovered!!!\n", s.result.Size())
	rets := make(Records, 0)
	for item := range s.result.IterItems() {
		rets = append(rets, item.Value.(Record))
	}
	sort.Sort(rets)
	for _, v := range rets {
		r, err := json.Marshal(v)
		if err != nil {
			log.Errorf("%v\n", err.Error())
			continue
		}
		s := strings.Replace(string(r), `\u0026`, `&`, -1)
		fmt.Printf("%s\n", s)
	}
}

func (spi *Spider) exec(u string) {

	out := spi.Run(u)

	lines := bytes.Split(out, []byte{'\n'})
	var (
		finalURL, typ, mtd, data, hash string
		ok                             bool
	)
	for _, line := range lines {

		if len(line) < 11 {
			continue
		}
		log.Printf("line: %s\n", line)
		// redirect 与 type 只能存在一个
		if bytes.Contains(line, []byte{'"', 'r', 'e', 'd', 'i', 'r', 'e', 'c', 't'}) {
			redirectURL, err := jsonparser.GetString(line, "redirect")
			if err == nil && strings.Trim(redirectURL, " ") != "" {
				finalURL = redirectURL
			}
			typ = "link"
			mtd = "GET"
			data = ""
		} else if bytes.Contains(line, []byte{'{', '"', 't', 'y', 'p', 'e'}) {
			// 2. 包含 type

			finalURL, _ = jsonparser.GetString(line, "data", "url")

			if finalURL[len(finalURL)-1] == '?' {
				finalURL = finalURL[:len(finalURL)-1]
			}
			typ, _ = jsonparser.GetString(line, "type")
			mtd, _ = jsonparser.GetString(line, "data", "method")
			data, _ = jsonparser.GetString(line, "data", "data")
			if typ == "" || mtd == "" {
				continue
			}
		}

		hash, ok = hash_url(finalURL, mtd, typ, data)
		if !ok {
			continue
		}

		if !spi.IsSameDomain(finalURL) {
			continue
		}

		r := NewRecord(typ, mtd, finalURL, data)
		a := map[string]Record{
			hash: r,
		}
		waitQueue <- a
		deepth, err := urltool.URLDeepthByString(finalURL)
		if err != nil {
			return
		}
		for i := 1; i < deepth; i++ {
			path, err := urltool.URLSpecificNodeByString(finalURL, i)
			if err != nil {
				continue
			}

			m := NewRecord("request", "GET", spi.hostInfo+path, "")

			hash, ok = hash_url(m.Data.URL, m.Method, m.Type, m.Data.Data)
			h := map[string]Record{
				hash: m,
			}
			waitQueue <- h
		}

	}
}

// IsSameDomain 判断是否为同一个 Host
func (spi *Spider) IsSameDomain(u string) bool {
	uri, err := url.Parse(u)
	return err == nil && (uri.Scheme+"://"+uri.Host) == spi.hostInfo
}

func checkFinished() {
	time.Sleep(60 * time.Second)
	for {
		if runtime.NumGoroutine() == 4 {
			finish <- true
		} else {
			time.Sleep(5 * time.Second)
		}
	}
}
