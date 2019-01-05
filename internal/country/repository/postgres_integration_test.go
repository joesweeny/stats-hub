package repository

import (
	"testing"
	"fmt"
	"database/sql"
	"github.com/joesweeny/statshub/internal/model"
	"github.com/joesweeny/statshub/internal/config"
	"github.com/satori/go.uuid"
	"time"
	_ "github.com/lib/pq"
)

func TestInsert(t *testing.T) {
	conn, cleanUp := getConnection(t)
	repo := NewPostgresCountryRepository(conn)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()
		for i := 1; i < 4; i++ {
			var id = uuid.Must(uuid.NewV4(), nil)
			c := newCountry(id)

			if err := repo.Insert(c); err != nil {
				t.Errorf("Error when inserting record into the database: %s", err.Error())
			}

			row := conn.QueryRow("select count(*) from country")

			var count int

			if err := row.Scan(&count); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			if i != count {
				t.Fatalf("Expected %d, got %d", i, count)
			}
		}
	})

	t.Run("returns error when ID primary key violates unique constraint", func(t *testing.T) {
		t.Helper()
		defer cleanUp()
		c := newCountry(uuid.UUID{})

		if err := repo.Insert(c); err != nil {
			t.Errorf("Test failed, expected nil, got %s", err)
		}

		if e := repo.Insert(c); e == nil {
			t.Fatalf("Test failed, expected %s, got nil", e)
		}
	})

	conn.Close()
}

func TestUpdate(t *testing.T) {
	conn, cleanUp := getConnection(t)
	repo := NewPostgresCountryRepository(conn)

	t.Run("modifies existing record", func(t *testing.T) {
		t.Helper()
		defer cleanUp()
		var id = uuid.Must(uuid.NewV4(), nil)
		c := newCountry(id)

		if err := repo.Insert(c); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		s := "Scotland"
		c.Name = s

		if err := repo.Update(c); err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		r, err := repo.GetById(c.ID)

		if err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		got := r.Name
		want := s

		if got != want {
			t.Fatalf("got %s, want %s", got, want)
		}
	})

	t.Run("returns error if record does not exist", func (t *testing.T) {
		t.Helper()
		defer cleanUp()
		c := newCountry(uuid.NewV4())

		err := repo.Update(c)

		if err == nil {
			t.Fatalf("Test failed, expected nil, got %v", err)
		}
	})
}

var db = config.GetConfig().DB

func getConnection(t *testing.T) (*sql.DB, func()) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Name)

	db, err := sql.Open(db.Driver, psqlInfo)

	if err != nil {
		panic(err)
	}

	return db, func() {
		_, err := db.Exec("delete from country")
		if err != nil {
			t.Fatalf("Failed to clear database. %s", err.Error())
		}
	}
}

func newCountry(u uuid.UUID) model.Country {
	c := model.Country{
		u,
		1,
		"England",
		"Europe",
		"ENG",
		time.Now(),
		time.Now(),
	}

	return c
}
