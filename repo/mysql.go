package repo

import (
	"database/sql"
	"go-mem-profile/entity"
	"time"
)

type Repository interface {
	GetDataPointer() *entity.LargeData
	GetDataValue() entity.LargeData
}

// type Repository interface {
// 	GetDataPointer() *[]entity.LargeData
// 	GetDataValue() []entity.LargeData
// }

type repository struct {
	// db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{}
	// return &repository{
	// 	db: db,
	// }
	// return &repository{
	// 	db: db,
	// }
}

func (r *repository) GetDataPointer() *entity.LargeData {
	data := entity.LargeData{
		ID:          1,
		Name:        "John Doe",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Age:         30,
		Address:     "123 Main St",
		Email:       "john@example.com",
		Phone:       "123-456-7890",
		City:        "New York",
		State:       "NY",
		Country:     "USA",
		ZipCode:     "10001",
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Notes:       []string{"Note 1", "Note 2", "Note 3"},
		Tags:        []string{"tag1", "tag2", "tag3"},
		Metadata: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
			"key3": true,
		},
		ImageURL:  "https://example.com/image.jpg",
		Latitude:  40.7128,
		Longitude: -74.0060,
	}

	return &data
}

func (r *repository) GetDataValue() entity.LargeData {

	data := entity.LargeData{
		ID:          1,
		Name:        "John Doe",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Age:         30,
		Address:     "123 Main St",
		Email:       "john@example.com",
		Phone:       "123-456-7890",
		City:        "New York",
		State:       "NY",
		Country:     "USA",
		ZipCode:     "10001",
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Notes:       []string{"Note 1", "Note 2", "Note 3"},
		Tags:        []string{"tag1", "tag2", "tag3"},
		Metadata: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
			"key3": true,
		},
		ImageURL:  "https://example.com/image.jpg",
		Latitude:  40.7128,
		Longitude: -74.0060,
	}

	return data
}

// func (r *repository) GetDataPointer() *[]entity.LargeData {
// 	var datas []entity.LargeData
// 	// do database queries
// 	for i := 0; i < 100_000; i++ {
// 		data := entity.LargeData{
// 			ID:          i + 1,
// 			Name:        "John Doe",
// 			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
// 			Age:         30,
// 			Address:     "123 Main St",
// 			Email:       "john@example.com",
// 			Phone:       "123-456-7890",
// 			City:        "New York",
// 			State:       "NY",
// 			Country:     "USA",
// 			ZipCode:     "10001",
// 			IsActive:    true,
// 			CreatedAt:   time.Now(),
// 			UpdatedAt:   time.Now(),
// 			Notes:       []string{"Note 1", "Note 2", "Note 3"},
// 			Tags:        []string{"tag1", "tag2", "tag3"},
// 			Metadata: map[string]interface{}{
// 				"key1": "value1",
// 				"key2": 123,
// 				"key3": true,
// 			},
// 			ImageURL:  "https://example.com/image.jpg",
// 			Latitude:  40.7128,
// 			Longitude: -74.0060,
// 		}
// 		datas = append(datas, data)
// 	}

// 	return &datas
// }

// func (r *repository) GetDataValue() []entity.LargeData {
// 	var datas []entity.LargeData

// 	// do database queries
// 	for i := 0; i < 100_000; i++ {
// 		data := entity.LargeData{
// 			ID:          i + 1,
// 			Name:        "John Doe",
// 			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
// 			Age:         30,
// 			Address:     "123 Main St",
// 			Email:       "john@example.com",
// 			Phone:       "123-456-7890",
// 			City:        "New York",
// 			State:       "NY",
// 			Country:     "USA",
// 			ZipCode:     "10001",
// 			IsActive:    true,
// 			CreatedAt:   time.Now(),
// 			UpdatedAt:   time.Now(),
// 			Notes:       []string{"Note 1", "Note 2", "Note 3"},
// 			Tags:        []string{"tag1", "tag2", "tag3"},
// 			Metadata: map[string]interface{}{
// 				"key1": "value1",
// 				"key2": 123,
// 				"key3": true,
// 			},
// 			ImageURL:  "https://example.com/image.jpg",
// 			Latitude:  40.7128,
// 			Longitude: -74.0060,
// 		}
// 		datas = append(datas, data)
// 	}

// 	return datas
// }
