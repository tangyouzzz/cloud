package controller

import(
	// "errors"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"os"
	"main/models"

)


// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type UserInfo struct {
	Username string
	Password string
}

const TokenExpireDuration = time.Hour * 2
var MySecret = []byte("嘉然今天吃什么")




func Filedelete(c *gin.Context){
	username :=c.MustGet("username").(string)
	fname := c.Request.FormValue("name")
	id,path :=models.FindFilebyname(username,fname)
	if id==0{
		c.JSON(http.StatusOK,gin.H{
			"code":2001,
			"msg":"wenjianbucunzai",
		})
	}else{
		models.DeleteFile(id)
		os.Remove(path+fname)
		c.JSON(http.StatusOK,gin.H{
			"code":2000,
			"msg":"shanchuchenggong",
		})	
	}

}

func Filedownload(c *gin.Context)  {
	username :=c.MustGet("username").(string)
	fname := c.Request.FormValue("name")
	_,path :=models.FindFilebyname(username,fname)
	if path ==""{
		c.JSON(http.StatusOK,gin.H{
			"code":2001,
			"msg":"文件不存在",
		})
		 
	}else{
		c.Header("Access-Control-Expose-Headers","Content-Disposition")
		c.Header("Content-Type","application/octet-stream")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Cache-Control", "no-cache")
		c.Header("response-type","blob")
		c.File(path+fname)
	}

	
}


func Fileupload(c *gin.Context){
	username :=c.MustGet("username").(string)
	file,err :=c.FormFile("file")
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":2001,
			"msg":"上传文件出错",
		})
		return
	}
	id,_ :=models.FindFilebyname(username,file.Filename)
	if id!=0{
		c.JSON(http.StatusOK,gin.H{
			"code":2002,
			"msg":"文件已存在",
		})
	}else {
		_ = os.Mkdir("./file/"+username,0777)
		uploaddate := time.Now().Format("2006-01-02 15:04:05")
		c.SaveUploadedFile(file,"./file/"+username+"/"+file.Filename,)
		models.CreateFile(file.Filename,"./file/"+username+"/",file.Size,username,uploaddate)
		c.JSON(http.StatusOK,gin.H{
			"code":2000,
			"msg":"上传文件成功",
		})
	}


}




// chaxunyonghusuoshudewenjian
func Filefind(c *gin.Context){
	username :=c.MustGet("username").(string)
	fil :=models.FindFIle(username)
	c.JSON(http.StatusOK,gin.H{
		"file": fil ,
	})
}

func AuthHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "check your email or password",
		})
		return
	}

	// 校验用户名和密码是否正确
	_,p :=models.GetUser(user.Username)
	if p==""{
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "check your email or password",
		})
		return

	}
	if user.Password == p{
		// 生成Token
		tokenString, _ := GenToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "check your email or password",
	})
	return
}

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "tangyou",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}


func HomeHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
	// c.HTML(200,"index.html",nil)
}


