package znet

import "ljc/NetBattle/ziface"

type BaseRouter struct{}

// 在处理conn业务之前的钩子方法Hook
func (br *BaseRouter) preHandle(request ziface.IRequest) {

}

// 在处理conn业务的主方法
func (br *BaseRouter) Handle(request ziface.IRequest) {

}

// 在处理conn业务之后的钩子方法Hook
func (br *BaseRouter) PostHandle(request ziface.IRequest) {

}
