-- Deploy underground:005-RESOLUTION_ENUM to pg

DO $$
DECLARE
    schema text := 'circle';
    name text := 'resolution_enum';
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_enum
        JOIN pg_type ON pg_enum.enumtypid = pg_type.oid
        JOIN pg_namespace ON pg_type.typnamespace = pg_namespace.oid
        WHERE pg_namespace.nspname = schema
        AND pg_type.typname = name
    ) THEN
        CREATE TYPE circle.resolution_enum AS ENUM (
            '1M',
            '5M',
            '15M',
            '30M',
            '60M',
            '1D'
        );
        RAISE NOTICE 'Enum %.% created.', schema, name;
    ELSE
        RAISE NOTICE 'Enum %.% already exists.', schema, name;
    END IF;
END $$;

