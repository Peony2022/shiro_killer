package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/panjf2000/ants"
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
	flag.IntVar(&Ant, "t", 50, "Number of goroutines")
	flag.StringVar(&CheckContent, "chk", "rO0ABXNyADJvcmcuYXBhY2hlLnNoaXJvLnN1YmplY3QuU2ltcGxlUHJpbmNpcGFsQ29sbGVjdGlvbqh/WCXGowhKAwABTAAPcmVhbG1QcmluY2lwYWxzdAAPTGphdmEvdXRpbC9NYXA7eHBwdwEAeA==", "Check Content")
	flag.StringVar(&NRemeberMe, "rm", "rememberMe", "Name of rememberMe")
	flag.Parse()
}
func StartTask(TargetUrl string) {
	if !ShiroCheck(TargetUrl) {
		_, result := KeyCheck(TargetUrl)
		fmt.Println(TargetUrl, ": \n", result)
	} else {
		fmt.Println(TargetUrl, ": ", "Shiro not exist!")
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
		defer ants.Release()
		pool, _ := ants.NewPool(Ant)

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
				ShiroKeys = append(ShiroKeys, string(UnFormatted))
			}
		}

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
			pool.Submit(func() {
				StartTask(string(TargetUrl))
				wg.Done()
			})
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
