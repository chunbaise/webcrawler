package downloader

import "base"

// 网络下载器的接口类型
type PageDownloader interface {
	Id() uint32                                        // 获得ID
	Download(req base.Request) (*base.Response, error) // 根据请求下载网页并返回响应
}

// ID生成器
type IdGenertor interface {
	GetUint32() uint32
}
