package analyzer

import (
	"base"
	"net/http"
)

// 分析器的接口类型
type Analyzer interface {
	Id() uint32 // 获得ID
	Analyze(
		respParsers []ParseResponse,
		resp base.Response) ([]base.Data, []error)
}

// 被用于解析HTTP响应的函数类型
type ParseResponse func(httpResp *http.Response, respDepth uint32) ([]base.Data, []error)

// 分析池接口
type AnalyzerPool interface {
	Take() (Analyzer, error)
	Return(analyzer Analyzer) error
	Total() uint32
	Used() uint32
}
