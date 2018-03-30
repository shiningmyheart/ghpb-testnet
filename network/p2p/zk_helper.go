package p2p

import (
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/pkg/errors"
	"net/url"
	"github.com/hpb-project/ghpb/network/p2p/discover"
	"github.com/hpb-project/ghpb/common/log"
)

var RootPath = "/nodes"

func StartZk(server *Server, readOnly bool) error {
	log.Info("########################################", "zk", server.ZkAddress)
	conn, connChan, err := zk.Connect([]string{server.ZkAddress}, time.Second*5)
	if err != nil {
		log.Error("Connect to zookeeper server error!", "zk", err)
		return err
	}
	// 等待连接成功
	for {
		isConnected := false
		select {
		case connEvent := <-connChan:
			if connEvent.State == zk.StateConnected {
				isConnected = true
				log.Info("Connect to zookeeper server success!")
			}
		case _ = <-time.After(time.Second * 3): // 3秒仍未连接成功则返回连接超时
			log.Error("Connect to zookeeper server timeout!")
			return errors.New("Connect to zookeeper server timeout!")
		}
		if isConnected {
			break
		}
	}

	if !readOnly {
		currNode := RootPath + "/" + url.QueryEscape(server.Self().String())
		path, err := conn.Create(currNode, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))

		if err == nil { // 创建成功
			if path == currNode { // 返回的path表示在zookeeper上创建的znode路径
				log.Info("Add local node success!")
			} else {
				log.Error("Add local node returned different path")
				return errors.New("Add local node returned different path " + currNode + " != " + path)
			}
		} else { // 创建失败
			log.Error("Add local node error", "Path : ", path, "Error :", err)
			return err
		}
	}

	go loop(server.quitZk, conn, server)
	//nodes,_ ,_:=conn.Children(RootPath)

	return nil
}

func loop(quit <-chan struct{}, conn *zk.Conn, server *Server) {
	// watch zk
	changeChan := make(chan struct{})
	go Watch(conn, changeChan, server)
	for {
		select {
		case <-quit:
			break
		case <-changeChan:
			// re-watch zk
			go Watch(conn, changeChan, server)
		}
	}
}

func Watch(conn *zk.Conn, changeChan chan struct{}, server *Server) {
	log.Info("watching")
	children, _, childCh, err := conn.ChildrenW(RootPath)
	if err != nil {
		log.Error("watch children error, ", err)
	}
	log.Info("get children nodes ", "Nodes : ", children)

	process(children, server)

	select {
	case childEvent := <-childCh:
		if childEvent.Type == zk.EventNodeChildrenChanged {
			log.Info("Event coming : ", "Event Type : ", childEvent.Type, "Event Path : ", childEvent.Path)
			changeChan <- struct{}{}
		}
	}
}

func process(nodes []string, server *Server) {
	for _, nodeStr := range nodes {
		//server.StaticNodes = append(server.StaticNodes,node)
		nodeStr, err := url.QueryUnescape(nodeStr)
		if err != nil {
			log.Error(fmt.Sprintf("Node URL Unescape Error %s: %v\n", nodeStr, err))
		}
		node, err := discover.ParseNode(nodeStr)
		if err != nil {
			log.Error(fmt.Sprintf("Node URL %s: %v\n", nodeStr, err))
			continue
		}
		server.StaticNodes = append(server.StaticNodes,node)
		server.addstatic <- node
	}
}
