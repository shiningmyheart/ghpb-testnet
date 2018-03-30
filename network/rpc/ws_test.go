package rpc

import (
	"testing"
	"context"
	"fmt"
)

func TestWS(t *testing.T) {
	client,_ := DialWebsocket(context.Background(),"ws://127.0.0.1:8546","")
	defer client.Close()
	var resp interface{}
	if err := client.Call(&resp, "hpb_mining", nil); err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}