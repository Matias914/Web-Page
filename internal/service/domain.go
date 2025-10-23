package service

import (
	"time"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

const MaxMovieTitleLength = 255
const MaxGenresNameLength = 255
const MaxRoleLength = 255
const MaxUsernameLength = 255
const MaxPasswordLength = 255
const MaxSynopsisLength = 5000
const MaxCommentLength = 5000

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
	GenreID int `json:"genre_id"`
	MovieID int `json:"movie_id"`
}

type Celebrity struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
}

type Genre struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Movie struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Synopsis        string    `json:"synopsis"`
	ReleasedAt      time.Time `json:"released_at"`
	DurationMinutes int       `json:"duration_minutes"`
	PosterUrl       string    `json:"poster_url"`
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
	Title           string    `json:"title"`
	Synopsis        string    `json:"synopsis"`
	ReleasedAt      time.Time `json:"released_at"`
	DurationMinutes int       `json:"duration_minutes"`
	PosterUrl       string    `json:"poster_url"`
}
