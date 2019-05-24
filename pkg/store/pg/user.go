package pg

import (
	"time"

	"github.com/go-pg/pg"

	"gorum/pkg/auth"
)

type user struct {
	ID           int64
	Name         string
	Email        string
	Secret       string
	FullName     string
	TagLine      string
	Role         int
	Respect      int
	RegisteredAt time.Time
	UpdatedAt    time.Time
}

func (m *user) From(u *auth.User) {
	m.ID = u.ID
	m.Name = string(u.Name)
	m.Email = string(u.Email)
	m.Secret = u.Secret
	m.FullName = u.FullName
	m.TagLine = u.TagLine
	m.Role = int(u.Role)
	m.Respect = u.Respect
	m.RegisteredAt = u.RegisteredAt
	m.UpdatedAt = u.UpdatedAt
}

func (m *user) To(u *auth.User) {
	u.ID = m.ID
	u.Name = auth.UserName(m.Name)
	u.Email = auth.Email(m.Email)
	u.Secret = m.Secret
	u.FullName = m.FullName
	u.TagLine = m.TagLine
	u.Role = auth.Role(m.Role)
	u.Respect = m.Respect
	u.RegisteredAt = m.RegisteredAt
	u.UpdatedAt = m.UpdatedAt
}

// NewUserStorer creates a new Postgres auth.UserStorer.
func NewUserStorer(db *pg.DB) auth.UserStorer {
	return &userStore{db: db}
}

type userStore struct {
	db *pg.DB
}

func (us *userStore) Find(u *auth.User, id int64) error {
	return us.findOne(u, "id = ?", id)
}

func (us *userStore) FindByName(u *auth.User, name auth.UserName) error {
	return us.findOne(u, "name = ?", string(name))
}

func (us *userStore) ExistsByName(name auth.UserName) (bool, error) {
	return us.exists("name = ?", string(name))
}

func (us *userStore) FindByEmail(u *auth.User, email auth.Email) error {
	return us.findOne(u, "email = ?", string(email))
}

func (us *userStore) ExistsByEmail(email auth.Email) (bool, error) {
	return us.exists("email = ?", string(email))
}

func (us *userStore) Persist(u *auth.User) error {
	model := user{}
	model.From(u)

	if model.ID == 0 {
		err := us.db.Insert(&model)
		if err != nil {
			return err
		}

		u.ID = model.ID
		return nil
	}

	return us.db.Update(&model)
}

func (us *userStore) findOne(u *auth.User, where string, params ...interface{}) error {
	model := user{}
	model.From(u)

	err := us.db.Model(&model).Where(where, params...).First()
	if err != nil {
		if err == pg.ErrNoRows {
			return auth.ErrNoUser
		}

		return err
	}

	model.To(u)

	return nil
}

func (us *userStore) exists(where string, params ...interface{}) (bool, error) {
	return us.db.Model(&user{}).Where(where, params...).Exists()
}
