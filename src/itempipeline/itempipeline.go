package itempipeline

import "base"

// 条目处理管道的接口类型
type ItemPipeLine interface {
	// 发送条目
	Send(item base.Item) []error
	// FailFast方法会返回一个布尔值。该值表示当前的条目处理
	// 管道是否是快速失败的。这里的快速失败指的是：只要对某个条目的处理的流程
	// 在某一个步骤上出错，那么条目处理管道就会忽略掉后续的所有处理步骤并报告错误。
	FailFast() bool
	// 设置是否快速失败
	SetFailFast(failFast bool)
	// 获得已发送、已接收和已处理的条目的计数值
	// 更确切地说，作为结果值的切片总会有3个元素值，这3个值会分别代表前述的3个计数
	Count() []uint64
	// 获取正在被处理的条目的数量
	ProcessingNumber() uint64
	// 获取摘要信息
	Summary() string
}

type ProcessItem func(item base.Item) (result base.Item, err error)
