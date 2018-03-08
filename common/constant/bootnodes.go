// Copyright 2018 The go-hpb Authors
// This file is part of the go-hpb.
//
// The go-hpb is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-hpb is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-hpb. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Hpb network.
var MainnetBootnodes = []string{
	// Hpbereum Foundation Go Bootnodes
	"hnode://91f7e7efce9fe8c8f48e07c87b8d64729fa4de6da16bcbc9ed5e3967b575c6faa12b88f039fe4c1d6c5d02867cba4d854881781110b346ca8eaca1cde6ba576b&1@127.0.0.1:30301",
	//"hnode://5c643c86d7cccae0f25916aae34a2ee076634a8123088b34b01390f68e3d71126a5fc16161c24b00d7ca05e38f6e85c556325b6dbe399564556d87aa6f08643e&1@47.75.51.144:30301",
	//"hnode://63c43fc19bca6770a602eaca3669e131ac38d07b3e40f119cc094af24eca9a98d3d9ad9ef0981f3878b0334a51b59f36bb1ecf3671a483dd5f6261df549c8aa6&1@47.91.47.79:30301",
	//"hnode://af6568c2913a99401fa567182a39f89bad7a0a273d2d7ba5a4ec1d02ad9c790c3be3f17ac92da84c5a9ed604cb7d44482783c85792d587f2bfc42b1dccd3d7e5&1@47.92.26.84:30301",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
}


// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}
