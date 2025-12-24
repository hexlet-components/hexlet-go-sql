-- name: CreateCourse :one
INSERT INTO courses (title, description)
VALUES ($1, $2)
RETURNING id, title, description, created_at;

-- name: GetCourse :one
SELECT id, title, description, created_at
  FROM courses
 WHERE id = $1;

-- name: ListCourses :many
SELECT id, title, description, created_at
  FROM courses
 ORDER BY created_at DESC
 LIMIT $1 OFFSET $2;
