package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/antchfx/xpath"
)

func main() {

	s := `<?xml version="1.0" encoding="UTF-8"?>
	
	<bookstore>
	
	<book category="COOKING">
	  <title lang="en">Everyday Italian</title>
	  <author>Giada De Laurentiis</author>
	  <year>2005</year>
	  <price>30.00</price>
	</book>
	
	<book category="CHILDREN">
	  <title lang="en">Harry Potter</title>
	  <author>J K. Rowling</author>
	  <year>2005</year>
	  <price>29.99</price>
	</book>
	
	<book category="WEB">
	  <title lang="en">XQuery Kick Start</title>
	  <author>James McGovern</author>
	  <author>Per Bothner</author>
	  <author>Kurt Cagle</author>
	  <author>James Linn</author>
	  <author>Vaidyanathan Nagarajan</author>
	  <year>2003</year>
	  <price>49.99</price>
	</book>
	
	<book category="WEB">
	  <title lang="en">Learning XML</title>
	  <author>Erik T. Ray</author>
	  <year>2003</year>
	  <price>39.95</price>
	</book>
	
	</bookstore> `

	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}

	expr, err := xpath.Compile("count(//book)")
	if err != nil {
		panic(err)
	}

	var root xpath.NodeNavigator
	root = xmlquery.CreateXPathNavigator(doc)
	val := expr.Evaluate(root)
	fmt.Println(val.(float64))

	expr = xpath.MustCompile("//book")
	val = expr.Evaluate(root)
	iter := val.(*xpath.NodeIterator)
	for iter.MoveNext() {
		fmt.Println(iter.Current().Value())
	}

	iter = expr.Select(root)
	for iter.MoveNext() {
		fmt.Println(iter.Current().Value())
	}
}
