package main

import (
	_ "docker-redis-sample/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

