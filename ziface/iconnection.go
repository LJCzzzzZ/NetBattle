package ziface

import "net"

type IConnection interface {
	// 启动链接
	Start()
	// 停止链接
	Stop()
	// 获取当前链接绑定的socket conn
	GetTCPConnection() *net.TCPConn
	// 获取当前链接模块的链接ID
	GetConnID() uint32
	// 获取远程客户端的 TCP状态 IP Port
	RemoteAddr() net.Addr
	// 发送数据
	Send(data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
