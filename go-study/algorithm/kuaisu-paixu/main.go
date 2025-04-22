package main

type RecordExceptionInfo struct {
	UserId              string   `json:"user_id"`
	BusinessId          string   `json:"business_id"` //保持和错误码一致，例如account：1，鉴权2 ，通知：3
	URL                 string   `json:"url"`
	Email               string   `json:"email"`
	StationId           uint64   `json:"station_id"`
	TraceId             string   `json:"trace_id"`
	ErrorCode           int      `json:"error_code"`
	ErrorMsg            string   `json:"error_msg"`
	PointNumber         int      `json:"point_number"`       //Repo层 ，service层，中间层，第三层调用层，
	ExceptionCategory   int      `json:"exception_category"` // 异常分类/紧急度 ， 按照系统界别，中间件级别，数据库级别，Redis级别，业务级别，
	Details             string   `json:"details"`
	InputParams         string   `json:"input_params"`
	OutputParams        string   `json:"output_params"`
	HasPermission       bool     `json:"has_permission"`       // 权限是否拥有
	RequiredPermissions []string `json:"required_permissions"` // 需要什么权限
	Ctime               uint64   `json:"ctime"`
	Mtime               uint64   `json:"mtime"`
}

type RecordExceptionHandleInfo struct {
	TraceId       string `json:"trace_id"`
	URL           string `json:"url"`
	ErrorCode     int    `json:"error_code"`
	ErrorMsg      string `json:"error_msg"`
	ErrorDetails  string `json:"error_details"`
	MethodType    int    `json:"method_type"` // 1：get 2：post
	RequestBody   string `json:"request_body"`
	ReqeustHeader string `json:"reqeust_Header"`
	ResponseInfo  string `json:"response_info"`
	Ctime         uint64 `json:"ctime"`
	Mtime         uint64 `json:"mtime"`
}

func main() {
}
