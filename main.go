package main

import (
	"./models"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Msg struct {
	Time string `json:"time"`
	Brand string `json:"brand"`
	Folder string `json:"folder"`
}

func GetRes(c *gin.Context)  {
	now := time.Now()
	logDate := fmt.Sprintf(now.Format("2006_01_02") + ".log")
	logPath := path.Join("/data/wwwroot/wow-trend-com/wow-admin/Runtime/Logs/task/UploadAutoTask", logDate)
	err, out, _ := models.RunCmd("tail " + logPath)
	if err != nil {
		fmt.Println(err)
	}

	//cc := new(models.ClientConfig)
	//cc.SSHConnect("192.168.1.8", "root", "wow-trend.com", 45184)
	//cc.RunCMD("tail " + logPath)

	logs := strings.Split(strings.TrimSpace(out), "\n")
	Time := strings.Split(logs[len(logs)-1], ": ")[0]
	diff := models.GetStampDiff(Time)
	//fmt.Println(diff)
	if diff > 300 {
		c.HTML(http.StatusOK, "clean.html", "当前无任务或进程卡死！")
		return
	}

	var Flag bool
	var msg *Msg
	var res []Msg

	for _, log := range logs {
		if strings.Contains(log, "/data/work/upload") {
			Flag = true
			Time := strings.Split(log, ": ")[0]
			Brand, Folder := strings.Split(log, "/")[4], strings.Split(log, "/")[5]
			msg = &Msg{Time: Time, Brand: Brand, Folder: Folder}
			res = append(res, *msg)
		}
		if strings.Contains(log, "执行错误") {
			Flag = true
			Time := strings.Split(log, ": ")[0]
			Brand := "执行错误"
			err := strings.Split(log, "\"")
			errInfo := err[len(err)-2]
			byte, _ := models.UnescapeUnicode([]byte(errInfo))
			Folder := strings.Replace(string(byte), "\\", "", -1)
			msg = &Msg{Time: Time, Brand: Brand, Folder: Folder}
			res = append(res, *msg)
		}
	}

	if !Flag {
		c.HTML(http.StatusOK, "clean.html", "当前无任务或进程卡死！")
		return
	}

	c.HTML(http.StatusOK, "view.html", res)
}

func Clean(c *gin.Context)  {
	n, errInfo, tag := models.DoMysql()
	msg := fmt.Sprintf("成功清理%v条缓存!", n)
	if tag {
		c.HTML(http.StatusOK, "clean.html", msg)
	} else {
		c.HTML(http.StatusOK, "clean.html", errInfo)
	}
}

func Restart(c *gin.Context)  {
	//cc := new(models.ClientConfig)
	//cc.SSHConnect("192.168.1.8", "root", "wow-trend.com", 45184)
	//cc.RunCMD("bash /root/scripts/banmian.sh")
	n, errInfo, tag := models.DoMysql()
	var msg string
	if tag {
		msg = fmt.Sprintf("，并清理%v条缓存！", n)
	} else {
		msg = "，" + errInfo
	}

	err, out, e := models.RunCmd("bash /root/scripts/banmian.sh")
	if err != nil {
		fmt.Println(err)
	}
	if e != "" {
		msg = "重启失败，请重新尝试！"
	} else {
		msg = "重启成功" + msg
	}
	fmt.Println(out, e)
	c.HTML(http.StatusOK, "restart.html", msg)
}

func loadTemplates(templateDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templateDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templateDir + "/includes/*")
	if err != nil {
		panic(err.Error())
	}
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func main() {
	r := gin.Default()

	r.Static("/static", "static")
	r.HTMLRender = loadTemplates("./templates")
	//r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/view", GetRes)
	r.GET("/clean", Clean)
	r.GET("/restart", Restart)

	r.Run(":9090")
}