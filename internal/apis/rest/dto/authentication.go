package dto

// Login paramter
type Login struct {
	Username string `json:"username" binding:"required" example:"admin"`  // 用户名
	Password string `json:"password" binding:"required" example:"123456"` // 密码
}
