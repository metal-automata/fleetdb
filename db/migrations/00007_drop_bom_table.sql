-- +goose Up
-- +goose StatementBegin
DROP TABLE public.bom_info CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
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
-- +goose StatementEnd