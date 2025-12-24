-- name: CreateLesson :one
INSERT INTO lessons (course_id, title, description)
VALUES ($1, $2, $3)
RETURNING id, course_id, title, description, created_at;

-- name: ListLessonsByCourse :many
SELECT id, course_id, title, description, created_at
  FROM lessons
 WHERE course_id = $1
 ORDER BY created_at;
