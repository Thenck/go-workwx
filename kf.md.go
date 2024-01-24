// Code generated by sdkcodegen; DO NOT EDIT.

package workwx

// KfAccount 客服账号
type KfAccount struct {
	// OpenKfID 客服账号ID
	OpenKfID string `json:"open_kfid"`
	// Name 客服名称
	Name string `json:"name"`
	// Avatar 客服头像URL
	Avatar string `json:"avatar"`
	// ManagePrivilege 当前调用接口的应用身份，是否有该客服账号的管理权限（编辑客服账号信息、分配会话和收发消息）。组件应用不返回此字段
	ManagePrivilege bool `json:"manage_privilege,omitempty"`
}

// KfServicer 客服接待人员
type KfServicer struct {
	// UserID 接待人员的userid。第三方应用获取到的为密文userid，即open_userid
	UserID string `json:"userid,omitempty"`
	// Status 接待人员的接待状态。0:接待中,1:停止接待。
	Status int `json:"status"`
	// StopType 接待人员的接待状态为「停止接待」的子类型。0:停止接待,1:暂时挂起
	StopType int `json:"stop_type"`
	// DepartmentID 接待人员部门的id
	DepartmentID int64 `json:"department_id,omitempty"`
}

// KfServicerResult 客户群列表数据
type KfServicerResult struct {
	// UserID 接待人员的userid
	UserID string `json:"userid,omitempty"`
	// DepartmentID 接待人员部门的id
	DepartmentID int64 `json:"department_id,omitempty"`
	// ErrCode 该条记录的结果
	ErrCode int64 `json:"errcode"`
	// ErrMsg 结果信息
	ErrMsg string `json:"errmsg"`
}

// KfServiceState 客服会话状态
//
// 0 未处理 新会话接入
// 1 由智能助手接待
// 2 待接入池排队中
// 3 由人工接待
// 4 已结束/未开始
type KfServiceState int

const (
	// KfServiceStateUntreated 未处理 新会话接入
	KfServiceStateUntreated KfServiceState = iota
	// KfServiceStateRobotReception 由智能助手接待
	KfServiceStateRobotReception
	// KfServiceStateInQueue 待接入池排队中
	KfServiceStateInQueue
	// KfServiceStateManualReception 由人工接待
	KfServiceStateManualReception
	// KfServiceStateFinished 已结束/未开始
	KfServiceStateFinished
)
