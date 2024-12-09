-- +goose Up
-- +goose StatementBegin

-- This table holds records for firwmare installed on server components.
CREATE TABLE public.component_status (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_component_id UUID NOT NULL,
    health TEXT NOT NULL,
    state TEXT NOT NULL,
    info TEXT,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT component_status_pkey PRIMARY KEY (id),
    CONSTRAINT fk_component_status_server_component_id FOREIGN KEY (server_component_id) REFERENCES public.server_components(id) ON DELETE CASCADE
);

-- If we do require to store multiple records - consider an ETL or another time series database
-- restrict duplicate status records for component
CREATE UNIQUE INDEX idx_component_status_uniq ON public.component_status(server_component_id);

-- component metadata table, data here should be limited to a one level depth JSON
CREATE TABLE public.component_metadata (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_component_id UUID NOT NULL,
    namespace TEXT NOT NULL,
    data JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT component_metadata_pkey PRIMARY KEY (id),
    CONSTRAINT fk_component_metadata_server_component_id FOREIGN KEY (server_component_id) REFERENCES public.server_components(id) ON DELETE CASCADE,
    CONSTRAINT check_data_not_empty CHECK (data != '{}'::jsonb)
);

-- restrict duplicate metadata records for component
CREATE UNIQUE INDEX id_component_metadata_uniq ON public.component_metadata(server_component_id, namespace);

-- This table holds component capabilities
CREATE TABLE public.component_capabilities (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_component_id UUID NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    enabled bool default false,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT component_capabilities_pkey PRIMARY KEY (id),
    CONSTRAINT fk_component_capabilities_server_component_id FOREIGN KEY (server_component_id) REFERENCES public.server_components(id) ON DELETE CASCADE
);


-- Add component description, oem columns
ALTER TABLE public.server_components ADD COLUMN description TEXT;
ALTER TABLE public.server_components ADD COLUMN oem bool;


-- server components unique index updated - since now serial cannot be NULL 
ALTER TABLE public.server_components ALTER COLUMN serial SET NOT NULL;
DROP INDEX public.idx_server_components;
CREATE UNIQUE INDEX idx_server_components ON public.server_components (server_id, serial, server_component_type_id);


-- add index to enable upserts - since the upsert conflicts clause is on the server_component_id
CREATE UNIQUE INDEX idx_installed_firmware_uniq_for_conflict_clause ON public.installed_firmware(server_component_id);

-- removes deleted_at since we hold only a single record per server component firmware
ALTER TABLE public.installed_firmware DROP COLUMN deleted_at;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP INDEX public.idx_component_status_uniq;
DROP TABLE public.component_status;

DROP INDEX public.idx_component_metadata_uniq;
DROP TABLE public.component_metadata;

DROP INDEX public.idx_component_capabilities_uniq;
DROP TABLE public.component_capabilities;

ALTER TABLE public.server_components DROP COLUMN description;
ALTER TABLE public.server_components DROP COLUMN oem;

DROP UNIQUE INDEX idx_installed_firmware_uniq_for_conflict_clause;
ALTER TABLE public.installed_firmware ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE NULL;

ALTER TABLE public.server_components ALTER COLUMN serial TEXT DROP NOT NULL;
CREATE UNIQUE INDEX idx_server_components ON public.server_components (server_id, serial, server_component_type_id) 
    WHERE server_id IS NOT NULL AND serial IS NOT NULL AND server_component_type_id IS NOT NULL;
-- +goose StatementEnd
