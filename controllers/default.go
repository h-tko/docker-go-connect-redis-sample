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

    c.Data["NowCache"] = cacheManager.Get("cache")
    c.TplName = "index.tpl"
}

func (c *MainController) Post() {
    cacheManager := cache.GetCacheManager()

    item := c.GetString("Item")

    cacheManager.Put("cache", item, 0)

    c.Data["NowCache"] = item

    c.TplName = "index.tpl"
}
