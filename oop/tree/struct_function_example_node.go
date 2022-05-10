package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//go仅支持封装，不支持集成和多肽
// struct
//go没有构造函数这种概念，如果我们需要自己控制来进行struct生成，可以用工厂函数的方法
//返回的局部变量的地址
//c++局部变量是分配到栈上，函数退出就销毁了；java这种具有垃圾回收机制有些是分配到堆上
//go 是编译器根据情况是分配到堆上还是栈上（！）
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

//对象函数 显式接受者
func (node treeNode) print() {
	fmt.Println(node.value)
}

//穿值的 不加星号穿值
func (node treeNode) setValue(value int) {
	node.value = value
}

//穿值的 不加星号穿值
func (node *treeNode) setValueRef(value int) {
	//nil指针也可以调用方法,但是不能够调用
	if node == nil {
		fmt.Println("nil value ignore")
		return
	}
	node.value = value
}

//遍历树
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

//关于设计
//改变内容使用指针接受者
//结构太大考虑使用指针接受者
//一致性：如果有指针接受者，最好都是指针接受者
//值接收者 是go特有
//指针 值都可以通过指针接受者调用
func main() {
	var root treeNode

	//.访问函数
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.left.right = new(treeNode)
	root.right.left = createNode(2)
	fmt.Println(root)
	//setvalue穿值
	root.right.left.print()
	root.right.left.setValue(4)
	root.right.left.print()
	root.right.left.setValueRef(4)
	root.right.left.print()
	//如果对象的地址能调用函数吗
	pRoot := &root
	pRoot.print()
	pRoot.setValueRef(100)
	pRoot.print()

	//遍历
	root.traverse()
	var nilroot *treeNode
	nilroot.setValueRef(1)
	//slice
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, nil},
	}
	fmt.Println(nodes)
}
