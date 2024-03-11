package storage

import (
	"database/sql"
	"fmt"
	"lms_back/config"

	_ "github.com/lib/pq"
)

type Store struct {
	DB     *sql.DB
	Branch branchRepo
	Teacher TeacherRepo
	Student StudentRepo
	Group   GroupRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}
	NewTeacher := NewTeacher(db)
	NewBranch  := NewBranch(db)
	NewStudent := NewStudent(db)
	NewGroup   := NewGroup(db)
	return Store{
		DB:     db,
		Branch: NewBranch,
		Teacher: NewTeacher,
		Student: NewStudent,
		Group: NewGroup,
	}, nil

}


