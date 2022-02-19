package main

import (
	"github.com/perfect6566/bingpaper/cmd"
	"github.com/robfig/cron"
)

const paperapiurl = "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN"
func main()  {
	c:=cron.New()
	c.AddFunc("0 50 18 * * *", func() {
		cmd.Fetchpapers(paperapiurl)
	})
	c.Run()
}