// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "new-proto/api/models"

	mock "github.com/stretchr/testify/mock"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

// CompleteTask provides a mock function with given fields: id
func (_m *TaskRepository) CompleteTask(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: task
func (_m *TaskRepository) Create(task models.Task) int {
	ret := _m.Called(task)

	var r0 int
	if rf, ok := ret.Get(0).(func(models.Task) int); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *TaskRepository) GetAll() []models.Task {
	ret := _m.Called()

	var r0 []models.Task
	if rf, ok := ret.Get(0).(func() []models.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Task)
		}
	}

	return r0
}

// GetById provides a mock function with given fields: id
func (_m *TaskRepository) GetById(id int) (models.Task, error) {
	ret := _m.Called(id)

	var r0 models.Task
	if rf, ok := ret.Get(0).(func(int) models.Task); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: task
func (_m *TaskRepository) UpdateTask(task models.Task) error {
	ret := _m.Called(task)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
