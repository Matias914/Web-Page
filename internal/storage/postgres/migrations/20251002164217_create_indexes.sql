-- Modify "categories" table
ALTER TABLE "categories" DROP CONSTRAINT "categories_pkey", ADD PRIMARY KEY ("genre_id", "movie_id");
-- Create index "idx_categories_movies_id" to table: "categories"
CREATE INDEX "idx_categories_movies_id" ON "categories" ("movie_id");
-- Modify "ratings" table
ALTER TABLE "ratings" DROP CONSTRAINT "ratings_pkey", ADD PRIMARY KEY ("user_id", "movie_id");
-- Create index "idx_ratings_movies_id" to table: "ratings"
CREATE INDEX "idx_ratings_movies_id" ON "ratings" ("movie_id");
-- Create index "idx_reviews_movies_id" to table: "reviews"
CREATE INDEX "idx_reviews_movies_id" ON "reviews" ("movie_id");
-- Modify "roles" table
ALTER TABLE "roles" DROP CONSTRAINT "roles_pkey", DROP CONSTRAINT "pk_roles_celebrities", DROP CONSTRAINT "pk_roles_movie", ADD PRIMARY KEY ("celebrity_id", "movie_id", "role"), ADD CONSTRAINT "fk_roles_celebrities" FOREIGN KEY ("celebrity_id") REFERENCES "celebrities" ("id") ON UPDATE RESTRICT ON DELETE CASCADE, ADD CONSTRAINT "fk_roles_movie" FOREIGN KEY ("movie_id") REFERENCES "movies" ("id") ON UPDATE RESTRICT ON DELETE CASCADE;
-- Create index "idx_roles_movies_id" to table: "roles"
CREATE INDEX "idx_roles_movies_id" ON "roles" ("movie_id");
-- Create index "idx_roles_role" to table: "roles"
CREATE INDEX "idx_roles_role" ON "roles" ("role");
