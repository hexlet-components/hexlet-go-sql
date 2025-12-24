-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS course_members (
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id BIGINT NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, course_id)
);
CREATE INDEX IF NOT EXISTS idx_course_members_course_id ON course_members(course_id);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS course_members;
