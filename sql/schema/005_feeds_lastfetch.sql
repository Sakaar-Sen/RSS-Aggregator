-- +goose Up

alter table feeds add column last_fetch timestamp;

-- +goose Down

alter table feeds drop column last_fetch;