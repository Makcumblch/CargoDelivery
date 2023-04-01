package cargodelivery

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" db:"password_hash"`
	Salt     string `json:"-" db:"salt"`
}
