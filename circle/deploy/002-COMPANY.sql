-- Deploy underground:002-COMPANY to pg

CREATE TABLE IF NOT EXISTS CIRCLE.COMPANY (
    id BIGSERIAL PRIMARY KEY,
    time_l integer,
    time_utc timestamp with time zone,
    name varchar(50),
    exchange varchar(10),
    ticker varchar(5),
    sector varchar(20),
    industry varchar(30),
    country varchar(20),
    market_cap varchar(20),
    volume integer,
    scan_type varchar(20)
);
