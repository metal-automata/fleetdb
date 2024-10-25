-- +goose Up
-- +goose StatementBegin

-- This table holds records for firwmare installed on server components.
CREATE TABLE public.installed_firmware (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_component_id UUID NOT NULL,
    version TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT installed_firmware_pkey PRIMARY KEY (id),
    CONSTRAINT fk_server_component_id FOREIGN KEY (server_component_id) REFERENCES public.server_components(id) ON DELETE CASCADE
);

-- restrict duplicate records for the same component, firmware version
CREATE UNIQUE INDEX idx_installed_firmware_uniq ON public.installed_firmware(server_component_id, version);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_installed_firmware_uniq;
DROP TABLE public.installed_firmware;
-- +goose StatementEnd