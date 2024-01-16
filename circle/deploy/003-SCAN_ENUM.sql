-- Deploy underground:003-SCAN_ENUM to pg

DO $$
DECLARE
    schema text := 'circle';
    name text := 'scan_enum';
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_enum
        JOIN pg_type ON pg_enum.enumtypid = pg_type.oid
        JOIN pg_namespace ON pg_type.typnamespace = pg_namespace.oid
        WHERE pg_namespace.nspname = schema
        AND pg_type.typname = name
    ) THEN
        CREATE TYPE CIRCLE.SCAN_ENUM AS ENUM (
            'top-gainer',
            'tv-premarket',
            'halt',
            'finviz',
            'momentum-finviz'
        );
        RAISE NOTICE 'Enum %.% created.', schema, name;
    ELSE
        RAISE NOTICE 'Enum %.% already exists.', schema, name;
    END IF;
END $$;
