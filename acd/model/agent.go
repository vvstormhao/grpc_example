package model

// Agent agents表结构
type Agent struct {
	SoftDelBase
	AppID       string `gorm:"size:191;not null"`
	AgentID     string `gorm:"size:191;not null;index"` // id@appid 唯一索引
	Role        string `gorm:"size:191;default normal"` // 角色 normal 普通坐席 admin 坐席管理员
	WorkerID    string `gorm:"size:191"`
	WorkMod     int    `gorm:"size:1;default:0"`          // 坐席工作模式 0 IP话机模式 1 手机模式
	PhoneNumber string `gorm:"size:11"`                   // 坐席手机号
	State       int    `gorm:"size:4;not null;default:0"` // 坐席状态 0 离线 1 空闲 2 忙碌 3 小休 4 准备
	VoipNumber  string `gorm:"size:16"`
	Password    string `gorm:"size:16"`
	ExtState    int    `gorm:"size:4;not null;default:0"` // 分机状态 0 offline 1 online
}
