// models包定义了数据模型
package models

// User 表示一个用户实体
type User struct {
	ID       string `json:"id"`       // 用户唯一标识符
	PhoneNo  string `json:"phone_no"` // 用户手机号
	Password string `json:"password"` // 用户密码
	Name     string `json:"name"`     // 用户姓名
	Remark   string `json:"备注"`      // 备注信息
}