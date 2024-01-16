-- Deploy underground:001-CIRCLE to pg

DO $$
BEGIN
    CREATE SCHEMA IF NOT EXISTS circle;
END $$;