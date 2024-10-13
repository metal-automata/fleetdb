-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.servers (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NULL,
    facility_code TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT servers_pkey PRIMARY KEY (id)
);

CREATE INDEX idx_facility ON public.servers (facility_code);

CREATE TABLE public.server_component_types (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    slug TEXT NOT NULL,
    CONSTRAINT server_component_types_pkey PRIMARY KEY (id),
    CONSTRAINT server_component_types_name_key UNIQUE (name),
    CONSTRAINT server_component_types_slug_key UNIQUE (slug)
);

CREATE TABLE public.server_components (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NULL,
    vendor TEXT NULL,
    model TEXT NULL,
    serial TEXT NULL,
    server_component_type_id UUID NOT NULL,
    server_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT server_components_pkey PRIMARY KEY (id),
    CONSTRAINT fk_server_component_type_id_ref_server_component_types FOREIGN KEY (server_component_type_id) REFERENCES public.server_component_types(id),
    CONSTRAINT fk_server_id_ref_servers FOREIGN KEY (server_id) REFERENCES public.servers(id) ON DELETE CASCADE
);

CREATE INDEX idx_server_component_type_id ON public.server_components (server_component_type_id);
CREATE INDEX idx_server_components_server_id ON public.server_components (server_id);
CREATE UNIQUE INDEX idx_server_components ON public.server_components (server_id, serial, server_component_type_id) 
    WHERE server_id IS NOT NULL AND serial IS NOT NULL AND server_component_type_id IS NOT NULL;

CREATE TABLE public.versioned_attributes (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_id UUID NULL,
    namespace TEXT NOT NULL,
    data JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    server_component_id UUID NULL,
    tally BIGINT NOT NULL DEFAULT 0,
    CONSTRAINT versioned_attributes_pkey PRIMARY KEY (id),
    CONSTRAINT check_server_id_server_component_id CHECK (
        (CASE WHEN server_id IS NOT NULL THEN 1 ELSE 0 END +
         CASE WHEN server_component_id IS NOT NULL THEN 1 ELSE 0 END) = 1
    ),
    CONSTRAINT fk_server_id_ref_servers FOREIGN KEY (server_id) REFERENCES public.servers(id) ON DELETE CASCADE,
    CONSTRAINT fk_server_component_id_ref_server_components FOREIGN KEY (server_component_id) REFERENCES public.server_components(id) ON DELETE CASCADE
);

CREATE INDEX idx_versioned_attributes_server_id ON public.versioned_attributes (server_id) WHERE server_id IS NOT NULL;
CREATE INDEX idx_versioned_attributes_server_namespace ON public.versioned_attributes (server_id, namespace, created_at) WHERE server_id IS NOT NULL;
-- composite index with server_id, namespace using B-tree index type
-- and with data using the GIN on the jsonb_ops operator class
CREATE INDEX idx_versioned_attributes_server_data ON public.versioned_attributes (server_id, namespace, data jsonb_ops) WHERE server_id IS NOT NULL;
CREATE INDEX idx_versioned_attributes_server_component_id ON public.versioned_attributes (server_component_id) WHERE server_component_id IS NOT NULL;
CREATE INDEX idx_versioned_attributes_server_component_namespace ON public.versioned_attributes (server_component_id, namespace, created_at) WHERE server_component_id IS NOT NULL;
-- composite index with server_component_id, namespace using B-tree index type
-- and with data using the GIN on the jsonb_ops operator class
CREATE INDEX idx_versioned_attributes_server_component_data ON public.versioned_attributes (server_component_id, namespace, data jsonb_ops) WHERE server_component_id IS NOT NULL;

CREATE TABLE public.attributes (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_id UUID NULL,
    server_component_id UUID NULL,
    namespace TEXT NOT NULL,
    data JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT attributes_pkey PRIMARY KEY (id),
    CONSTRAINT check_server_id_server_component_id CHECK (
        (CASE WHEN server_id IS NOT NULL THEN 1 ELSE 0 END +
         CASE WHEN server_component_id IS NOT NULL THEN 1 ELSE 0 END) = 1
    ),
    CONSTRAINT fk_server_id_ref_servers FOREIGN KEY (server_id) REFERENCES public.servers(id) ON DELETE CASCADE,
    CONSTRAINT fk_server_component_id_ref_server_components FOREIGN KEY (server_component_id) REFERENCES public.server_components(id) ON DELETE CASCADE
);

CREATE INDEX idx_attributes_server_id ON public.attributes (server_id) WHERE server_id IS NOT NULL;
CREATE UNIQUE INDEX idx_attributes_server_namespace ON public.attributes (server_id, namespace) WHERE server_id IS NOT NULL;
-- composite index with server_id, namespace using B-tree index type
-- and with data using the GIN on the jsonb_ops operator class
CREATE INDEX idx_attributes_server_data ON public.attributes (server_id, namespace, data jsonb_ops) WHERE server_id IS NOT NULL;
CREATE INDEX idx_attributes_server_component_id ON public.attributes (server_component_id) WHERE server_component_id IS NOT NULL;
-- composite index with server_component_id, namespace using B-tree index type
-- and with data using the GIN on the jsonb_ops operator class
CREATE UNIQUE INDEX idx_attributes_server_component_namespace ON public.attributes (server_component_id, namespace) WHERE server_component_id IS NOT NULL;
CREATE INDEX idx_attributes_server_component_data ON public.attributes (server_component_id, namespace, data jsonb_ops) WHERE server_component_id IS NOT NULL;

CREATE TABLE public.component_firmware_version (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    component TEXT NOT NULL,
    vendor TEXT NOT NULL,
    model TEXT[] NOT NULL,
    filename TEXT NOT NULL,
    version TEXT NOT NULL,
    checksum TEXT NOT NULL,
    upstream_url TEXT NOT NULL,
    repository_url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    install_inband BOOLEAN NOT NULL DEFAULT false,
    oem BOOLEAN NOT NULL DEFAULT false,
    CONSTRAINT component_firmware_version_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX vendor_component_version_filename_unique ON public.component_firmware_version (vendor, component, version, filename);

CREATE TABLE public.server_credential_types (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    slug TEXT NOT NULL,
    builtin BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT server_credential_types_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX server_secret_types_slug_key ON public.server_credential_types (slug);

CREATE TABLE public.server_credentials (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_id UUID NOT NULL,
    server_credential_type_id UUID NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    username TEXT NOT NULL,
    CONSTRAINT server_credentials_pkey PRIMARY KEY (id),
    CONSTRAINT fk_server_id_ref_servers FOREIGN KEY (server_id) REFERENCES public.servers(id) ON DELETE CASCADE,
    CONSTRAINT fk_server_secret_type_id_ref_server_secret_types FOREIGN KEY (server_credential_type_id) REFERENCES public.server_credential_types(id)
);

CREATE UNIQUE INDEX idx_server_secrets_by_type ON public.server_credentials (server_id, server_credential_type_id);

CREATE TABLE public.component_firmware_set (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT component_firmware_set_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX idx_name ON public.component_firmware_set (name);

CREATE TABLE public.component_firmware_set_map (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    firmware_set_id UUID NOT NULL,
    firmware_id UUID NOT NULL,
    CONSTRAINT component_firmware_set_map_pkey PRIMARY KEY (id),
    CONSTRAINT fk_firmware_set_id_ref_component_firmware_set FOREIGN KEY (firmware_set_id) REFERENCES public.component_firmware_set(id) ON DELETE CASCADE,
    CONSTRAINT fk_firmware_id_ref_component_firmware_version FOREIGN KEY (firmware_id) REFERENCES public.component_firmware_version(id) ON DELETE RESTRICT
);

CREATE UNIQUE INDEX component_firmware_set_map_firmware_set_id_firmware_id_key ON public.component_firmware_set_map (firmware_set_id, firmware_id);

CREATE TABLE public.attributes_firmware_set (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    firmware_set_id UUID NULL,
    namespace TEXT NOT NULL,
    data JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT attributes_firmware_set_pkey PRIMARY KEY (id),
    CONSTRAINT fk_firmware_set_id_ref_component_firmware_set FOREIGN KEY (firmware_set_id) REFERENCES public.component_firmware_set(id) ON DELETE CASCADE
);

CREATE INDEX idx_firmware_set_id ON public.attributes_firmware_set (firmware_set_id) WHERE firmware_set_id IS NOT NULL;
-- composite index with firmware_set_id, namespace using B-tree index type
-- and with data using the GIN on the jsonb_ops operator class
CREATE INDEX idx_firmware_set_data ON public.attributes_firmware_set (firmware_set_id, namespace, data jsonb_ops) WHERE firmware_set_id IS NOT NULL;
CREATE UNIQUE INDEX idx_firmware_set_namespace ON public.attributes_firmware_set (firmware_set_id, namespace) WHERE firmware_set_id IS NOT NULL;

CREATE TABLE public.bom_info (
    serial_num TEXT NOT NULL,
    aoc_mac_address TEXT,
    bmc_mac_address TEXT,
    num_defi_pmi TEXT,
    num_def_pwd TEXT,
    metro TEXT,
    PRIMARY KEY (serial_num)
);

CREATE TABLE public.aoc_mac_address (
    aoc_mac_address TEXT NOT NULL,
    serial_num TEXT NOT NULL,
    PRIMARY KEY (aoc_mac_address),
    FOREIGN KEY (serial_num) REFERENCES public.bom_info(serial_num) ON DELETE CASCADE
);

CREATE TABLE public.bmc_mac_address (
    bmc_mac_address TEXT NOT NULL,
    serial_num TEXT NOT NULL,
    PRIMARY KEY (bmc_mac_address),
    FOREIGN KEY (serial_num) REFERENCES public.bom_info(serial_num) ON DELETE CASCADE
);

CREATE TABLE public.bios_config_sets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    version TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    UNIQUE (name)
);

CREATE TABLE public.bios_config_components (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fk_bios_config_set_id UUID NOT NULL,
    name TEXT NOT NULL,
    vendor TEXT NOT NULL,
    model TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    UNIQUE (fk_bios_config_set_id, name),
    FOREIGN KEY (fk_bios_config_set_id) REFERENCES public.bios_config_sets(id) ON DELETE CASCADE
);

CREATE TABLE public.bios_config_settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fk_bios_config_component_id UUID NOT NULL,
    settings_key TEXT NOT NULL,
    settings_value TEXT NOT NULL,
    raw JSONB,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    UNIQUE (fk_bios_config_component_id, settings_key),
    FOREIGN KEY (fk_bios_config_component_id) REFERENCES public.bios_config_components(id) ON DELETE CASCADE
);

CREATE TABLE public.event_history (
    event_id UUID NOT NULL,
    event_type TEXT NOT NULL,
    event_start TIMESTAMP WITH TIME ZONE NOT NULL,
    event_end TIMESTAMP WITH TIME ZONE NOT NULL,
    target_server UUID NOT NULL,
    parameters JSONB,
    final_state TEXT NOT NULL,
    final_status JSONB,
    PRIMARY KEY (event_id, event_type, target_server),
    FOREIGN KEY (target_server) REFERENCES public.servers(id) ON DELETE CASCADE
);

CREATE TABLE public.firmware_set_validation_facts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    firmware_set_id UUID NOT NULL,
    target_server_id UUID NOT NULL,
    performed_on TIMESTAMP WITH TIME ZONE NOT NULL,
    FOREIGN KEY (firmware_set_id) REFERENCES public.component_firmware_set(id) ON DELETE CASCADE
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE public.firmware_set_validation_facts;
DROP TABLE public.event_history;
DROP TABLE public.bios_config_settings;
DROP TABLE public.bios_config_components;
DROP TABLE public.bios_config_sets;
DROP TABLE public.bmc_mac_address;
DROP TABLE public.aoc_mac_address;
DROP TABLE public.bom_info;
DROP TABLE public.attributes_firmware_set;
DROP TABLE public.component_firmware_set_map;
DROP TABLE public.component_firmware_set;
DROP TABLE public.server_credentials;
DROP TABLE public.server_credential_types;
DROP TABLE public.component_firmware_version;
DROP TABLE public.attributes;
DROP TABLE public.versioned_attributes;
DROP TABLE public.server_components;
DROP TABLE public.server_component_types;
DROP TABLE public.servers;
-- +goose StatementEnd