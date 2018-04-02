package hpb

import (
	//"github.com/hpb-project/ghpb/consensus/prometheus"
	"net/url"
	"time"
	"github.com/hpb-project/ghpb/common/log"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/hpb-project/ghpb/network/rpc"
	"encoding/json"
	"github.com/hpb-project/ghpb/core"
	"github.com/hpb-project/ghpb/consensus/prometheus"
)

var RootPath = "/miners"

type ZkMiner struct {
	blockChain *core.BlockChain
	hpb        *Hpb
	prometheus *prometheus.Prometheus
	conn       *zk.Conn
}

func NewZkMineHelper(chain *core.BlockChain, hpb *Hpb, prometheus *prometheus.Prometheus, zkAddr string) *ZkMiner {
	m := &ZkMiner{
		blockChain: chain,
		hpb:        hpb,
		prometheus: prometheus,
	}
	log.Info(" ZkMiner ########################################", "zk", zkAddr)
	conn, connChan, err := zk.Connect([]string{zkAddr}, time.Second*5)
	m.conn = conn
	if err != nil {
		log.Error("ZkMiner Connect to zookeeper server error!", "zk", err)
	}
	// 等待连接成功
	for {
		isConnected := false
		select {
		case connEvent := <-connChan:
			if connEvent.State == zk.StateConnected {
				isConnected = true
				log.Info("ZkMiner Connect to zookeeper server success!")
			}
		case _ = <-time.After(time.Second * 3): // 3秒仍未连接成功则返回连接超时
			log.Error("ZkMiner Connect to zookeeper server timeout!")
		}
		if isConnected {
			break
		}
	}
	return m
}

func (z *ZkMiner) MineZkStart() {
	// 拿到 /miners 下的所有节点
	// 调用RPC 请求proposal
	exists, _, err := z.conn.Exists(RootPath)
	if err != nil {
		log.Error("ZkMiner Zk exists error!", "zk", err)
	}
	if !exists {
		_, err := z.conn.Create(RootPath, []byte{0}, 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			log.Error("ZkMiner Zk Create error!", "zk", err)
		}
		log.Info("ZkMiner Create RootPath success!", "Path:", RootPath)
		return
	}
	miners, _, _ := z.conn.Children(RootPath)

	for _, miner := range miners {
		var m zkminer
		mStr, _ := url.QueryUnescape(miner)
		json.Unmarshal([]byte(mStr), &m)
		client, _ := rpc.Dial(m.RPC)
		var rest interface{}
		base, _ := z.hpb.Hpberbase()
		client.Call(&rest, "prometheus_propose", base, prometheus.RND(z.blockChain), true)
		log.Info("ZkMiner Call remote prometheus_propose", "response", rest)
	}

}

type zkminer struct {
	RPC  string `json:"rpc,omitempty"`
	ADDR string `json:"addr,omitempty"`
	RND  string `json:"rnd,omitempty"`
}

func (msg *zkminer) String() string {
	b, _ := json.Marshal(msg)
	return string(b)
}

func (z *ZkMiner) MinerRegister(host string) {
	//将本节点的address和random注册到 /miners
	addr, _ := z.hpb.Hpberbase()
	miner := &zkminer{
		RPC:  host,
		ADDR: addr.Hex(),
		RND:  prometheus.RND(z.blockChain),
	}

	data, _ := json.Marshal(&miner)

	currNode := RootPath + "/" + url.QueryEscape(string(data[:]))
	path, err := z.conn.Create(currNode, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))

	if err == nil { // 创建成功
		if path == currNode { // 返回的path表示在zookeeper上创建的znode路径
			log.Info("ZkMiner Add local node success!")
		} else {
			log.Error("ZkMiner Add local node returned different path")
		}
	} else { // 创建失败
		log.Error("ZkMiner Add local node error", "Path : ", path, "Error :", err)
	}
}
