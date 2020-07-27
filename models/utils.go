package models

import (
	"bytes"
	"database/sql"
	"os/exec"
	"strconv"
	"strings"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db * sql.DB
	err error
)

func GetStampDiff(t string) (diff int64) {
	now := time.Now()
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", t, loc)
	logStamp := theTime.Unix()
	theStamp := now.Unix()

	return theStamp - logStamp
}

func RunCmd(command string) (error, string, string) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stderr, cmd.Stdout = &stderr, &stdout
	err := cmd.Run()

	return err, stdout.String(), stderr.String()
}

func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func DoMysql() (n int64, errInfo string, tag bool) {
	Db, err = sql.Open("mysql", "ops:wow-trend.com@tcp(120.78.161.44:3306)/wow")
	//Db, err = sql.Open("mysql", "root:123456@tcp(192.168.1.16:3306)/test")
	if err != nil {
		errInfo = "连接数据库失败。"
		return 0, errInfo, false
	//	panic(err.Error())
	}

	sqlStr := "delete from users"
	ret, err := Db.Exec(sqlStr)
	if err != nil {
		errInfo = "执行失败。"
		return 0, errInfo, false
	}
	n, err = ret.RowsAffected()
	if err != nil {
		errInfo = "获取结果失败。"
		return 0, errInfo, false
	}
	errInfo = "清除成功。"

	return n, errInfo, true
}