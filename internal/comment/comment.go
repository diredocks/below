package comment

type CommentStatus string

const (
	StatusSent    CommentStatus = "Sent"
	StatusRemoved CommentStatus = "Removed"
)

type Comment struct {
	ID        int64         `xorm:"'id' pk autoincr"`
	Site      string        `xorm:"'site' not null" validate:"required,url" json:"site"`
	Page      string        `xorm:"'page' not null" validate:"required,uri" json:"page"`
	Nickname  string        `xorm:"not null default 'Anonymous'" validate:"required" json:"nickname"`
	Email     string        `xorm:"not null" validate:"omitempty,email" json:"email"`
	Content   string        `xorm:"text not null" validate:"required" json:"content"`
	CreatedAt string        `xorm:"created"`
	Status    CommentStatus `xorm:"not null default 'Sent'"`
}

type CommentQuery struct {
	Site string `validate:"required,url" json:"site"`
	Page string `validate:"required,uri" json:"page"`
}
