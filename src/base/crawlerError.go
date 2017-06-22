package base

import (
	"bytes"
	"fmt"
)

type ErrorType string

type CrawlerError interface {
	Type() ErrorType
	Error() string
}

// 爬虫错误的实现
type myCrawlerError struct {
	errType    ErrorType
	errMsg     string
	fullErrMsg string
}

// 错误类型常量
const (
	DOWNLOADER_ERROR     ErrorType = "Downloader Error"
	ANALYZER_ERROR       ErrorType = "Analyzer Error"
	ITEM_PROCESSOR_ERROR ErrorType = "Item Processor Error"
)

// 创建一个新的掉错误
func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawlerError{errType: errType, errMsg: errMsg}
}

// 获取错误类型
func (ce *myCrawlerError) Type() ErrorType {
	return ce.errType
}

// 获取错误信息
func (ce *myCrawlerError) Error() string {
	if ce.fullErrMsg == "" {
		ce.genFullErrMsg()
	}
	return ce.fullErrMsg
}

// 生成错误提示信息，并给相应的字段赋值
func (ce *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("Crawler Error: ")
	if ce.errType != "" {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf("%s\n", buffer.String())
	return
}
