package db

//数据库配置.
type Config struct {
	Host     string `json:"host"`
	Dbname   string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
	Sslmode  string `json:"ssl"`
}

//存储服务器.
type StorageServer struct {
	Id        int `db:"id, bigserial, primary key, not null" json:"id"`
	Ip        string `db:"ip, text, not null" json:"ip"`
	Status    int `db:"status, int, not null" json:"status"`
	CreatedAt int `db:"created_at, int, not null"`
}

//待同步.
type WaitSync struct {
	Id        int `db:"id, bigserial, primary key, not null" json:"id"`
	Path      string `db:"path, text, not null" json:"path"`
	TargetIp  string `db:"target, text, not null" json:"target_ip"`
	SyncToIp  string `db:"sync_to_ip, text, not null" json:"sync_to_ip"`
	Status    int `db:"status, int, not null" json:"status"`
	CreatedAt int `db:"created_at, int, not null" json:"created_at"`
}

type SyncMessage struct {
	Id        int `db:"id, bigserial, primary key, not null"`
	Path      string `db:"path, text, not null"`
	SyncIp    string `db:"sync_ip, text, not null"`
	Status    int `db:"status, int, not null" json:"status"`
	CreatedAt int `db:"created_at, int, not null" json:"created_at"`
}

const (
	SYNC_MESSAGE_WAIT_SYNC    = 0
	SYNC_MESSAGE_SYNC_SUCCESS = 1
)
