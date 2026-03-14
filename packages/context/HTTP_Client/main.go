package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getUrlWithTimeout(ctx context.Context, url string) error {
	// 创建一个带有超时的请求
	req, err := http.NewRequestWithContext(ctx, "", url, nil)
	if err != nil {
		return err
	}

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(os.Stdout, "staus: %v\n", resp.Status)
	return err
}

func main() {
	ctx, cancel := context.WithTimeoutCause(context.Background(), 100*time.Millisecond, nil)
	defer cancel()
	err := getUrlWithTimeout(ctx, "https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
}
