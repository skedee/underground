-- Deploy underground:002-COMPANY to pg

DO $$
DECLARE
    schema text := 'circle';
    name text := 'company';
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.tables
        WHERE table_schema = schema
          AND table_name = name
    ) THEN
        -- The table does not exist, you can create it here
        CREATE TABLE CIRCLE.COMPANY (
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
        RAISE NOTICE 'Table %.% created.', schema, name;
    ELSE
         RAISE NOTICE 'Table %.% already exist.', schema, name;
    END IF;
END $$;
