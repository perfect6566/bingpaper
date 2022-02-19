package server
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Serverstart() {
	r:=gin.Default()
	r.StaticFS("/",http.Dir("./papers"))
	r.Run(":8888")
}
