package zk

import (
	"testing"
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"
	//"net/url"
	"sync"
)

var hosts = []string{"39.107.116.40:2181"}
var flags int32 = zk.FlagEphemeral
var acls = zk.WorldACL(zk.PermAll)

var RootPath = "/nodes"
var all = make(map[string]struct{})

func process(nodes []string) {
	for _, node := range nodes {
		if _, has := all[node]; has {
			continue
		}
		all[node] = struct{}{}
		fmt.Println("process done current nodes : ", nodes)
	}
}

var wg = sync.WaitGroup{}

func TestZk(t *testing.T) {
	wg.Add(1)
	quit := make(chan struct{})

	conn, connChan, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		t.Error("error", err)
	}
	// 等待连接成功
	for {
		isConnected := false
		select {
		case connEvent := <-connChan:
			if connEvent.State == zk.StateConnected {
				isConnected = true
				fmt.Println("Connect to zookeeper server success!")
			}
		case _ = <-time.After(time.Second * 3): // 3秒仍未连接成功则返回连接超时
			t.Error("Connect to zookeeper server timeout!")
		}
		if isConnected {
			break
		}
	}
	//req := "hnode://d7c14e72eb3ce550c2fe95664f16d7b8843296f2d775c679b7523728f852e7b30fefd9c2534e1b236ae06619a36b564805ed8a11dc17147b1f5f8623b62b9be8&4@116.62.175.114:3001"
	//encodeurl := url.QueryEscape(req)
	currNode := RootPath + "/" + "1"
	path, err := conn.Create(currNode, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))

	if err == nil { // 创建成功
		if path == currNode { // 返回的path表示在zookeeper上创建的znode路径
			fmt.Println("Add local node success!")
		} else {
			t.Error("Add local node returned different path " + currNode + " != " + path)
		}
	} else { // 创建失败
		t.Error("Add local node faild", err)
	}

	go loop(quit, conn)
	wg.Wait()
}

func loop(quit chan struct{}, conn *zk.Conn) {
	// watch zk
	changeChan := make(chan struct{})
	go watch(conn, changeChan)
	for {
		select {
		case <-quit:
			break
		case <-changeChan:
			// re-watch zk
			go watch(conn, changeChan)
		}
	}
}

func watch(conn *zk.Conn, changeChan chan struct{}) {
	fmt.Println("watching")
	children, _, childCh, err := conn.ChildrenW(RootPath)
	if err != nil {
		fmt.Println("watch children error, ", err)
	}
	fmt.Println("children",children)
	process(children)

	select {
	case childEvent := <-childCh:
		if childEvent.Type == zk.EventNodeChildrenChanged {
			fmt.Println("Event coming : ", childEvent.Type, childEvent.Path)
			changeChan <- struct{}{}
		}
	}
}

func TestAddNodes(t *testing.T) {
	//wg.Add(1)
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//req := "hnode://d7c14e72eb3ce550c2fe95664f16d7b8843296f2d775c679b7523728f852e7b30fefd9c2534e1b236ae06619a36b564805ed8a11dc17147b1f5f8623b62b9be8&4@116.62.175.114:3002"
	//encodeurl := url.QueryEscape(req)

	resp, err := conn.Create(RootPath , []byte{0}, 0, acls)
	fmt.Println(resp)
	fmt.Println(err)
	//wg.Wait()
}


func TestGetNodes(t *testing.T) {
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	children,_,_:=conn.Children(RootPath)
	fmt.Println(children)
}
