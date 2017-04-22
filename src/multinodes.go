package main

import (
	"os"
	"github.com/awalterschulze/gographviz"
)

func main() {
	g := gographviz.NewGraph()
	if err := g.SetName("G"); err != nil {
		panic(err)
	}
	if err := g.SetDir(true); err != nil { // 有向グラフか
		panic(err)
	}
	
	if err := g.AddNode("G", "Hello", nil); err != nil {
		panic(err)
	}
	if err := g.AddNode("G", "Beautiful", nil); err != nil {
		panic(err)
	}
	if err := g.AddNode("G", "World", nil); err != nil {
		panic(err)
	}

	// src(第1引数)からdst(第2引数)へ方向を付ける
	if err := g.AddEdge("Hello", "World", true, nil); err != nil {
		panic(err)
	}
	if err := g.AddEdge("Beautiful", "World", true, nil); err != nil {
		panic(err)
	}
	s := g.String()
	// fmt.Println(s)
	
	file, err := os.Create(`./multi.dot`)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(s))

}
