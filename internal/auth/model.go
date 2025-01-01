package auth

type User struct {
	ID       int64  `xorm:"'id' pk autoincr"`
	Username string `xorm:"not null" validate:"required" json:"username"`
	Password string `xorm:"not null" validate:"required" json:"-"`
}
