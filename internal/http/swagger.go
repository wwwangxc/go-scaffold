/*
Package app 生成swagger文档

文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format

使用方式：
	go get -u github.com/swaggo/swag/cmd/swag
	swag init --generalInfo ./internal/http/swagger.go --output ./internal/http/swagger */

package http

// @title go-scaffold/
// @version 1.0.0
// @description 网关HTTP服务
// @schemes http
