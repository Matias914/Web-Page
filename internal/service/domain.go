package service

import (
	"time"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

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
	ID              int       `json:"id" example:"1"`
	Title           string    `json:"title" example:"Lord of the Strings"`
	Synopsis        string    `json:"synopsis" example:"A man that can manipulate every string variable"`
	ReleasedAt      time.Time `json:"released_at" example:"2025-10-30T00:00:00Z"`
	DurationMinutes int       `json:"duration_minutes" example:"255"`
	PosterUrl       string    `json:"poster_url" example:"https://imgs.search.brave.com/8beLkj9LsaTyqIIEdbEELICTyQJkSQnTz6UQfl4x5oI/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pLmVi/YXlpbWcuY29tL2lt/YWdlcy9nL3hwNEFB/ZVN3cGNwb2xKRFQv/cy1sMjI1LmpwZw"`
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
