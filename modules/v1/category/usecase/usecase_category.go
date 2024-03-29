package usecase

import (
	"fmt"
	"log"
	"pretest-privyid/modules/v1/category/model"
	"pretest-privyid/modules/v1/category/repository"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type CategoryUsecaseImpl struct {
	CategoryRepository repository.CategoryRepository
}

func NewWorkCategoryUseCase(CategoryRepo repository.CategoryRepository) CategoryUsecase {
	return &CategoryUsecaseImpl{
		CategoryRepository: CategoryRepo,
	}
}

func (rp *CategoryUsecaseImpl) CreateCategory(param model.Category) ResultUseCase {
	output := ResultUseCase{}

	saveResult := rp.CategoryRepository.CreateCategory(param)
	if saveResult.Error != nil {
		log.Println("Error menyimpan data :", saveResult.Error.Error())
		output = ResultUseCase{Error: saveResult.Error}
		return output
	}
	data, ok := saveResult.Result.(model.Category)
	if !ok {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		output = ResultUseCase{Error: err}
		return output
	}
	output = ResultUseCase{Result: data}
	return output
}

