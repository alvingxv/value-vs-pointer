package service

import (
	"go-mem-profile/entity"
	"go-mem-profile/repo"
)

type Service interface {
	GetDataPointer() *entity.LargeData
	GetDataValue() entity.LargeData
}

type service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetDataPointer() *entity.LargeData {
	// var datas []interface{}

	data := s.repo.GetDataPointer()
	// counter := 0

	// for i := 0; i < 100_000; i++ {
	// 	counter += data.Age

	// 	// datas = append(datas, data)
	// }

	// _ = data

	return data
	// return strconv.Itoa(counter)

}

func (s *service) GetDataValue() entity.LargeData {
	// var datas []interface{}
	data := s.repo.GetDataValue()

	// counter := 0

	// for i := 0; i < 100_000; i++ {
	// 	counter += data.Age
	// 	// datas = append(datas, data)
	// }

	// // _ = data

	return data
}

// func (s *service) GetDataPointer() string {

// 	data := s.repo.GetDataPointer()

// 	counter := 0

// 	for _, d := range *data {

// 		for i := 0; i < 100; i++ {
// 			counter += d.Age
// 		}
// 	}

// 	return strconv.Itoa(counter)

// }

// func (s *service) GetDataValue() string {

// 	data := s.repo.GetDataValue()

// 	counter := 0

// 	for _, d := range data {

// 		for i := 0; i < 100; i++ {
// 			counter += d.Age
// 		}
// 	}

// 	return strconv.Itoa(counter)
// }
