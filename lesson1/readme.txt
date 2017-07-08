go run hello.go  //表示运行hello.go文件
gofmt -w hello.go  //格式化hello.go文件
goimports  //等同于gofmt，可以引用使用了而未引用的包，或删除了引用了而未使用的包

go build hello.go  //生成一个二进制文件hello

GOOS=linux go build hello.go  //生成一个LINUX系统的可执行二进制文件
GOOS=windows go build -o hello.exe hello.go  //生成一个windows系统的可执行文件hello.exe

#指针：操作变量的变量
var x int   //声明变量
x = 1   //变量赋值
var p *int  //声明一个指针
p = &x
fmt.Println(p)   //打印指针地址
fmt.Println(*p)   //打印指针的值
*p = 2     //指针重新赋值
fmt.Println(x)  //打印X值，此时X值变为2

go get -d cd github.com/51reboot/golang-01-homework
go get github.com/icexin/golib   //download代码

cd $GOPATH/src/github.com/51reboot/golang-01-homework
git pull
cd lesson1
mkdir zhaoyuanhai

git add .
git commit -m 'first test push'
git push origin master
