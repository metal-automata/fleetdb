-- +goose Up
-- +goose StatementBegin

-- Add inventory refreshed TS column
ALTER TABLE public.servers ADD COLUMN inventory_refreshed_at TIMESTAMP WITH TIME ZONE NULL;

-- server status table
CREATE TABLE public.server_status (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_id UUID NOT NULL,
    health TEXT NOT NULL,
    state TEXT NOT NULL,
    info TEXT,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT server_status_pkey PRIMARY KEY (id),
    CONSTRAINT fk_server_status_server_id FOREIGN KEY (server_id) REFERENCES public.servers(id) ON DELETE CASCADE
);

-- restrict duplicate status records for server
-- If we do require to store multiple records - consider an ETL, audit triggers or another time series database
CREATE UNIQUE INDEX idx_server_status_uniq ON public.server_status(server_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX public.idx_server_status_uniq;
DROP TABLE public.server_status;

ALTER TABLE public.servers DROP COLUMN inventory_refreshed_at;
-- +goose StatementEnd