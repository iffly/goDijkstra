package main

import (
	"fmt"
)

type nodeInfo struct {
	name string
}

type connInfo struct {
	from nodeInfo
	to   nodeInfo
	min  []nodeInfo
	dis  int
	done bool
}

var node = make([]nodeInfo, 0)
var conn = make([]connInfo, 0)
var defaultDisMax = 10000

func main() {
	// test data
	node = append(node, nodeInfo{name: "北京"})
	node = append(node, nodeInfo{name: "西安"})
	node = append(node, nodeInfo{name: "苏州"})
	node = append(node, nodeInfo{name: "上海"})
	node = append(node, nodeInfo{name: "杭州"})
	node = append(node, nodeInfo{name: "厦门"})
	node = append(node, nodeInfo{name: "深圳"})
	conn = append(conn, connInfo{from: nodeInfo{name: "北京"}, to: nodeInfo{name: "西安"}, dis: 1144})
	conn = append(conn, connInfo{from: nodeInfo{name: "北京"}, to: nodeInfo{name: "苏州"}, dis: 1140})
	conn = append(conn, connInfo{from: nodeInfo{name: "北京"}, to: nodeInfo{name: "上海"}, dis: 1212})
	conn = append(conn, connInfo{from: nodeInfo{name: "北京"}, to: nodeInfo{name: "杭州"}, dis: 1271})
	conn = append(conn, connInfo{from: nodeInfo{name: "上海"}, to: nodeInfo{name: "苏州"}, dis: 100})
	conn = append(conn, connInfo{from: nodeInfo{name: "上海"}, to: nodeInfo{name: "杭州"}, dis: 175})
	conn = append(conn, connInfo{from: nodeInfo{name: "上海"}, to: nodeInfo{name: "厦门"}, dis: 1021})
	conn = append(conn, connInfo{from: nodeInfo{name: "上海"}, to: nodeInfo{name: "深圳"}, dis: 1427})
	conn = append(conn, connInfo{from: nodeInfo{name: "杭州"}, to: nodeInfo{name: "厦门"}, dis: 877})
	conn = append(conn, connInfo{from: nodeInfo{name: "杭州"}, to: nodeInfo{name: "深圳"}, dis: 1300})
	conn = append(conn, connInfo{from: nodeInfo{name: "厦门"}, to: nodeInfo{name: "深圳"}, dis: 578})

	// ctor map
	calcMap := make([][]connInfo, 0)
	for _, from := range node {
		calcMapOne := make([]connInfo, 0)
		for _, to := range node {
			connInfoOne := connInfo{
				from: from,
				to:   to,
				min:  make([]nodeInfo, 0),
			}

			disMin := getDis(from, to)
			disTmp := getDis(to, from)
			if disMin > disTmp {
				disMin = disTmp
			}
			if from == to {
				disMin = 0
			}
			connInfoOne.dis = disMin

			calcMapOne = append(calcMapOne, connInfoOne)
		}
		calcMap = append(calcMap, calcMapOne)
	}

	vs := calcMap[0]
	// Dijkstra
	for {
		// find min dis, mark done
		disMin := defaultDisMax
		k := 0
		for i, info := range vs {
			if true == info.done {
				continue
			}

			if info.dis > disMin {
				continue
			}

			disMin = info.dis
			k = i
		}
		vs[k].done = true

		// when to finish the Dijkstra
		if defaultDisMax == disMin {
			break
		}

		// use min dis to relay
		for i, info := range calcMap[k] {
			if disMin+info.dis >= calcMap[0][i].dis {
				continue
			}

			vs[i].dis = disMin + info.dis
			vs[i].min = append(vs[k].min, node[k])
		}
	}

	// print
	for _, info := range vs {
		fmt.Println(info.from.name, "->", info.to.name, "路过：", info.min, "距离：", info.dis)
	}
}

func getDis(from, to nodeInfo) (dis int) {
	for _, info := range conn {
		if info.from == from && info.to == to {
			return info.dis
		}
	}

	return defaultDisMax
}
