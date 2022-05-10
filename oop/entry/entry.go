package main

import (
	"fmt"
	"github.com/huahua/learngo/oop/tree"
)

//关于封装
// go 明明camelcase
//大小写表示public和private
//包
//每个目录一个包
//为结构定义的方法必须放在同一个包内
//可以是不同文件

func main() {
	var root tree.Node

	//.访问函数
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Left.Right = new(tree.Node)
	root.Right.Left = tree.CreateNode(2)
	fmt.Println(root)
	//setvalue穿值
	root.Right.Left.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	root.Right.Left.SetValueRef(4)
	root.Right.Left.Print()
	//如果对象的地址能调用函数吗
	pRoot := &root
	pRoot.Print()
	pRoot.SetValueRef(100)
	pRoot.Print()

	//遍历
	root.Traverse()
	var nilroot *tree.Node
	nilroot.SetValueRef(1)
	//slice
	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, nil},
	}
	fmt.Println(nodes)
}
