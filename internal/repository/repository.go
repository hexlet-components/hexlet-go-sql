package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	dbpkg "github.com/hexlet-components/hexlet-go-sql/internal/db"
)

// User описывает сущность пользователя.
type User struct {
	ID    int64   `json:"id"`
	Email string  `json:"email"`
	Name  *string `json:"name,omitempty"`
}

// Course описывает курс.
type Course struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// Member описывает участника курса.
type Member struct {
	UserID   int64  `json:"user_id"`
	CourseID int64  `json:"course_id"`
	JoinedAt string `json:"joined_at"`
}

// Repo предоставляет методы для работы с БД.
type Repo struct {
	db      *sql.DB
	queries *dbpkg.Queries
}

// New создаёт репозиторий.
func New(db *sql.DB) *Repo {
	return &Repo{db: db, queries: dbpkg.New(db)}
}

// CreateUser создаёт пользователя.
func (r *Repo) CreateUser(ctx context.Context, email string, name *string) (User, error) {
	u, err := r.queries.CreateUser(ctx, dbpkg.CreateUserParams{
		Email: email,
		Name:  toNullString(name),
	})
	if err != nil {
		return User{}, err
	}
	return fromDBUser(u), nil
}

// CreateCourse создаёт курс.
func (r *Repo) CreateCourse(ctx context.Context, title, description string) (Course, error) {
	c, err := r.queries.CreateCourse(ctx, dbpkg.CreateCourseParams{
		Title:       title,
		Description: toNullString(&description),
	})
	if err != nil {
		return Course{}, err
	}
	return Course{
		ID:          c.ID,
		Title:       c.Title,
		Description: nullStringToString(c.Description),
	}, nil
}

// JoinCourse добавляет пользователя в курс и создаёт заказ.
func (r *Repo) JoinCourse(ctx context.Context, tx *sql.Tx, userID, courseID int64) error {
	qtx := dbpkg.New(tx)
	if err := qtx.JoinCourse(ctx, dbpkg.JoinCourseParams{
		UserID:   userID,
		CourseID: courseID,
	}); err != nil {
		return fmt.Errorf("add member: %w", err)
	}

	return nil
}

// ListCourseMembers возвращает JSON с участниками.
func (r *Repo) ListCourseMembers(ctx context.Context, courseID int64) ([]byte, error) {
	items, err := r.queries.ListCourseMembers(ctx, courseID)
	if err != nil {
		return nil, err
	}
	members := make([]Member, len(items))
	for i, item := range items {
		members[i] = Member{
			UserID:   item.UserID,
			CourseID: item.CourseID,
			JoinedAt: item.JoinedAt.Format(time.RFC3339),
		}
	}
	return json.MarshalIndent(members, "", "  ")
}

func toNullString(value *string) sql.NullString {
	if value == nil || *value == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *value, Valid: true}
}

func fromDBUser(u dbpkg.User) User {
	return User{
		ID:    u.ID,
		Email: u.Email,
		Name:  nullStringToPtr(u.Name),
	}
}

func nullStringToPtr(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	val := ns.String
	return &val
}

func nullStringToString(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}
