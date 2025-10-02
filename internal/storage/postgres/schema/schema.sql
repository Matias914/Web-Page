-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2025-10-02 16:40:08.698

-- tables
-- Table: categories
CREATE TABLE categories (
    genre_id int  NOT NULL,
    movie_id bigint  NOT NULL,
    CONSTRAINT PK_CATEGORIES PRIMARY KEY (genre_id,movie_id)
);

CREATE INDEX idx_categories_movies_id on categories (movie_id ASC);

-- Table: celebrities
CREATE TABLE celebrities (
    id bigserial  NOT NULL,
    name varchar(255)  NOT NULL,
    birth_date date  NOT NULL,
    CONSTRAINT PK_CELEBRITIES PRIMARY KEY (id)
);

-- Table: genres
CREATE TABLE genres (
    id serial  NOT NULL,
    name varchar(255)  NOT NULL,
    CONSTRAINT AK_NAME UNIQUE (name) NOT DEFERRABLE  INITIALLY IMMEDIATE,
    CONSTRAINT PK_GENRES PRIMARY KEY (id)
);

-- Table: movies
CREATE TABLE movies (
    id bigserial  NOT NULL,
    title varchar(255)  NOT NULL,
    synopsis text  NOT NULL,
    released_at date  NOT NULL,
    poster_url text  NOT NULL,
    duration_minutes int  NOT NULL,
    CONSTRAINT check_duration CHECK (duration_minutes > 0) NOT DEFERRABLE INITIALLY IMMEDIATE,
    CONSTRAINT PK_MOVIES PRIMARY KEY (id)
);

-- Table: ratings
CREATE TABLE ratings (
    user_id bigint  NOT NULL,
    movie_id bigint  NOT NULL,
    rating int  NOT NULL,
    created_at timestamp with time zone  NOT NULL DEFAULT NOW(),
    CONSTRAINT check_rating CHECK (rating between 1 and 10) NOT DEFERRABLE INITIALLY IMMEDIATE,
    CONSTRAINT PK_RATINGS PRIMARY KEY (user_id,movie_id)
);

CREATE INDEX idx_ratings_movies_id on ratings (movie_id ASC);

-- Table: reviews
CREATE TABLE reviews (
    user_id bigint  NOT NULL,
    movie_id bigint  NOT NULL,
    comment text  NOT NULL,
    created_at timestamp with time zone  NOT NULL DEFAULT NOW(),
    CONSTRAINT comment_check CHECK (length(trim(comment)) > 0) NOT DEFERRABLE INITIALLY IMMEDIATE,
    CONSTRAINT PK_REVIEWS PRIMARY KEY (user_id,movie_id)
);

CREATE INDEX idx_reviews_movies_id on reviews (movie_id ASC);

-- Table: roles
CREATE TABLE roles (
    celebrity_id bigint  NOT NULL,
    movie_id bigint  NOT NULL,
    role varchar(255)  NOT NULL,
    CONSTRAINT PK_ROLES PRIMARY KEY (celebrity_id,movie_id,role)
);

CREATE INDEX idx_roles_movies_id on roles (movie_id ASC);

CREATE INDEX idx_roles_role on roles (role ASC);

-- Table: users
CREATE TABLE users (
    id bigserial  NOT NULL,
    username varchar(255)  NOT NULL,
    mail varchar(255)  NOT NULL,
    created_at timestamp with time zone  NOT NULL DEFAULT NOW(),
    CONSTRAINT AK_USERNAME UNIQUE (username) NOT DEFERRABLE  INITIALLY IMMEDIATE,
    CONSTRAINT AK_MAIL UNIQUE (mail) NOT DEFERRABLE  INITIALLY IMMEDIATE,
    CONSTRAINT mail_check CHECK (position('@' in mail) > 1) NOT DEFERRABLE INITIALLY IMMEDIATE,
    CONSTRAINT PK_USERS PRIMARY KEY (id)
);

-- foreign keys
-- Reference: FK_CATEGORIES_GENRES (table: categories)
ALTER TABLE categories ADD CONSTRAINT FK_CATEGORIES_GENRES
    FOREIGN KEY (genre_id)
    REFERENCES genres (id)
    ON DELETE  CASCADE
    ON UPDATE  RESTRICT
    NOT DEFERRABLE
    INITIALLY IMMEDIATE
;

-- Reference: FK_CATEGORIES_MOVIES (table: categories)
ALTER TABLE categories ADD CONSTRAINT FK_CATEGORIES_MOVIES
    FOREIGN KEY (movie_id)
    REFERENCES movies (id)
    ON DELETE  CASCADE
    ON UPDATE  RESTRICT
    NOT DEFERRABLE
    INITIALLY IMMEDIATE
;

-- Reference: FK_RATINGS_MOVIES (table: ratings)
ALTER TABLE ratings ADD CONSTRAINT FK_RATINGS_MOVIES
    FOREIGN KEY (movie_id)
    REFERENCES movies (id)
    ON DELETE  CASCADE
    ON UPDATE  RESTRICT
    NOT DEFERRABLE
    INITIALLY IMMEDIATE
;

-- Reference: FK_RATINGS_USERS (table: ratings)
ALTER TABLE ratings ADD CONSTRAINT FK_RATINGS_USERS
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE  CASCADE
    ON UPDATE  RESTRICT
    NOT DEFERRABLE
    INITIALLY IMMEDIATE
;

-- Reference: FK_REVIEWS_MOVIES (table: reviews)
ALTER TABLE reviews ADD CONSTRAINT FK_REVIEWS_MOVIES
    FOREIGN KEY (movie_id)
    REFERENCES movies (id)
    ON DELETE  CASCADE
    ON UPDATE  RESTRICT
    NOT DEFERRABLE
    INITIALLY IMMEDIATE
;

-- Reference: FK_REVIEWS_USERS (table: reviews)
ALTER TABLE reviews ADD CONSTRAINT FK_REVIEWS_USERS
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE  CASCADE
    ON UPDATE  RESTRICT
    NOT DEFERRABLE
    INITIALLY IMMEDIATE
;

-- Reference: FK_ROLES_CELEBRITIES (table: roles)
ALTER TABLE roles ADD CONSTRAINT FK_ROLES_CELEBRITIES
    FOREIGN KEY (celebrity_id)
    REFERENCES celebrities (id)
    ON DELETE  CASCADE
    ON UPDATE  RESTRICT
    NOT DEFERRABLE
    INITIALLY IMMEDIATE
;

-- Reference: FK_ROLES_MOVIE (table: roles)
ALTER TABLE roles ADD CONSTRAINT FK_ROLES_MOVIE
    FOREIGN KEY (movie_id)
    REFERENCES movies (id)
    ON DELETE  CASCADE
    ON UPDATE  RESTRICT
    NOT DEFERRABLE
    INITIALLY IMMEDIATE
;

-- End of file.