-- Deploy underground:005-RESOLUTION_ENUM to pg

CREATE TYPE circle.resolution_enum AS ENUM (
    '1M',
    '5M',
    '15M',
    '30M',
    '60M',
    '1D'
);