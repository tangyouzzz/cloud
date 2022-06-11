package main

import(
	// "github.com/gin-gonic/gin"
	// "net/http"
	"main/dao"
	// "fmt"
	// "main/controller"
	"main/routers"
	// "github.com/dgrijalva/jwt-go"
	// "time"
	// "errors"

	// "path"
	// "os"
	// "iouil"
)

func main(){
	dao.Initdb()
	r :=routers.SetupRouter()
	r.Run(":8848")
}

