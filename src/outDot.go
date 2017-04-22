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
	if err := g.SetDir(true); err != nil {
		panic(err)
	}
	if err := g.AddNode("G", "Hello", nil); err != nil {
		panic(err)
	}
	if err := g.AddNode("G", "World", nil); err != nil {
		panic(err)
	}
	if err := g.AddEdge("Hello", "World", true, nil); err != nil {
		panic(err)
	}
	s := g.String()
	// fmt.Println(s)
	
	file, err := os.Create(`./output.dot`)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(s))

}
