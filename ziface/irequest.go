package ziface

type IRequest interface {
	// 当前链接
	GetConnection() IConnection
	// 请求的消息数据
	GetData() []byte
}
