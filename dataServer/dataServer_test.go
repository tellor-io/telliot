package dataServer

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestDataServer(t *testing.T) {
	exitCh := make(chan int)
	ds, err := CreateServer()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	ds.Start(ctx, exitCh)

	time.Sleep(5000 * time.Millisecond)

	resp, err := http.Get("http://localhost:5000/balance")
	defer resp.Body.Close()
	fmt.Printf("Finished: %+v", resp)

}
