package controllers

import (
    "github.com/astaxie/beego"
    "docker-redis-sample/cache"
)

type MainController struct {
    beego.Controller
}

func (c *MainController) Get() {
    cacheManager := cache.GetCacheManager()

    cacheManager.Put("key", "test", 0)

    c.Data["Website"] = "beego.me"
    c.Data["Email"] = cacheManager.Get("key")
    c.TplName = "index.tpl"
}
