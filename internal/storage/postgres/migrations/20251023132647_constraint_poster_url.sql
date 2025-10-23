-- Modify "movies" table
ALTER TABLE "movies" ADD CONSTRAINT "chk_poster" CHECK ((poster_url IS NULL) OR (length(TRIM(BOTH FROM poster_url)) > 0)), ALTER COLUMN "poster_url" DROP NOT NULL;
