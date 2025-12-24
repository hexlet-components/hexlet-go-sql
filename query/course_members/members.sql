-- name: JoinCourse :exec
INSERT INTO course_members(user_id, course_id)
VALUES ($1, $2);

-- name: ListCourseMembers :many
SELECT user_id, course_id, joined_at
  FROM course_members
 WHERE course_id = $1
 ORDER BY joined_at;
