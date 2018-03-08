## Go HPB

Official golang implementation of the HPB protocol.

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](#)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](#)



## Building the source

Building ghpb requires both a Go (version 1.7 or later)

    git clone git@github.com:hpb-project/ghpb.git $GOPATH/src/github.com/hpb-project

then

    go install -a -v ./command/ghpb
or
```
    cd ./github.com/hpb-project/ghpb
    make ghpb
```
when you do this,there is a shell cmd generated at './bulid/bin',called ghpb

note: if you want to make all,maybe have some problem

## Running ghpb
```
$ ghpb --identity "private hpb"  --rpcaddr 127.0.0.1  --rpc   --rpcport 8545  --maxpeers 2  --networkid 100  --datadir "./chain"  --nodiscover
```

## Attach to the node
```
ghpb attach ipc://path-to-chain-directory/ghpb.ipc
```