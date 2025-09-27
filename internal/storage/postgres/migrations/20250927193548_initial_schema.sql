-- Create "genres" table
CREATE TABLE "genres" ("id" serial NOT NULL, "name" character varying(255) NOT NULL, PRIMARY KEY ("id"));
-- Create index "ak_name" to table: "genres"
CREATE UNIQUE INDEX "ak_name" ON "genres" ("name");
-- Create "movies" table
CREATE TABLE "movies" ("id" bigserial NOT NULL, "title" character varying(255) NOT NULL, "synopsis" text NOT NULL, "released_at" date NOT NULL, "poster_url" text NOT NULL, "duration_minutes" integer NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "check_duration" CHECK (duration_minutes > 0));
-- Create "categories" table
CREATE TABLE "categories" ("genre_id" integer NOT NULL, "movie_id" bigint NOT NULL, PRIMARY KEY ("movie_id", "genre_id"), CONSTRAINT "fk_categories_genres" FOREIGN KEY ("genre_id") REFERENCES "genres" ("id") ON UPDATE RESTRICT ON DELETE CASCADE, CONSTRAINT "fk_categories_movies" FOREIGN KEY ("movie_id") REFERENCES "movies" ("id") ON UPDATE RESTRICT ON DELETE CASCADE);
-- Create "users" table
CREATE TABLE "users" ("id" bigserial NOT NULL, "username" character varying(255) NOT NULL, "mail" character varying(255) NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id"), CONSTRAINT "mail_check" CHECK (POSITION(('@'::text) IN (mail)) > 1));
-- Create index "ak_mail" to table: "users"
CREATE UNIQUE INDEX "ak_mail" ON "users" ("mail");
-- Create index "ak_username" to table: "users"
CREATE UNIQUE INDEX "ak_username" ON "users" ("username");
-- Create "ratings" table
CREATE TABLE "ratings" ("user_id" bigint NOT NULL, "movie_id" bigint NOT NULL, "rating" integer NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("movie_id", "user_id"), CONSTRAINT "fk_ratings_movies" FOREIGN KEY ("movie_id") REFERENCES "movies" ("id") ON UPDATE RESTRICT ON DELETE CASCADE, CONSTRAINT "fk_ratings_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE RESTRICT ON DELETE CASCADE, CONSTRAINT "check_rating" CHECK ((rating >= 1) AND (rating <= 10)));
-- Create "reviews" table
CREATE TABLE "reviews" ("user_id" bigint NOT NULL, "movie_id" bigint NOT NULL, "comment" text NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("user_id", "movie_id"), CONSTRAINT "fk_reviews_movies" FOREIGN KEY ("movie_id") REFERENCES "movies" ("id") ON UPDATE RESTRICT ON DELETE CASCADE, CONSTRAINT "fk_reviews_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE RESTRICT ON DELETE CASCADE, CONSTRAINT "comment_check" CHECK (length(TRIM(BOTH FROM comment)) > 0));
-- Create "celebrities" table
CREATE TABLE "celebrities" ("id" bigserial NOT NULL, "name" character varying(255) NOT NULL, "birth_date" date NOT NULL, PRIMARY KEY ("id"));
-- Create "roles" table
CREATE TABLE "roles" ("movie_id" bigint NOT NULL, "celebrity_id" bigint NOT NULL, "role" character varying(255) NOT NULL, PRIMARY KEY ("movie_id", "celebrity_id"), CONSTRAINT "pk_roles_celebrities" FOREIGN KEY ("celebrity_id") REFERENCES "celebrities" ("id") ON UPDATE RESTRICT ON DELETE CASCADE, CONSTRAINT "pk_roles_movie" FOREIGN KEY ("movie_id") REFERENCES "movies" ("id") ON UPDATE RESTRICT ON DELETE CASCADE);
