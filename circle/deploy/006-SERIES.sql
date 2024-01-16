-- Deploy underground:004-SCAN to pg

DO $$
DECLARE
    schema text := 'circle';
    name text := 'series';
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.tables
        WHERE table_schema = schema
          AND table_name = name
    ) THEN
        -- The table does not exist, you can create it here
        CREATE TABLE CIRCLE.SERIES (
            id BIGSERIAL PRIMARY KEY,
            company_id int REFERENCES CIRCLE.COMPANY(id),
            time_l integer,
            time_utc timestamp with time zone,
            resolution circle.resolution_enum
        );
        RAISE NOTICE 'Table %.% created.', schema, name;
    ELSE
         RAISE NOTICE 'Table %.% already exist.', schema, name;
    END IF;
END $$;
