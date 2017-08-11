package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var globalRoom *Room = NewRoom()

type Room struct {
	users map[string]net.Conn
}

func NewRoom() *Room {
	return &Room{
		users: make(map[string]net.Conn),
	}
}

func (r *Room) Join(user string, conn net.Conn) {
	//先判断用户是否在线，在线则踢出聊天室后再重新加入新的用户
	_, ok := r.users[user]
	if ok {
		r.Leave(user)
	}
	r.users[user] = conn
	// fmt.Printf("system: %s Join room\n", user)
	// conn.Write([]byte("System: " + user + " Join room\n"))
}

func (r *Room) Leave(user string) {
	//关闭连接
	//将用户踢出聊天室
	conn, ok := r.users[user]
	if !ok {
		return
	}
	conn.Close()
	delete(r.users, user)
	// fmt.Printf("system: %s Leave room\n", user)
}

func (r *Room) Broadcast(who string, msg string) {
	//遍历所有用户，发送消息
	// fmt.Println(r.users)
	tosend := fmt.Sprintf("%s:%s\n", who, msg)
	for user, conn := range r.users {
		if user == who {
			continue
		}
		conn.Write([]byte(tosend))
	}
}

//输入帐号和密码
//发送数据
//所有人可以收到
//关闭连接

//接收新连接
//验证用户名和密码
//等待用户输入
//向所有在线用户广播用户的输入
func handConn(conn net.Conn) { //conn类型为net.conn
	defer conn.Close() //函数退出，关闭连接
	r := bufio.NewReader(conn)
	content, err := r.ReadString('\n')
	if err != nil {
		log.Print(err)
		return
	}
	content = strings.TrimSpace(content)
	cmdfn := strings.Fields(content)
	if len(cmdfn) != 2 {
		conn.Write([]byte("bad input"))
		return
	}
	user := cmdfn[0]
	passwd := cmdfn[1]
	if passwd != "123456" {
		return
	}
	conn.Write([]byte("system: " + user + "登陆成功" + "\n"))
	// fmt.Printf("system: %s 登陆成功\n", user)
	globalRoom.Join(user, conn)
	// fmt.Println("---", globalRoom.users)
	// globalRoom.Join(user, conn)
	globalRoom.Broadcast("System: ", fmt.Sprintf("%s Join Room\n", user))
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		// conn.Write([]byte(line))
		line = strings.TrimSpace(line)
		// fmt.Println(line)
		globalRoom.Broadcast(user, line)
	}
	globalRoom.Broadcast("System: ", fmt.Sprintf("%s Leave Room\n", user))
	globalRoom.Leave(user)

}

func main() {
	addr := ":7777"                          //addr为IP:端口，一般忽略IP
	listener, err := net.Listen("tcp", addr) //监听端口
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close() //程序退出，关闭监听
	for {                  //循环接收请求
		conn, err := listener.Accept() //接收新的连接，阻塞连接
		if err != nil {
			log.Fatal(err)
		}
		go handConn(conn) //启动处理连接协程，哪里阻塞go哪里

	}
}
