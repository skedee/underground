-- Deploy underground:004-SCAN to pg

DO $$
BEGIN
    CREATE TABLE IF NOT EXISTS CIRCLE.SERIES (
        id BIGSERIAL PRIMARY KEY,
        company_id int REFERENCES CIRCLE.COMPANY(id),
        time_l integer,
        time_utc timestamp with time zone,
        resolution circle.resolution_enum
    );
END $$;
