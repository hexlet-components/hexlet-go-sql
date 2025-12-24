package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/joho/godotenv"

	"github.com/hexlet-components/hexlet-go-sql/internal/config"
	"github.com/hexlet-components/hexlet-go-sql/internal/repository"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()
	var (
		cmd      = flag.String("cmd", "user-create", "command to run")
		email    = flag.String("email", "", "user email")
		name     = flag.String("name", "", "user name")
		course   = flag.String("course", "", "course title")
		userID   = flag.Int64("user", 0, "user id")
		courseID = flag.Int64("course-id", 0, "course id")
	)
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	repo := repository.New(db)

	switch *cmd {
	case "user-create":
		user, err := repo.CreateUser(ctx, *email, strPtr(*name))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", user)
	case "course-create":
		c, err := repo.CreateCourse(ctx, *course, "")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", c)
	case "join-course":
		err := repository.WithTx(ctx, db, func(tx *sql.Tx) error {
			return repo.JoinCourse(ctx, tx, *userID, *courseID)
		})
		if err != nil {
			log.Fatal(err)
		}
		data, _ := repo.ListCourseMembers(ctx, *courseID)
		fmt.Println(string(data))
	default:
		log.Fatalf("unknown command: %s", *cmd)
	}
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
