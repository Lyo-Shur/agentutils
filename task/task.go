package task

import "net/http"

// 前置处理
type Task interface {
	// 返回值 是否继续处理
	Do(w http.ResponseWriter, r *http.Request) bool
}
