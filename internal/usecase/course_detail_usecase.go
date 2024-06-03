package usecase

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
)

// CreateCourseDetail implements IFaceUsecase.
func (u *Usecase) CreateCourseDetail(ctx context.Context, data *request.UpserCourseDetailRequest) error {
	var courseDetail *model.CourseDetail

	course, err := u.Repo.FindOneCourse(ctx, &request.GetOneCourseRequest{
		CourseID: data.CourseID,
	})

	if err != nil {
		return err
	}

	courseDetail = &model.CourseDetail{
		Name:      data.Name,
		Position:  data.Position,
		CourseID:  course.CourseID,
		Objective: data.Objective,
	}

	for _, content := range data.CourseContent {
		courseContent := model.CourseContent{
			Title:         content.Title,
			CourseContent: content.Content,
			SubContent: []model.CourseSubContent{
				{
					Title:   content.SubContent.Title,
					Content: content.SubContent.Content,
				},
			},
		}

		for _, exercise := range content.Exercise {
			courseContent.Exercise = append(courseContent.Exercise, model.CourseExercise{
				Title:   exercise.Title,
				Content: exercise.Content,
			})
		}

		courseDetail.Content = append(courseDetail.Content, courseContent)

	}

	// for _, content := range data.CourseContent {
	// 	courseDetail.Content = append(courseDetail.Content, model.CourseContent{
	// 		Title:         content.Title,
	// 		CourseContent: content.Content,
	// 		SubContent: []model.CourseSubContent{
	// 			{
	// 				Title:   content.SubContent.Title,
	// 				Content: content.SubContent.Content,
	// 			},
	// 		},
	// 		Exercise: []model.CourseExercise{
	// 			{
	// 				Title:   content.Exercise.Title,
	// 				Content: content.Exercise.Content,
	// 			},
	// 		},
	// 	})
	// }

	return u.Repo.CreateCourseDetail(ctx, courseDetail)

}

// FindCourseDetail implements IFaceUsecase.
func (u *Usecase) FindCourseDetail(ctx context.Context, courseDetailID int) (*model.CourseDetail, error) {
	return u.Repo.FindOneCourseDetail(ctx, courseDetailID)
}
