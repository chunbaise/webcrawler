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

// 网页下载器池的接口类型
type PageDownloaderPool interface {
	Take() (PageDownloader, error)  // 从池中取出一个页面下载器
	Return(dl PageDownloader) error // 把一个网页下载器归还给池
	Total() uint32                  // 获得池的总容量
	Used() uint32                   // 获得正在被使用的下载器数量
}
