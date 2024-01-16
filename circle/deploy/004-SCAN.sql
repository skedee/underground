
-- Deploy underground:004-SCAN to pg

DO $$
BEGIN
    CREATE TABLE IF NOT EXISTS CIRCLE.SCAN (
        id BIGSERIAL PRIMARY KEY,
        time_l integer,
        time_utc timestamp with time zone,
        company_id int REFERENCES CIRCLE.COMPANY(id),
        scan CIRCLE.SCAN_ENUM
    );
END $$;