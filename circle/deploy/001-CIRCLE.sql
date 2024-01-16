-- Deploy underground:001-CIRCLE to pg

DO $$
DECLARE
    schema text := 'circle';
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_namespace WHERE nspname = schema) THEN
        CREATE SCHEMA schema;
        RAISE NOTICE 'Schema % created.', schema;
    ELSE
        RAISE NOTICE 'Schema % already exists.', schema;
    END IF;
END $$;