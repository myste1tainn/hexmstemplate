package main

import (
	"github.com/myste1tainn/hexmstemplate/httpserv"
	"github.com/myste1tainn/hexmstemplate/infrastructure"
)

func init() {
	infrastructure.InitConfig()
}
func main() {
	infrastructure.InitLogger()
	httpserv.Run()
}
