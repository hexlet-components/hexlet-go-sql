package repository

import (
    "context"
    "database/sql"
    "encoding/json"
    "fmt"
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
    db *sql.DB
}

// New создаёт репозиторий.
func New(db *sql.DB) *Repo {
    return &Repo{db: db}
}

// CreateUser создаёт пользователя.
func (r *Repo) CreateUser(ctx context.Context, email string, name *string) (User, error) {
    res, err := r.db.ExecContext(ctx,
        `INSERT INTO users(email, name) VALUES(?, ?)`,
        email, name,
    )
    if err != nil {
        return User{}, err
    }
    id, _ := res.LastInsertId()
    return User{ID: id, Email: email, Name: name}, nil
}

// CreateCourse создаёт курс.
func (r *Repo) CreateCourse(ctx context.Context, title, description string) (Course, error) {
    res, err := r.db.ExecContext(ctx,
        `INSERT INTO courses(title, description) VALUES(?, ?)`,
        title, description,
    )
    if err != nil {
        return Course{}, err
    }
    id, _ := res.LastInsertId()
    return Course{ID: id, Title: title, Description: description}, nil
}

// JoinCourse добавляет пользователя в курс и создаёт заказ.
func (r *Repo) JoinCourse(ctx context.Context, tx *sql.Tx, userID, courseID int64, amount int64) error {
    if _, err := tx.ExecContext(ctx,
        `INSERT INTO course_members(user_id, course_id) VALUES(?, ?)`,
        userID, courseID,
    ); err != nil {
        return fmt.Errorf("add member: %w", err)
    }

    if _, err := tx.ExecContext(ctx,
        `INSERT INTO orders(user_id, course_id, amount_cents) VALUES(?, ?, ?)`,
        userID, courseID, amount,
    ); err != nil {
        return fmt.Errorf("create order: %w", err)
    }

    return nil
}

// ListCourseMembers возвращает JSON с участниками.
func (r *Repo) ListCourseMembers(ctx context.Context, courseID int64) ([]byte, error) {
    rows, err := r.db.QueryContext(ctx,
        `SELECT user_id, course_id, joined_at FROM course_members WHERE course_id = ? ORDER BY joined_at`,
        courseID,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    members := []Member{}
    for rows.Next() {
        var m Member
        if err := rows.Scan(&m.UserID, &m.CourseID, &m.JoinedAt); err != nil {
            return nil, err
        }
        members = append(members, m)
    }
    return json.MarshalIndent(members, "", "  ")
}
