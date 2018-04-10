package main

import (
	"fmt"
	"bytes"
	"strconv"
	"encoding/json"
	"time"
	"runtime"
	"net/http"
	"strings"
)

func ByteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {
		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}

	return buffer.String()
}

type Tx struct {
	From string `json:"from"`

	To string `json:"to"`

	Value string `json:"value"`

	Gas string `json:"gas"`
}
type Data struct {
	Txs []Tx `json:"params"`

	Id int `json:"id"`

	Jsonrpc string `json:"jsonrpc"`

	Method string `json:"method"`
}

func send(url string, data string, c http.Client) {
	//fmt.Println(data)

	req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	//
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	//defer wg.Add(-1)

}

func main() {
	url := "http://127.0.0.1:8545"
	maxProcs := runtime.NumCPU()
	fmt.Println("cpu_count :", maxProcs)

	c := &http.Client{}
	datas := make([]interface{}, 0)
	num := float64(700)
	for i := 1; i < int(num); i++ {
		ss := Data{}
		ss.Method = "hpb_sendTransaction"
		ss.Jsonrpc = "2.0"
		ss.Id = 67
		tx_ := Tx{}
		tx_.From = "0x654b8d29253cb213c65a7d9f9aae3343809f54a6"
		tx_.To = "0x428a9787f2066e3887277a9a38d49f86cfbc7e91"

		//wg.Add(1)
		bb := byte(i * 100)
		vv := ByteToHex([]byte{bb})
		//fmt.Println(vv)
		tx_.Value = "0x" + string(vv)

		tx_.Gas = "0x" + strconv.FormatInt(int64(100), 16)
		ss.Txs = append(ss.Txs, tx_)
		datas = append(datas, ss)
		//send(url,string(b))
	}

	b, _ := json.Marshal(datas)
	send(url, string(b), *c)
	t1 := time.Now()

	t2 := time.Now()
	dd := t2.Sub(t1).Seconds()
	dd1 := t2.Sub(t1).Nanoseconds()
	fmt.Println(dd, "/s")
	fmt.Println(dd1, "/Nanoseconds")
	fmt.Println(num/dd, "counts/s")
}
