package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d  %s", secs, nbytes, url)
}

// https://www.google.com/ https://www.facebook.com/ https://www.reddit.com/ https://www.odoo.com/ https://www.youtube.com/ https://maktoob.yahoo.com/ https://slashdot.org/ https://www.linkedin.com/ https://github.com/ https://soundcloud.com/ https://twitter.com/ https://www.instagram.com https://www.bet365.com/ https://www.youm7.com/ https://www.file-upload.com/ https://www.masrawy.com/ https://www.wikipedia.org/ https://www.whatsapp.com/ https://www.olx.com.eg/ https://egypt.souq.com/eg-en/ https://www.elbalad.news/ https://www.almasryalyoum.com/ https://www.yallakora.com/ https://mawdoo3.com/ https://360vuz.com/ https://www.sm3ha.com/ https://www.premierleague.com/ https://www.tiktok.com/ https://www.rt.com/ http://ghenwety.com/ https://www.filgoal.com/ https://www.pinterest.com/ https://outlook.live.com/ https://www.mediafire.com/ https://www.jumia.com.eg/ https://www.elwatannews.com/ https://muhtwaplus.com/
