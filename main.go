package main
import (
	"github.com/perfect6566/bingpaper/cmd"
	"github.com/perfect6566/bingpaper/server"
	"github.com/robfig/cron"
)

const paperapiurl = "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN"
func main()  {
	go server.Serverstart()
	c:=cron.New()
	c.AddFunc("0 50 18 * * *", func() {
		cmd.Fetchpapers(paperapiurl)
	})
	c.Run()
}