-- +goose Up
-- +goose StatementBegin
CREATE TABLE likes (
    user_id UUID NOT NULL,
    post_id UUID NOT NULL,
    PRIMARY KEY (user_id, post_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE likes;
-- +goose StatementEnd
