-- +goose Up
-- +goose StatementBegin

CREATE TABLE public.servers (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NULL,
    facility_code STRING NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    INDEX idx_facility (facility_code ASC)
);

CREATE TABLE public.server_component_types (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    slug STRING NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    UNIQUE INDEX idx_name (name ASC),
    UNIQUE INDEX server_component_types_slug_key (slug ASC)
);

CREATE TABLE public.server_components (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NULL,
    vendor STRING NULL,
    model STRING NULL,
    serial STRING NULL,
    server_component_type_id UUID NOT NULL,
    server_id UUID NOT NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    INDEX idx_server_component_type_id (server_component_type_id ASC),
    INDEX idx_server_id (server_id ASC),
    UNIQUE INDEX idx_server_components (server_id ASC, serial ASC, server_component_type_id ASC) WHERE ((server_id IS NOT NULL) AND (serial IS NOT NULL)) AND (server_component_type_id IS NOT NULL),
    CONSTRAINT fk_server_component_type_id_ref_server_component_types FOREIGN KEY (server_component_type_id) REFERENCES public.server_component_types(id),
    CONSTRAINT fk_server_id_ref_servers FOREIGN KEY (server_id) REFERENCES public.servers(id) ON DELETE CASCADE
);

CREATE TABLE public.versioned_attributes (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_id UUID NULL,
    namespace STRING NOT NULL,
    data JSONB NOT NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    server_component_id UUID NULL,
    tally INT8 NOT NULL DEFAULT 0:::INT8,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    INDEX idx_server_id (server_id ASC) WHERE server_id IS NOT NULL,
    INDEX idx_server_namespace (server_id ASC, namespace ASC, created_at ASC) WHERE server_id IS NOT NULL,
    INVERTED INDEX idx_server_data (server_id ASC, namespace ASC, data) WHERE server_id IS NOT NULL,
    INDEX idx_server_component_id (server_component_id ASC) WHERE server_component_id IS NOT NULL,
    INDEX idx_server_component_namespace (server_component_id ASC, namespace ASC, created_at ASC) WHERE server_component_id IS NOT NULL,
    INVERTED INDEX idx_server_component_data (server_component_id ASC, namespace ASC, data) WHERE server_component_id IS NOT NULL,
    CONSTRAINT check_server_id_server_component_id CHECK (((server_id IS NOT NULL)::INT8 + (server_component_id IS NOT NULL)::INT8) = 1:::INT8),
    CONSTRAINT fk_server_id_ref_servers FOREIGN KEY (server_id) REFERENCES public.servers(id) ON DELETE CASCADE,
    CONSTRAINT fk_server_component_id_ref_server_components FOREIGN KEY (server_component_id) REFERENCES public.server_components(id) ON DELETE CASCADE
);

CREATE TABLE public.attributes (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_id UUID NULL,
    server_component_id UUID NULL,
    namespace STRING NOT NULL,
    data JSONB NOT NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    INDEX idx_server_id (server_id ASC) WHERE server_id IS NOT NULL,
    UNIQUE INDEX idx_server_namespace (server_id ASC, namespace ASC) WHERE server_id IS NOT NULL,
    INVERTED INDEX idx_server_data (server_id ASC, namespace ASC, data) WHERE server_id IS NOT NULL,
    INDEX idx_server_component_id (server_component_id ASC) WHERE server_component_id IS NOT NULL,
    UNIQUE INDEX idx_server_component_namespace (server_component_id ASC, namespace ASC) WHERE server_component_id IS NOT NULL,
    INVERTED INDEX idx_server_component_data (server_component_id ASC, namespace ASC, data) WHERE server_component_id IS NOT NULL,
    CONSTRAINT check_server_id_server_component_id CHECK (((server_id IS NOT NULL)::INT8 + (server_component_id IS NOT NULL)::INT8) = 1:::INT8),
    CONSTRAINT fk_server_id_ref_servers FOREIGN KEY (server_id) REFERENCES public.servers(id) ON DELETE CASCADE,
    CONSTRAINT fk_server_component_id_ref_server_components FOREIGN KEY (server_component_id) REFERENCES public.server_components(id) ON DELETE CASCADE
);

CREATE TABLE public.component_firmware_version (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    component STRING NOT NULL,
    vendor STRING NOT NULL,
    model STRING[] NOT NULL,
    filename STRING NOT NULL,
    version STRING NOT NULL,
    checksum STRING NOT NULL,
    upstream_url STRING NOT NULL,
    repository_url STRING NOT NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    install_inband BOOL NOT NULL DEFAULT false,
    oem BOOL NOT NULL DEFAULT false,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    UNIQUE INDEX vendor_component_version_filename_unique (vendor ASC, component ASC, version ASC, filename ASC)
);

CREATE TABLE public.server_credential_types (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL,
    slug STRING NOT NULL,
    builtin BOOL NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    UNIQUE INDEX server_secret_types_slug_key (slug ASC)
);

CREATE TABLE public.server_credentials (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    server_id UUID NOT NULL,
    server_credential_type_id UUID NOT NULL,
    password STRING NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    username STRING NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    UNIQUE INDEX idx_server_secrets_by_type (server_id ASC, server_credential_type_id ASC),
    CONSTRAINT fk_server_id_ref_servers FOREIGN KEY (server_id) REFERENCES public.servers(id) ON DELETE CASCADE,
    CONSTRAINT fk_server_secret_type_id_ref_server_secret_types FOREIGN KEY (server_credential_type_id) REFERENCES public.server_credential_types(id)
);

CREATE TABLE public.component_firmware_set (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    UNIQUE INDEX idx_name (name ASC)
);

CREATE TABLE public.component_firmware_set_map (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    firmware_set_id UUID NOT NULL,
    firmware_id UUID NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    UNIQUE INDEX component_firmware_set_map_firmware_set_id_firmware_id_key (firmware_set_id ASC, firmware_id ASC),
    CONSTRAINT fk_firmware_set_id_ref_component_firmware_set FOREIGN KEY (firmware_set_id) REFERENCES public.component_firmware_set(id) ON DELETE CASCADE,
    CONSTRAINT fk_firmware_id_ref_component_firmware_version FOREIGN KEY (firmware_id) REFERENCES public.component_firmware_version(id) ON DELETE RESTRICT
);

CREATE TABLE public.attributes_firmware_set (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    firmware_set_id UUID NULL,
    namespace STRING NOT NULL,
    data JSONB NOT NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    INDEX idx_firmware_set_id (firmware_set_id ASC) WHERE firmware_set_id IS NOT NULL,
    INVERTED INDEX idx_firmware_set_data (firmware_set_id ASC, namespace ASC, data) WHERE firmware_set_id IS NOT NULL,
    UNIQUE INDEX idx_firmware_set_namespace (firmware_set_id ASC, namespace ASC) WHERE firmware_set_id IS NOT NULL,
    CONSTRAINT fk_firmware_set_id_ref_component_firmware_set FOREIGN KEY (firmware_set_id) REFERENCES public.component_firmware_set(id) ON DELETE CASCADE
);

CREATE TABLE public.bom_info (
    serial_num STRING NOT NULL,
    aoc_mac_address STRING NULL,
    bmc_mac_address STRING NULL,
    num_defi_pmi STRING NULL,
    num_def_pwd STRING NULL,
    metro STRING NULL,
    CONSTRAINT "primary" PRIMARY KEY (serial_num ASC)
);

CREATE TABLE public.aoc_mac_address (
    aoc_mac_address STRING NOT NULL,
    serial_num STRING NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (aoc_mac_address ASC),
    CONSTRAINT fk_serial_num_ref_bom_info FOREIGN KEY (serial_num) REFERENCES public.bom_info(serial_num) ON DELETE CASCADE
);

CREATE TABLE public.bmc_mac_address (
    bmc_mac_address STRING NOT NULL,
    serial_num STRING NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (bmc_mac_address ASC),
    CONSTRAINT fk_serial_num_ref_bom_info FOREIGN KEY (serial_num) REFERENCES public.bom_info(serial_num) ON DELETE CASCADE
);

CREATE TABLE public.bios_config_sets (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL,
    version STRING NOT NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    CONSTRAINT bios_config_sets_pkey PRIMARY KEY (id ASC),
    UNIQUE INDEX bios_config_sets_name_key (name ASC)
);

CREATE TABLE public.bios_config_components (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    fk_bios_config_set_id UUID NOT NULL,
    name STRING NOT NULL,
    vendor STRING NOT NULL,
    model STRING NOT NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    CONSTRAINT bios_config_components_pkey PRIMARY KEY (id ASC),
    UNIQUE INDEX bios_config_components_fk_bios_config_set_id_name_key (fk_bios_config_set_id ASC, name ASC),
    CONSTRAINT bios_config_components_fk_bios_config_set_id_fkey FOREIGN KEY (fk_bios_config_set_id) REFERENCES public.bios_config_sets(id) ON DELETE CASCADE
);

CREATE TABLE public.bios_config_settings (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    fk_bios_config_component_id UUID NOT NULL,
    settings_key STRING NOT NULL,
    settings_value STRING NOT NULL,
    raw JSONB NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    CONSTRAINT bios_config_settings_pkey PRIMARY KEY (id ASC),
    UNIQUE INDEX bios_config_settings_fk_bios_config_component_id_settings_key_key (fk_bios_config_component_id ASC, settings_key ASC),
    CONSTRAINT bios_config_settings_fk_bios_config_component_id_fkey FOREIGN KEY (fk_bios_config_component_id) REFERENCES public.bios_config_components(id) ON DELETE CASCADE
);

CREATE TABLE public.event_history (
    event_id UUID NOT NULL,
    event_type STRING NOT NULL,
    event_start TIMESTAMPTZ NOT NULL,
    event_end TIMESTAMPTZ NOT NULL,
    target_server UUID NOT NULL,
    parameters JSONB NULL,
    final_state STRING NOT NULL,
    final_status JSONB NULL,
    CONSTRAINT event_history_pkey PRIMARY KEY (event_id ASC, event_type ASC, target_server ASC),
    CONSTRAINT event_history_target_server_fkey FOREIGN KEY (target_server) REFERENCES public.servers(id) ON DELETE CASCADE
);

CREATE TABLE public.firmware_set_validation_facts (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    firmware_set_id UUID NOT NULL,
    target_server_id UUID NOT NULL,
    performed_on TIMESTAMPTZ NOT NULL,
    CONSTRAINT firmware_set_validation_facts_pkey PRIMARY KEY (id ASC),
    CONSTRAINT firmware_set_validation_facts_firmware_set_id_fkey FOREIGN KEY (firmware_set_id) REFERENCES public.component_firmware_set(id) ON DELETE CASCADE
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