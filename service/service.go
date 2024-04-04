package service

import (
	"go-mem-profile/repo"
	"strconv"
)

type Service interface {
	GetDataPointer() string
	GetDataValue() string
}

type service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetDataPointer() string {

	data := s.repo.GetDataPointer()
	counter := 0

	for i := 0; i < 100_000; i++ {
		counter += data.Age
	}

	// _ = data

	// return data
	return strconv.Itoa(counter)

}

func (s *service) GetDataValue() string {
	data := s.repo.GetDataValue()

	counter := 0

	for i := 0; i < 100_000; i++ {
		counter += data.Age
	}

	// _ = data

	return strconv.Itoa(counter)
	// return data
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
