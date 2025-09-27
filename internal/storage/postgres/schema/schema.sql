-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2025-09-20 20:08:09.162

-- tables
-- Table: categories
CREATE TABLE categories (
    genre_id int  NOT NULL,
    movie_id bigint  NOT NULL,
    CONSTRAINT PK_CATEGORIES PRIMARY KEY (movie_id,genre_id)
);

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
    CONSTRAINT PK_RATINGS PRIMARY KEY (movie_id,user_id)
);

-- Table: reviews
CREATE TABLE reviews (
    user_id bigint  NOT NULL,
    movie_id bigint  NOT NULL,
    comment text  NOT NULL,
    created_at timestamp with time zone  NOT NULL DEFAULT NOW(),
    -- SOLUCIÓN EN 'reviews'
    CONSTRAINT comment_check CHECK (length(trim(comment)) > 0) NOT DEFERRABLE INITIALLY IMMEDIATE,
    CONSTRAINT PK_REVIEWS PRIMARY KEY (user_id,movie_id)
);

-- Table: roles
CREATE TABLE roles (
    movie_id bigint  NOT NULL,
    celebrity_id bigint  NOT NULL,
    role varchar(255)  NOT NULL,
    CONSTRAINT PK_ROLES PRIMARY KEY (movie_id,celebrity_id)
);

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

-- Reference: PK_ROLES_CELEBRITIES (table: roles)
ALTER TABLE roles ADD CONSTRAINT PK_ROLES_CELEBRITIES
    FOREIGN KEY (celebrity_id)
    REFERENCES celebrities (id)
    ON DELETE  CASCADE
    ON UPDATE  RESTRICT
    NOT DEFERRABLE
    INITIALLY IMMEDIATE
;

-- Reference: PK_ROLES_MOVIE (table: roles)
ALTER TABLE roles ADD CONSTRAINT PK_ROLES_MOVIE
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

-- End of file.