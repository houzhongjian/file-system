package model

import (
	"fmt"
	"github.com/houzhongjian/file-system/db"
	"log"
)

//添加同步数据的消息.
func InsertSyncMessage (msg db.SyncMessage) error {
	esql := fmt.Sprintf(`INSERT INTO %s(path, sync_ip, status, created_at) VALUES ('%s', '%s', %d, %d)`,
		db.TABLR_SYNC_MESSAGE,
		msg.Path,
		msg.SyncIp,
		msg.Status,
		msg.CreatedAt,
	)

	st, err := db.Db.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return err
	}
	_, err = st.Exec()
	if err != nil {
		log.Printf("%+v\n", err)
		return err
	}

	return nil
}
