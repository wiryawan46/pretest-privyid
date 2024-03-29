package main

import (
	"pretest-privyid/config"
	CategoryPresenterPackage "pretest-privyid/modules/v1/category/presenter"
	CategoryRepoPackage "pretest-privyid/modules/v1/category/repository"
	CategoryUseCasePackage "pretest-privyid/modules/v1/category/usecase"
)

// Service structure
type Service struct {
	CategoryHandler *CategoryPresenterPackage.HTTPCategoryHandler
	CategoryUsecase CategoryUseCasePackage.CategoryUsecase
}

// MakeHandler function for initializing service
func MakeHandler() *Service {

	dbConnection := config.ConnectDB()

	CategoryRepo := CategoryRepoPackage.NewWorkCategoryRepoPostgres(dbConnection)
	CategoryUsecase := CategoryUseCasePackage.NewWorkCategoryUseCase(CategoryRepo)
	CategoryHandler := CategoryPresenterPackage.NewHTTPHandler(CategoryUsecase)

	return &Service{
		CategoryHandler: CategoryHandler,
		CategoryUsecase: CategoryUsecase,
	}
}
