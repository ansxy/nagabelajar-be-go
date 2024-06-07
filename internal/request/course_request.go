package request

import "github.com/google/uuid"

type ListCourseRequest struct {
	BaseQuery
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
	UserID     string `json:"-"`
}

type UpsertCourseRequest struct {
	Name        string `json:"name" validate:"required"`
	Code        string `json:"_"`
	CategoryID  int    `json:"category_id" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsPaid      *bool  `json:"is_paid"`
	IsArchived  *bool  `json:"is_archived"`
	MediaID     string `json:"media_id" validate:"required"`
	Author      string `json:"author" validate:"required"`
}

type TakeCourseRequest struct {
	UserID   string `json:"-"`
	CourseID string `json:"-"`
}

type GetOneCourseRequest struct {
	CourseID string     `json:"course_id" validate:"required"`
	UserID   *uuid.UUID `json:"-"`
}

type UpsertCourseContentRequest struct {
	CourseContentID *int                           `json:"course_content_id"`
	Title           string                         `json:"title" validate:"required"`
	Content         string                         `json:"content" validate:"required"`
	SubContent      *UpsertSubCourseContentRequest `json:"sub_content"`
	Exercise        []UpsertCourseExerciseRequest  `json:"exercise"`
}

type UpsertSubCourseContentRequest struct {
	SubContentID    *int   `json:"sub_content_id"`
	CourseContentID int    `json:"course_content_id"`
	Title           string `json:"title" validate:"required"`
	Content         string `json:"content" validate:"required"`
}

type UpsertCourseExerciseRequest struct {
	ExerciseID      *int   `json:"exercise_id"`
	CourseContentID int    `json:"course_content_id"`
	Title           string `json:"title" validate:"required"`
	Content         string `json:"content" validate:"required"`
}

type UpserCourseDetailRequest struct {
	CourseID      string                       `json:"course_id" validate:"required"`
	Name          string                       `json:"name" validate:"required"`
	Position      int                          `json:"position" validate:"required"`
	Objective     string                       `json:"objective" validate:"required"`
	CourseContent []UpsertCourseContentRequest `json:"course_content" validate:"required"`
}

type EnrollCourseRequest struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	CourseID string    `json:"course_id" validate:"required"`
}

type UpdateProgressRequest struct {
	ProgressID int    `json:"progress_id" validate:"required"`
	UserID     string `json:"-" validate:"required"`
}
