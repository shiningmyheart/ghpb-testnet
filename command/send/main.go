package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"time"
	"runtime"
	"net/http"
	"strings"
	"os"
)

type Tx struct {
	From string `json:"from"`

	To string `json:"to"`

	Value string `json:"value"`
	Gas   string `json:"gas"`
}
type Data struct {
	Txs []Tx `json:"params"`

	Id int `json:"id"`

	Jsonrpc string `json:"jsonrpc"`

	Method string `json:"method"`
}
type Acc struct {
	Params []string `json:"params"`

	Id int `json:"id"`

	Jsonrpc string `json:"jsonrpc"`

	Method string `json:"method"`
}

func geta(url string, data string, c http.Client) []string {
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
	body, _ := ioutil.ReadAll(resp.Body)

	type re struct {
		Id int `json:"id"`

		Jsonrpc string `json:"jsonrpc"`

		Result []string `json:"result"`
	}
	var res re
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Println(err)
	}
	return res.Result
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
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", bytes.NewBuffer(body).String())
}

//var wg sync.WaitGroup
func main() {

	url := "http://127.0.0.1:8545"
	//
	var acc Acc
	acc.Method = "hpb_accounts"
	acc.Jsonrpc = "2.0"
	acc.Id = 68
	acc.Params = nil
	acc_d, _ := json.Marshal(acc)
	maxProcs := runtime.NumCPU()
	fmt.Println("cpu_count :", maxProcs)

	c := &http.Client{}
	datas := make([]interface{}, 0)
	ress := geta(url, string(acc_d), *c)

	num, _ := strconv.Atoi(os.Args[1])
	loop, _ := strconv.Atoi(os.Args[2])
	fmt.Printf("Prepare send transaction count : %d .\n", num*loop)
	for i := 1; i <= num; i++ {
		ss := Data{}
		ss.Method = "hpb_sendTransaction"
		ss.Jsonrpc = "2.0"
		ss.Id = 67
		tx_ := Tx{}
		tx_.From = ress[0]
		tx_.To = ress[1]
		//wg.Add(1)
		tx_.Value = "0x" + strconv.FormatInt(int64(i*100), 16)
		tx_.Gas = "0x" + strconv.FormatInt(int64(100), 16)

		ss.Txs = append(ss.Txs, tx_)
		datas = append(datas, ss)
		//send(url,string(b))
	}
	//wg.Wait()

	b, _ := json.Marshal(datas)
	//fmt.Println(string(b))
	//nnn :=300
	for i := 1; i <= loop; i++ {
		send(url, string(b), *c)
		wait, _ := strconv.ParseBool(os.Args[3])
		if wait {
			time.Sleep(2 * time.Second)
		}
	}
	fmt.Println("Done.")
}
