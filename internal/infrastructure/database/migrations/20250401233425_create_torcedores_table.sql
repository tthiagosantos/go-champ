-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS torcedores (
                                          id SERIAL PRIMARY KEY,
                                          nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    time VARCHAR(255) NOT NULL
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS torcedores;
-- +goose StatementEnd
