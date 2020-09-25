// HTTP 服务入口

package main

import (
	"fmt"
	"go-scaffold/internal/http"
	"os"
)

func main() {
	fmt.Printf("processID: %d\n", os.Getppid())
	http.Run()
}
