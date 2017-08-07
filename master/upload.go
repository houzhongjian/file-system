package main

import (
	"net/http"
	"github.com/houzhongjian/file-system/system"
	"log"
	"io/ioutil"
	"bytes"
	"github.com/houzhongjian/file-system/db"
	"github.com/houzhongjian/file-system/conf"
	"github.com/houzhongjian/file-system/db/model"
	"time"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	path, err := system.FileStoredPath()
	if err != nil {
		log.Printf("%+v\n", err)
		http.Error(w, "服务器异常", http.StatusInternalServerError)
		return
	}

	by, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	var buffer bytes.Buffer
	buffer.WriteString(path)
	buffer.WriteString(r.Header.Get("Name"))
	file := buffer.String()

	if !system.CreateFile(file, by) {
		w.Write([]byte("上传失败"))
		return
	}

	//发送消息到队列中，通知备份服务器从主服务器备份文件.
	msg := db.SyncMessage{
		Path:file,
		SyncIp:conf.MASTER_UPLOAD_GW,
		Status:db.SYNC_MESSAGE_WAIT_SYNC,
		CreatedAt:int(time.Now().Unix()),
	}
	if err = model.InsertSyncMessage(msg); err != nil {
		log.Printf("%+v\n", err)
		w.Write([]byte("上传失败"))
		return
	}
	//返回主服务器上的文件路径给用户.

	log.Println("文件上传成功:", file)
	w.Write([]byte(conf.MASTER_UPLOAD_GW+file))
}
