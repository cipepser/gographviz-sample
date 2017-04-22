package main

import (
	"os"
	"github.com/awalterschulze/gographviz"
)

func main() {
	// 新規グラフ作成
	g := gographviz.NewGraph()
	if err := g.SetName("G"); err != nil {
		panic(err)
	}
	// 有向グラフか
	if err := g.SetDir(true); err != nil { 
		panic(err)
	}
	
	// グラフ全体の設定
	if err := g.AddAttr("G", "bgcolor", "\"#343434\""); err != nil { 
		panic(err)
	}
	
	// Node設定
	nodeAttrs := make(map[string]string)
	// if _, err := gographviz.NewAttr("colorscheme"); err != nil {
	//  	panic(err)
	// }
	nodeAttrs["colorscheme"] = "rdylgn11"
	nodeAttrs["style"] = "\"solid,filled\""
	nodeAttrs["fontcolor"] = "6"
	nodeAttrs["fontname"] = "\"Migu 1M\""
	nodeAttrs["color"] = "7"
	nodeAttrs["fillcolor"] = "11"
	nodeAttrs["shape"] = "doublecircle"
	
	// Node追加
	if err := g.AddNode("G", "Hello", nodeAttrs); err != nil {
		panic(err)
	}
	if err := g.AddNode("G", "Beautiful", nodeAttrs); err != nil {
		panic(err)
	}
	if err := g.AddNode("G", "World", nodeAttrs); err != nil {
		panic(err)
	}
	
	// Edge設定
	edgeAttrs := make(map[string]string)
	edgeAttrs["color"] = "white"

	// Edge(Node関係の関係)追加
	// src(第1引数)からdst(第2引数)へ方向を付ける
	if err := g.AddEdge("Hello", "World", true, edgeAttrs); err != nil {
		panic(err)
	}
	if err := g.AddEdge("Beautiful", "World", true, edgeAttrs); err != nil {
		panic(err)
	}
	s := g.String()
	
	// dotファイル出力
	file, err := os.Create(`./dressup.dot`)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(s))

}
