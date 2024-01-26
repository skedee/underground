-- Deploy underground:003-SCAN_ENUM to pg

CREATE TYPE CIRCLE.SCAN_ENUM AS ENUM (
    'top-gainer',
    'tv-premarket',
    'halt',
    'finviz',
    'momentum-finviz'
);