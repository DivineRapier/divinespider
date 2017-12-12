/*
 * @Author: divinerapier
 * @Date: 2017-11-18 22:15:36
 * @Last Modified by: divinerapier
 * @Last Modified time: 2017-11-18 23:06:34
 */

package core

import (
	"encoding/json"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lunny/log"
)

// 获取指定深度的目录或文件
//
func getDirectory(url_ string, level int) []string {
	var (
		ending string
		tmp    string
	)
	// 1. 判断结束符
	if url_[len(url_)-1] == '/' {
		ending = "/"
	} else {
		ending = ""
	}

	// 2. 判读啊 scheme
	if len(url_) < 6 {
		return nil
	}
	if url_[:7] != "http://" && url_[:8] != "https://" {
		url_ = "http://" + url_
	}

	// 记录最后一个字符是否为分隔符

	uri, err := url.Parse(url_)
	if err != nil {
		log.Errorf("%v, 提取 %d 级目录失败!", err.Error(), level)
	}

	if ending == "/" {
		tmp = uri.Path[:len(uri.Path)-1]
	} else {
		tmp = uri.Path
	}
	path := strings.Split(tmp, "/")
	//return path[level]
	return path
}

// GetExecPath 返回可执行文件的绝对路径
func GetExecPath() (string, error) {
	execRelPath, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	execAbsPath, err := filepath.Abs(execRelPath)
	if err != nil {
		return "", err
	}
	return execAbsPath, nil
}

// GetExecDir 返回可执行文件的目录路径
func GetExecDir() (string, error) {
	path, err := GetExecPath()
	if err != nil {
		return "", err
	}
	return filepath.Dir(path), nil
}

func setCookie(cookie string) string {
	var ret []cookieType
	// 解析 cookie 到 cookieType 变量
	cookies := strings.Split(cookie, ";")
	for _, c := range cookies {
		kv := strings.Split(c, "=")
		if len(kv) != 2 {
			continue
		}
		tmp := cookieType{
			Name:  strings.Trim(kv[0], " "),
			Value: strings.Trim(kv[1], " "),
		}
		ret = append(ret, tmp)
	}

	curDir, err := GetExecDir()
	if err != nil {
		log.Fatalf("%v\n", err.Error())
		os.Exit(-1)
	}

	cookiePath := curDir + string(os.PathSeparator) + "cookie_" + strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Int()%100) + ".txt"
	log.Printf("cookiePath: %s\n", cookiePath)
	cookieFile, err := os.Create(cookiePath)
	if err != nil {
		log.Fatalf("创建cookie文件失败!")
		os.Exit(-1)
	}
	input, err := json.Marshal(ret)
	log.Printf("input: %s \n", input)
	if err != nil {
		log.Fatalf("Json序列化失败")
		os.Exit(-1)
	}
	cookieFile.Write(input)
	return cookiePath
}

func hash_url(URL, mtd, data_type, data string) (string, bool) {
	uri, err := url.Parse(URL)
	if err != nil {
		return "", false
	}
	return strings.Join([]string{uri.Scheme + "://" +
		uri.Host + uri.Path,
		strings.ToLower(mtd),
		hash_data(uri.RawQuery)},
		"|"), true
}

func hash_data(data string) string {
	if data == "" {
		return ""
	}
	s := strings.Split(data, "&")
	var ret []string
	for _, v := range s {
		kv := strings.Split(v, "=")
		ret = append(ret, kv[0])
	}
	sort.Strings(ret)
	return strings.Join(ret, "=&")
}
