package znet

import (
	"fmt"
	"ljc/NetBattle/ziface"
	"net"
)

type Connection struct {
	// socket TCP套接字
	Conn *net.TCPConn
	// 链接的ID
	ConnID uint32
	// 当前链接的状态
	isClosed bool

	// 退出告知channel
	ExitChan chan bool

	//该链接处理的方法Router
	Router ziface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// 读业务
func (c *Connection) StartReader() {
	fmt.Println("Reader Goruntine is running...")
	defer fmt.Println("connID = ", c.ConnID, " Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("revc buf err", err)
			continue
		}
		//得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			data: buf,
		}
		// 从路由中找到注册绑定的Conn对应的Router调用
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start() ConnID = ", c.ConnID)
	// 启动读数据业务
	go c.StartReader()
	// TODO 写数据业务

}

// 停止链接
func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID = ", c.ConnID)

	if c.isClosed == true {
		return
	}

	c.isClosed = true

	c.Conn.Close()

	close(c.ExitChan)
}

// 获取当前链接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的 TCP状态 IP Port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send(data []byte) error {
	return nil
}
