package main
import(
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/router"
)
func main()  {
	r:=gin.Default()
	router.Init(r)
	_=r.Run(":8081")
}
