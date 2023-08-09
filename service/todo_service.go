package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/refaldyrk/todo-go-grpc/connector/todo"
)

type CacheService struct {
	cache *cache.Cache
}

func NewCacheService(cch *cache.Cache) *CacheService {
	return &CacheService{cch}
}

func (cs *CacheService) AddTodo(ctx context.Context, name string) (todo.Todo, error) {
	if name == "" {
		return todo.Todo{}, errors.New("name can't be empty")
	}

	newTodo := todo.Todo{
		Id:     uuid.NewString(),
		Name:   name,
		IsDone: false,
	}

	err := cs.cache.Add(newTodo.Id, newTodo, 24*time.Hour)
	if err != nil {
		return todo.Todo{}, err
	}

	return newTodo, nil
}

func (cs *CacheService) UpdateTodo(ctx context.Context, req *todo.Todo) (todo.Todo, error) {
	if req.Id == "" || req.Name == "" {
		return todo.Todo{}, errors.New("can't be empty")
	}

	updatedTodo := todo.Todo{
		Id:     req.Id,
		Name:   req.Name,
		IsDone: req.IsDone,
	}

	err := cs.cache.Replace(req.Id, updatedTodo, 24*time.Hour)
	if err != nil {
		return todo.Todo{}, err
	}

	return updatedTodo, nil
}

func (cs *CacheService) DeleteTodo(ctx context.Context, id string) (todo.Todo, error) {
	if id == "" {
		return todo.Todo{}, errors.New("id can't be empty")
	}
	//Get Result
	data, found := cs.cache.Get(id)
	if !found {
		return todo.Todo{}, errors.New("not found")
	}

	cs.cache.Delete(id)

	return data.(todo.Todo), nil
}
