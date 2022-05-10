package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//go仅支持封装，不支持集成和多肽
// struct
//go没有构造 函数这种概念，如果我们需要自己控制来进行struct生成，可以用工厂函数的方法
//返回的局部变量的地址
//c++局部变量是分配到栈上，函数退出就销毁了；java这种具有垃圾回收机制有些是分配到堆上
//go 是编译器根据情况是分配到堆上还是栈上（！）
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//对象函数 显式接受者
func (node Node) Print() {
	fmt.Println(node.Value)
}

//穿值的 不加星号穿值
func (node Node) SetValue(value int) {
	node.Value = value
}

//穿值的 不加星号穿值
func (node *Node) SetValueRef(value int) {
	//nil指针也可以调用方法,但是不能够调用
	if node == nil {
		fmt.Println("nil value ignore")
		return
	}
	node.Value = value
}

//遍历树
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

//关于设计
//改变内容使用指针接受者
//结构太大考虑使用指针接受者
//一致性：如果有指针接受者，最好都是指针接受者
//值接收者 是go特有
//指针 值都可以通过指针接受者调用
