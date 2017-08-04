package main

import (
	"os"
	"github.com/houzhongjian/file-system/system"
	"log"
)

//初始化存储系统.
func InitSystem() {
	_, err := os.Stat(system.FILE_SYSTEM_PATH)
	if err != nil {
		log.Printf("监测到文件系统存储根目录不存在，开始执行自动创建... %+v\n", err)
		if err = os.MkdirAll(system.FILE_SYSTEM_PATH, os.ModePerm); err != nil {
			log.Printf("初始化失败:%+v\n", err)
			return
		}
		log.Printf("文件系统存储根目录创建成功")
	}
}
