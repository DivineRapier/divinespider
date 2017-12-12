// Spider: 并行爬虫
// Author: https://github.com/DivineRapier/Spider
//         poriter.coco@gmail.com

package core

import "testing"

// TestRelativeToAbsolute ...
func TestRelativeToAbsolute(t *testing.T) {
	/*
		{
			if ret := RelativeToAbsolute("http://baidu.com/a/b", ""); ret != "http://baidu.com/a/b" {
				t.Errorf("ret : %s\n", ret)
			}
			if ret := RelativeToAbsolute("http://baidu.com/a/b", "1"); ret != "http://baidu.com/a/1" {
				t.Errorf("ret : %s\n", ret)
			}
			if ret := RelativeToAbsolute("http://baidu.com/a/b", "../1"); ret != "http://baidu.com/1" {
				t.Errorf("ret : %s\n", ret)
			}
			if ret := RelativeToAbsolute("http://baidu.com/a/b/c/../d", ""); ret != "http://baidu.com/a/b/d" {
				t.Errorf("ret : %s\n", ret)
			}
			if ret := RelativeToAbsolute("http://baidu.com/a/b/c/../d", "1"); ret != "http://baidu.com/a/b/1" {
				t.Errorf("ret : %s\n", ret)
			}
			if ret := RelativeToAbsolute("http://baidu.com/a/b/c/../d", "./1"); ret != "http://baidu.com/a/b/1" {
				t.Errorf("ret : %s\n", ret)
			}
			if ret := RelativeToAbsolute("http://baidu.com/a/b/c/../d", "../1"); ret != "http://baidu.com/a/1" {
				t.Errorf("ret : %s\n", ret)
			}
			if ret := RelativeToAbsolute("http://baidu.com/a/b/c/../d/.", "1"); ret != "http://baidu.com/a/b/d/1" {
				t.Errorf("ret : %s\n", ret)
			}
			if ret := RelativeToAbsolute("http://baidu.com/a/b/c/../../d/.", "1"); ret != "http://baidu.com/a/d/1" {
				t.Errorf("ret : %s\n", ret)
			}
			if ret := RelativeToAbsolute("http://baidu.com/a/b/c/../../d/.", "/1"); ret != "http://baidu.com/1" {
				t.Errorf("ret : %s\n", ret)
			}
		}
	*/
}

func TestGetDeep(t *testing.T) {
	/*
	   {
	       if a := getDeep(""); a != 0 {
	           t.Errorf("expect deep to be 0, but %d\n", a)
	       }
	       if a := getDeep("/"); a != 0 {
	           t.Errorf("expect deep to be 0, but %d\n", a)
	       }
	       if a := getDeep("/a"); a != 1 {
	           t.Errorf("expect deep to be 1, but %d\n", a)
	       }
	       if a := getDeep("/a/"); a != 1 {
	           t.Errorf("expect deep to be 1, but %d\n", a)
	       }
	       if a := getDeep("/a/b"); a != 2 {
	           t.Errorf("expect deep to be 2, but %d\n", a)
	       }
	       if a := getDeep("/a/b/"); a != 2 {
	           t.Errorf("expect deep to be 2, but %d\n", a)
	       }
	       if a := getDeep("/a/b/c"); a != 3 {
	           t.Errorf("expect deep to be 3, but %d\n", a)
	       }
	       if a := getDeep("/a/b/c/"); a != 3 {
	           t.Errorf("expect deep to be 3, but %d\n", a)
	       }
	       if a := getDeep("/a/b/c/d"); a != 4 {
	           t.Errorf("expect deep to be 4, but %d\n", a)
	       }
	       if a := getDeep("/a/b/c/d/"); a != 4 {
	           t.Errorf("expect deep to be 4, but %d\n", a)
	       }
	       if a := getDeep("/a/b/c/d/e"); a != 5 {
	           t.Errorf("expect deep to be 5, but %d\n", a)
	       }
	       if a := getDeep("/a/b/c/d/e/"); a != 5 {
	           t.Errorf("expect deep to be 5, but %d\n", a)
	       }
	   }
	*/
}

func TestGetDirectory(t *testing.T) {
	{
		t.Error(getDirectory("http://www.baidu.com", 0))
		t.Error(getDirectory("baidu.com", 1))
		t.Error(getDirectory("baidu.com/", 0))
		t.Error(getDirectory("baidu.com/", 1))
		t.Error(getDirectory("baidu.com/", 2))
		t.Error(getDirectory("baidu.com/a", 0))
		t.Error(getDirectory("baidu.com/a", 1))
		t.Error(getDirectory("baidu.com/a", 2))
		t.Error(getDirectory("baidu.com/a", 3))
		t.Error(getDirectory("baidu.com/a/", 0))
		t.Error(getDirectory("baidu.com/a/b", 1))
		t.Error(getDirectory("baidu.com/a/b", 2))
		t.Error(getDirectory("baidu.com/a/b/c", 3))

	}
}
