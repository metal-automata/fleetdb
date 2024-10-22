-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.hardware_vendors (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    CONSTRAINT hardware_vendors_name UNIQUE (name),
    CONSTRAINT hardware_vendors_pkey PRIMARY KEY (id)
);

CREATE TABLE public.hardware_models (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    hardware_vendor_id UUID NOT NULL,
    name TEXT NOT NULL,
    CONSTRAINT hardware_models_name UNIQUE (name),
    CONSTRAINT hardware_models_pk PRIMARY KEY (id),
    CONSTRAINT fk_hardware_vendor_id FOREIGN KEY (hardware_vendor_id) REFERENCES public.hardware_vendors(id) ON DELETE CASCADE
);

CREATE TABLE public.bmcs(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_id UUID NOT NULL,
    hardware_vendor_id UUID NOT NULL,
    hardware_model_id UUID NOT NULL,
    username TEXT NOT NULL,
    ipaddress INET NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT bmcs_pk PRIMARY KEY (id),
    CONSTRAINT fk_bmc_server_id FOREIGN KEY (server_id) REFERENCES public.servers(id) ON DELETE CASCADE,
    CONSTRAINT fk_bmc_hardware_vendor FOREIGN KEY (hardware_vendor_id) REFERENCES public.hardware_vendors(id),
    CONSTRAINT fk_bmc_hardware_model FOREIGN KEY (hardware_model_id) REFERENCES public.hardware_models(id)
);

ALTER TABLE public.servers ADD COLUMN vendor_id UUID REFERENCES public.hardware_vendors(id);
ALTER TABLE public.servers ADD COLUMN serial_number TEXT;
ALTER TABLE public.servers ADD COLUMN model_id UUID REFERENCES public.hardware_models(id);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE public.bmcs;
DROP TABLE public.hardware_models;
DROP TABLE public.hardawre_vendors;

ALTER TABLE public.servers DROP COLUMN vendor_id;
ALTER TABLE public.servers DROP COLUMN model_id;
ALTER TABLE public.servers DROP COLUMN serial_number;
-- +goose StatementEnd