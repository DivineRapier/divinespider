package core

import (
	"flag"
	"fmt"
)

var (
	crawlURL string
	path     string
	script   string
	option   string
	header   string
	cookie   string
	mode     int
)

var executableFilePath = map[string]string{
	"darwin":  "/bin/darwin/phantomjs",
	"linux":   "/bin/linux/phantomjs",
	"windows": "/bin/windows/phantomjs.exe",
}

func usage() {
	fmt.Print(`
        Usage:

            ./divinespider option [arguments]

        The options are:

            -url               	url which you wanna crawl
            -path    			phantom javascript file
            -option    			arguments of phantomjs to execute
            -header            	custom header
            -cookie            	custom cookie
            -mode             	crawler mode
            -help
        `)
}

func init() {
	flag.StringVar(&crawlURL, "url", "", "url")
	flag.StringVar(&path, "path", "", "phantom path")
	flag.StringVar(&script, "script", "", "phantom javascript file")
	flag.StringVar(&option, "option", "", "phantom command option")
	flag.StringVar(&header, "header", "", "header")
	flag.StringVar(&cookie, "cookie", "", "cookie")
	flag.IntVar(&mode, "mode", 0, "crawler mode")
	flag.Usage = usage
	flag.Parse()
}
