package service

import (
	"errors"
	"time"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

/* ----------------------------------------------------------------------------- */
/*						       	     ERRORS  	        				         */
/* ----------------------------------------------------------------------------- */

const ErrCodeConstraintPK = "23505"
const ErrCodeConstraintCHK = "23514"
const ErrCodeConstraintFK = "23503"

var ErrCategoryNotFound = errors.New("the movie is not associated with the genre")
var ErrCategoryDuplicated = errors.New("the movie is already associated with the genre")
var ErrInvalidCategory = errors.New("some category fields were rejected")

var ErrCelebrityNotFound = errors.New("the celebrity does not exist")
var ErrCelebrityDuplicated = errors.New("the celebrity already exists")
var ErrInvalidCelebrity = errors.New("some celebrity fields were rejected")
var ErrCelebrityReferenced = errors.New("the celebrity cannot be deleted because it is referenced by other records")

var ErrGenreNotFound = errors.New("the genre does not exist")
var ErrGenreDuplicated = errors.New("the genre already exists")
var ErrInvalidGenre = errors.New("some genre fields were rejected")
var ErrGenreReferenced = errors.New("the genre cannot be deleted because it is referenced by other records")

var ErrMovieNotFound = errors.New("the movie does not exist")
var ErrMovieDuplicated = errors.New("the movie already exists")
var ErrInvalidMovie = errors.New("some movie fields were rejected")
var ErrMovieReferenced = errors.New("the movie cannot be deleted because it is referenced by other records")

var ErrRatingNotFound = errors.New("the rating does not exist")
var ErrRatingDuplicated = errors.New("the rating already exists")
var ErrInvalidRating = errors.New("some rating fields were rejected")

var ErrReviewNotFound = errors.New("the review does not exist")
var ErrReviewDuplicated = errors.New("the review already exists")
var ErrInvalidReview = errors.New("some review fields were rejected")

var ErrRoleNotFound = errors.New("the role does not exist")
var ErrRoleDuplicated = errors.New("the role already exists")
var ErrInvalidRole = errors.New("some role fields were rejected")

var ErrUserNotFound = errors.New("the user does not exist")
var ErrUserDuplicated = errors.New("the user already exists")
var ErrInvalidUser = errors.New("some user fields were rejected")
var ErrUserReferenced = errors.New("the user cannot be deleted because it is referenced by other records")

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
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"Linus Torvalds"`
	BirthDate time.Time `json:"birth_date" example:"1969-12-28T00:00:00Z"`
}

type Genre struct {
	ID   int32  `json:"id" example:"1"`
	Name string `json:"name" example:"horror"`
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
	UserID    int       `json:"user_id" example:"1"`
	MovieID   int       `json:"movie_id" example:"2"`
	Rating    int       `json:"rating" example:"5"`
	CreatedAt time.Time `json:"created_at" example:"2025-10-30T00:00:00Z"`
}

type Review struct {
	UserID    int       `json:"user_id" example:"1"`
	MovieID   int       `json:"movie_id" example:"2"`
	Comment   string    `json:"comment" example:"Seriously... Who made this? "`
	CreatedAt time.Time `json:"created_at" example:"2025-10-30T00:00:00Z"`
}

type Role struct {
	CelebrityID int    `json:"celebrity_id" example:"1"`
	MovieID     int    `json:"movie_id" example:"2"`
	Role        string `json:"role" example:"Tech Guy"`
}

type User struct {
	ID        int       `json:"id" example:"1"`
	Username  string    `json:"username" example:"thescrummaster1"`
	Password  string    `json:"password" example:"notofyourbusiness"`
	Mail      string    `json:"mail" example:"smokeseller@gmail.com"`
	CreatedAt time.Time `json:"created_at" example:"2025-10-30T00:00:00Z"`
}

/* ----------------------------------------------------------------------------- */
/*						       	     INPUTS  	        				         */
/* ----------------------------------------------------------------------------- */

type CategoryData struct {
	GenreID int `json:"genre_id" example:"1" validate:"required,gt=0"`
}

type CelebrityData struct {
	Name      string    `json:"name" example:"Linus Torvalds" validate:"required,max=255"`
	BirthDate time.Time `json:"birth_date" example:"1969-12-28T00:00:00Z" validate:"required,before_or_today"`
}

type GenreData struct {
	Name string `json:"name" example:"horror" validate:"required,max=255"`
}

type MovieData struct {
	Title           string    `json:"title" example:"Lord of the Strings" validate:"required,max=255"`
	Synopsis        string    `json:"synopsis" example:"A man that can manipulate every string variable" validate:"required,max=5000"`
	ReleasedAt      time.Time `json:"released_at" example:"2025-10-30T00:00:00Z" validate:"required,before_or_today"`
	DurationMinutes int       `json:"duration_minutes" example:"255" validate:"required,gt=0"`
	PosterUrl       string    `json:"poster_url" example:"https://www.amazon.nl/Lord-Strings-J-Khan/dp/B000BU9TAQ" validate:"omitempty"`
}

type RatingData struct {
	MovieID int `json:"movie_id" example:"2" validate:"required,gt=0"`
	Rating  int `json:"rating" example:"5" validate:"required,gt=0,lte=10"`
}

type UpdatableRatingData struct {
	Rating int `json:"rating" example:"5" validate:"required,gt=0,lt=11"`
}

type ReviewData struct {
	MovieID int    `json:"movie_id" example:"2" validate:"required,gt=0"`
	Comment string `json:"comment" example:"Seriously... Who made this?" validate:"required,max=255"`
}

type UpdatableReviewData struct {
	Comment string `json:"comment" example:"Seriously... Who made this?" validate:"required,max=255"`
}

type RoleData struct {
	CelebrityID int    `json:"celebrity_id" example:"1" validate:"required,gt=0"`
	Role        string `json:"role" example:"Tech Guy" validate:"required,max=255"`
}

type UpdatableRoleData struct {
	Role string `json:"role" example:"Tech Guy" validate:"required,max=255"`
}

type UserData struct {
	Username string `json:"username" example:"thescrummaster1" validate:"required,max=255"`
	Password string `json:"password" example:"notofyourbusiness" validate:"required,max=255"`
	Mail     string `json:"mail" example:"smokeseller@gmail.com" validate:"required,max=255"`
}

type UpdatableUserData struct {
	Username string `json:"username" example:"thescrummaster1" validate:"required,max=255"`
	Mail     string `json:"mail" example:"smokeseller@gmail.com" validate:"required,max=255"`
}
