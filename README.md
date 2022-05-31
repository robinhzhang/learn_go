# learn_go
learn go
go go go
merge request

小记 11

go mod tidy
清理多余mod的包

旧项目迁移到mod
添加go.mod

go mod init [名字]

go build ./...

小记12

按照官方包的方式整理了代码

go build ./...

2.1.2 duck typing

像鸭子一样就是鸭子

描述外部行为而非内部结构

python 支持 但是需要注释说明 C++通过模版类型实现，也需要说明
java则不支持，必须实现(extends)

2.1.3

go 是使用者来定义接口

隐式实现接口

/213interface_duck
2.1.4
/213interface_duck/real/Retriever.go
/213interface_duck/main.go