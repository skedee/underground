-- Deploy underground:008-METRIC to pg

DO $$
BEGIN
    CREATE TABLE IF NOT EXISTS CIRCLE.METRIC (
        id BIGSERIAL PRIMARY KEY,
        time_l integer,
        time_utc timestamp with time zone,
        company_id int REFERENCES CIRCLE.COMPANY(id),
        metric circle.metric_enum
    );
END $$;
