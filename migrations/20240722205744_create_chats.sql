-- +goose Up
-- +goose StatementBegin
CREATE TABLE Chats
(
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(255) NOT NULL,
    user_ids INT[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Chats;
-- +goose StatementEnd
