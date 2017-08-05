package db

type User struct {
	Id       int64 `db:"id, int, primary key" json:"id"`
	Username string `db:"username, text, not null" json:"username"`
	Password string `db:"password, text, not null" json:"password"`
}

type Account struct {

}