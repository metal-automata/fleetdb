-- +goose Up
-- +goose StatementBegin

-- This table holds component adds/deletes *for servers with an existing component inventory*
-- each record then can be reviewed by an operator/automation and applied to the server_components table,
-- after review the change is then merged into the server_components and relation tables and the row is deleted.
CREATE TABLE public.component_change_reports (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    report_id UUID NOT NULL,  -- report id is set by the API handler and is used to associate multiple change reports
    server_id UUID NOT NULL,
    server_component_id UUID, -- This column does not reference public.server_components.id since it will be uuid.Nil for new component additions
    server_component_name TEXT NOT NULL, -- component type slug for readability
    server_component_type_id UUID NOT NULL, -- component type ID to keep referential integrity
    remove_component boolean DEFAULT false,
    serial TEXT NOT NULL,
    collection_method TEXT NOT NULL,
    data JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT component_change_reports_pkey PRIMARY KEY (id),
    CONSTRAINT fk_component_change_reports_server_id FOREIGN KEY (server_id) REFERENCES public.servers(id) ON DELETE CASCADE,
    CONSTRAINT fk_component_change_reports_component_type_id FOREIGN KEY (server_component_type_id) REFERENCES public.server_component_types(id) ON DELETE CASCADE,
    CONSTRAINT check_data_not_empty CHECK (data != '{}'::jsonb)
);


-- uniq composite index on server_id, serial, component_type_id
CREATE UNIQUE INDEX idx_component_change_reports_uniq ON public.component_change_reports (server_id, serial, server_component_name, remove_component);
-- restrict duplicate component_capabilities records
CREATE UNIQUE INDEX idx_component_capabilities_uniq ON public.component_capabilities(server_component_id, name);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE public.component_change_reports;
-- +goose StatementEnd