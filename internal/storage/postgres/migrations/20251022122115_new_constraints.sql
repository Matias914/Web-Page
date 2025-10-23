-- Modify "celebrities" table
ALTER TABLE "celebrities" ADD CONSTRAINT "chk_name" CHECK (length(TRIM(BOTH FROM name)) > 0);
-- Modify "movies" table
ALTER TABLE "movies" ADD CONSTRAINT "chk_synopsis" CHECK (length(TRIM(BOTH FROM synopsis)) > 0), ADD CONSTRAINT "chk_title" CHECK (length(TRIM(BOTH FROM title)) > 0);
-- Create index "ak_title_release" to table: "movies"
CREATE UNIQUE INDEX "ak_title_release" ON "movies" ("title", "released_at");
-- Modify "roles" table
ALTER TABLE "roles" ADD CONSTRAINT "chk_role" CHECK (length(TRIM(BOTH FROM role)) > 0);
-- Modify "users" table
ALTER TABLE "users" ADD CONSTRAINT "chk_password" CHECK (length(TRIM(BOTH FROM password)) > 0), ADD CONSTRAINT "chk_username" CHECK (length(TRIM(BOTH FROM username)) > 0);
