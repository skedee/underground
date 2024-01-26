-- Deploy underground:007-METRIC_ENUM to pg

CREATE TYPE circle.metric_enum AS ENUM (
    'gainers',
    'top-gainers',
    'halt'
);

