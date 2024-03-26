-- +goose Up
-- +goose StatementBegin
CREATE TABLE calendar (
    id BIGSERIAL PRIMARY KEY,
    description TEXT NOT NULL DEFAULT '',
    startDate TIMESTAMP WITH TIME ZONE NOT NULL,
    endDate TIMESTAMP WITH TIME ZONE NOT NULL,
    createdAt TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updatedAt TIMESTAMP WITH TIME ZONE,
    EXCLUDE USING gist (tstzrange (startDate, endDate) WITH &&)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE calendar;
-- +goose StatementEnd
