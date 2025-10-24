package service

import (
	"errors"
	"time"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

/* ----------------------------------------------------------------------------- */
/*						       	     ERRORS  	        				         */
/* ----------------------------------------------------------------------------- */

var ErrMovieNotFound = errors.New("the movie does not exist")
var ErrMovieDuplicated = errors.New("the movie already exists")
var ErrInvalidMovie = errors.New("some movie fields were rejected")
var ErrMovieReferenced = errors.New("the movie cannot be deleted because it is referenced by other records")

var ErrCategoryNotFound = errors.New("the movie is not associated with the genre")
var ErrCategoryDuplicated = errors.New("the movie is already associated with the genre")

var ErrGenreNotFound = errors.New("the genre does not exist")
var ErrGenreDuplicated = errors.New("the genre already exists")
var ErrInvalidGenre = errors.New("some genre fields were rejected")
var ErrGenreReferenced = errors.New("the genre cannot be deleted because it is referenced by other records")

/* ----------------------------------------------------------------------------- */
/*						       	     SERVICES  	        				         */
/* ----------------------------------------------------------------------------- */

type CategoryService struct {
	Queries *sqlc.Queries
}

type CelebrityService struct {
	Queries *sqlc.Queries
}

type GenreService struct {
	Queries *sqlc.Queries
}

type MovieService struct {
	Queries *sqlc.Queries
}

type RatingService struct {
	Queries *sqlc.Queries
}

type ReviewService struct {
	Queries *sqlc.Queries
}

type RoleService struct {
	Queries *sqlc.Queries
}

type UserService struct {
	Queries *sqlc.Queries
}

/* ----------------------------------------------------------------------------- */
/*						       	     ENTITIES  	        				         */
/* ----------------------------------------------------------------------------- */

type Category struct {
	GenreID int `json:"genre_id" example:"1"`
	MovieID int `json:"movie_id" example:"2"`
}

type Celebrity struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
}

type Genre struct {
	ID   int32  `json:"id" example:"1"`
	Name string `json:"name" example:"terror"`
}

type Movie struct {
	ID              int       `json:"id" example:"1"`
	Title           string    `json:"title" example:"Lord of the Strings"`
	Synopsis        string    `json:"synopsis" example:"A man that can manipulate every string variable"`
	ReleasedAt      time.Time `json:"released_at" example:"2025-10-30T00:00:00Z"`
	DurationMinutes int       `json:"duration_minutes" example:"255"`
	PosterUrl       string    `json:"poster_url" example:"https://www.amazon.nl/Lord-Strings-J-Khan/dp/B000BU9TAQ"`
}

type Rating struct {
	UserID    int       `json:"user_id"`
	MovieID   int       `json:"movie_id"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

type Review struct {
	UserID    int       `json:"user_id"`
	MovieID   int       `json:"movie_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type Role struct {
	CelebrityID int    `json:"celebrity_id"`
	MovieID     int    `json:"movie_id"`
	Role        string `json:"role"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Mail      string    `json:"mail"`
	CreatedAt time.Time `json:"created_at"`
}

/* ----------------------------------------------------------------------------- */
/*						       	     INPUTS  	        				         */
/* ----------------------------------------------------------------------------- */

type AddMovieInput struct {
	Title           string    `json:"title" validate:"required,max=255"`
	Synopsis        string    `json:"synopsis" validate:"required,max=5000"`
	ReleasedAt      time.Time `json:"released_at" validate:"required,before_or_today"`
	DurationMinutes int       `json:"duration_minutes" validate:"required,gte=1"`
	PosterUrl       string    `json:"poster_url" validate:"omitempty"`
}

type UpdateMovieInput struct {
	Title           string    `json:"title" validate:"required,max=255"`
	Synopsis        string    `json:"synopsis" validate:"required,max=5000"`
	ReleasedAt      time.Time `json:"released_at" validate:"required,before_or_today"`
	DurationMinutes int       `json:"duration_minutes" validate:"required,gte=1"`
	PosterUrl       string    `json:"poster_url" validate:"omitempty"`
}

type AddCategoryInput struct {
	GenreID int `json:"genre_id" validate:"required,gt=1"`
}

type AddGenreInput struct {
	Name string `json:"name" validate:"required,max=255"`
}

type UpdateGenreInput struct {
	Name string `json:"name" validate:"required,max=255"`
}
