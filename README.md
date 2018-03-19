## Go HPB

Official golang implementation of the HPB protocol.

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](#)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](#)



## Building the source (requires a Go version 1.7 or later)

###Clone source
```
    git clone git@github.com:hpb-project/ghpb-testnet.git $GOPATH/src/github.com/hpb-project/ghpb
```

###Building ghpb
```
    go install -a -v ./command/ghpb
```
or
```
    cd $GOPATH/src/github.com/hpb-project/ghpb
    make ghpb
```

###Building promfile

## Init with genesis file
```
$ geth --datadir <some/location/where/to/create/chain> init genesis.json
```

## Running ghpb
```
$ ghpb --identity "private hpb"  --rpcaddr 127.0.0.1  --rpc   --rpcport 8545  --maxpeers 2  --networkid 100  --datadir "./chain"  --nodiscover
```

## Attach to the node
```
ghpb attach ipc://path-to-chain-directory/ghpb.ipc
```