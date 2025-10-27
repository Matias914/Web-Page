-- Modify "ratings" table
ALTER TABLE "ratings" ALTER COLUMN "created_at" TYPE timestamp;
-- Modify "reviews" table
ALTER TABLE "reviews" ALTER COLUMN "created_at" TYPE timestamp;
-- Drop index "idx_roles_role" from table: "roles"
DROP INDEX "idx_roles_role";
-- Modify "roles" table
ALTER TABLE "roles" DROP CONSTRAINT "roles_pkey", ADD PRIMARY KEY ("celebrity_id", "movie_id");
-- Modify "users" table
ALTER TABLE "users" ALTER COLUMN "created_at" TYPE timestamp;
