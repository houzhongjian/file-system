package system

import (
	"fmt"
	"time"
	"os"
	"log"
	"io/ioutil"
)

//文件系统路径.
const FILE_SYSTEM_PATH = "/file/"

//文件存放路径演示:/assets/2017/08/03/12312112423.jpg.
//获取文件存储的路径.
func FileStoredPath() (string, error) {
	path := fmt.Sprintf(FILE_SYSTEM_PATH+"%d/%d/%d/", time.Now().Year(), time.Now().Month(), time.Now().Day())
	status := IsExist(path)
	if !status {
		//创建文件的存储路径.
		log.Println("开始自动创建文件存储目录:",path)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Printf("目录创建失败:%+v\n", err)
			return "", err
		}
		log.Println("目录创建成功:", path)
	}

	return path, nil
}

//判断路径是否存在.
func IsExist(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("监测到文件存储目录不存在...%+v\n", err)
		return false
	}
	fileInfo, err := f.Stat()
	if err != nil {
		log.Printf("监测到文件存储目录不存在...%+v\n", err)
		return false
	}

	return fileInfo.IsDir()
}

//创建文件.
func CreateFile(path string, by []byte) bool {
	if err := ioutil.WriteFile(path, by, os.ModePerm); err != nil {
		log.Printf("%+v\n", err)
		return false
	}
	return true
}
