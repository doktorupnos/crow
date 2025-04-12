-- +goose Up
-- +goose StatementBegin
CREATE TABLE follows (
    follower UUID NOT NULL,
    followee UUID NOT NULL,
    PRIMARY KEY (follower, followee),
    FOREIGN KEY (follower) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (followee) REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE follows;
-- +goose StatementEnd
