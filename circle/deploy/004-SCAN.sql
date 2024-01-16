
-- Deploy underground:004-SCAN to pg

DO $$
BEGIN
    CREATE TABLE IF NOT EXISTS CIRCLE.SCAN (
        id BIGSERIAL PRIMARY KEY,
        company_id int REFERENCES CIRCLE.COMPANY(id),
        time_l integer,
        time_utc timestamp with time zone,
        scan CIRCLE.SCAN_ENUM
    );
END $$;