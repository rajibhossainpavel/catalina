package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/fetchbot"
)

func main() {
	f := fetchbot.New(fetchbot.HandlerFunc(handler))
	f.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.88 Safari/537.36"
	queue := f.Start()
	queue.SendStringGet("https://dsebd.org/latest_share_price_scroll_by_ltp.php")
	queue.Close()
}

func handler(ctx *fetchbot.Context, res *http.Response, err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Printf("[%d] %s %s\n", res.StatusCode, ctx.Cmd.Method(), ctx.Cmd.URL())
}
