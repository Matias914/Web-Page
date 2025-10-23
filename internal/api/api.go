package api

import (
	"github.com/Matias914/Web-Page/internal/service"
)

type API struct {
	ValidationService *service.ValidationService
	CategoryService   *service.CategoryService
	CelebrityService  *service.CelebrityService
	GenreService      *service.GenreService
	MovieService      *service.MovieService
	RatingService     *service.RatingService
	ReviewService     *service.ReviewService
	RoleService       *service.RoleService
	UserService       *service.UserService
}

type JsonError struct {
	Mssg string `json:"message"`
}
