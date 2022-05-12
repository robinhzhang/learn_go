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

//关于结构体的扩展
//扩充系统类型或者使用别人的类型
//定义别名
//使用组合

//比如我们想扩展tree.Node这个结构，组合
type myTreeNode struct {
	node *tree.Node
}

//扩展的结构定义方法
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()

}

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

	//组合 后续遍历
	myNode := myTreeNode{&root}
	myNode.postOrder()
}
