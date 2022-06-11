package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/juanmachuca95/conductores_go/domains/models"
	ri "github.com/juanmachuca95/conductores_go/interface/repository"
	"github.com/juanmachuca95/conductores_go/testutils"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func setup(t *testing.T) (r repository.AuthRepository, mock sqlmock.Sqlmock, teardown func()) {
	db, mock, err := testutils.NewDBMock(t)
	if err != nil {
		panic(err.Error())
	}

	r = ri.NewAuthRepository(db)
	return r, mock, func() {
		db.Close()
	}
}

func TestAuhtRepository_Login(t *testing.T) {
	r, mock, teardown := setup(t)
	defer teardown()

	l := models.Login{Email: "machucajuangabriel@gmail.com", Password: "123456"}
	u := models.User{Id: 1, Name: "Juan", Email: l.Email, Created_at: "2022-06-11", Updated_at: "2022-06-11"}
	pass, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	passHashed := string(pass)

	cases := map[string]struct {
		arrange func(t *testing.T)
		asserts func(t *testing.T, udb *models.User, u *models.User, err error)
	}{
		"sql_mocked": {
			arrange: func(t *testing.T) {
				mock.ExpectPrepare("SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?")).WithArgs(l.Email).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).AddRow(u.Id, u.Name, u.Email, passHashed, u.Created_at, u.Updated_at))
			},
			asserts: func(t *testing.T, udb *models.User, u *models.User, err error) {
				assert.Nil(t, err)
				assert.Equal(t, udb, u)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			tt.arrange(t)

			user, err := r.Login(&l)

			tt.asserts(t, &u, user, err)
		})
	}
}
