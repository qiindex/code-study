package demo_design

import (
	"fmt"
	"math/rand"

	"github.com/pkg/errors"
)

/*
设计一个物流状态更新模块，要求使用Go语言实现以下功能：
1. 并发处理：能够同时处理多个物流单号的状态更新请求
2. 顺序保证：相同物流单号的更新请求必须严格按提交顺序执行
3. 优雅关闭：收到关闭信号时，完成所有正在处理的请求后退出
4. 错误重试：（如有更好）当状态更新函数返回错误时，自动重试（最多3次）
*/
type UpdateRequest struct {
	TrackingID string // 物流单号
	NewStatus  string // 新状态
}

// 实际执行状态更新的函数（已存在，可能返回错误）
func RealUpdate(trackingID string, status string) error {
	// 模拟可能失败的业务逻辑
	if rand.Intn(10) < 3 { // 30%概率失败
		return errors.New("update failed")
	}
	fmt.Printf("Updated %s to %s\n", trackingID, status)
	return nil
}

type TypeName map[string]chan UpdateRequest
type Processor struct {
	TypeName
	//

}

// 构造函数（支持多个 wroker 并发处理）
func NewProcessor(workerNum int) *Processor {
	return &Processor{}
}

// 核心处理方法
func (p *Processor) Process(req UpdateRequest) error {
	// 初始化协程
	return nil
}

// 关闭方法
func (p *Processor) Shutdown() {

}

// 其他函数自行新增
