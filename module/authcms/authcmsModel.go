package authcms

type User struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"-"`
	CreatedAt int64  `db:"created_at" json:"-"`
}