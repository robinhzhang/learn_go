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

//另外一种 内嵌的方式
type myTreeNode struct {
	*tree.Node //Embedding 相当于把Node这些变量平铺到mytreeNode中
}

//Traverse 重载
func (myNode myTreeNode) Traverse() {
	fmt.Println("This is a shadow function")
}

//扩展的结构定义方法
//t通过内嵌可以看到 代码简洁了很多
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()

}

func main() {

	//直接把tree.node 换成 myTreeNode ，
	root := myTreeNode{&tree.Node{Value: 3}}
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
	root.Node.Traverse()
	//shadow后的
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
	myNode := root
	myNode.postOrder()

	//与继承的区别
	//如果可以继承，var baseRoot *tree.Node
	//则 baseRoot = root 是可以把子类赋值给父类而获得额外的方法的，但是go不行 这不是继承哦
}
