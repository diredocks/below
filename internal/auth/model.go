package auth

type User struct {
	ID       int64  ``
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"-"`
}
