package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	duration := 1 * time.Second
	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "hello")
	d := time.Now().Add(duration)
	ctx, cancel := context.WithDeadline(ctx, d)

	defer cancel()

	doRequest(ctx, "https://ya.ru")
}

func doRequest(ctx context.Context, requestStr string) {
	req, _ := http.NewRequest(http.MethodGet, requestStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("response completed, status code=%d\n", res.StatusCode)
		fmt.Println(ctx.Value("test"))
	case <-ctx.Done():
		fmt.Println("request too long")
	}
}
