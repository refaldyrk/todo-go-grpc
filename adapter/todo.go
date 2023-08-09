package adapter

import (
	"context"

	"github.com/refaldyrk/todo-go-grpc/connector/todo"
	"github.com/refaldyrk/todo-go-grpc/service"
)

type TodoAdapter struct {
	services *service.CacheService
}

func NewAdapter(services *service.CacheService) *TodoAdapter {
	return &TodoAdapter{services: services}
}

func (adapter *TodoAdapter) AddTodo(ctx context.Context, req *todo.AddTodoRequest) (*todo.TodoResponse, error) {
	resp, err := adapter.services.AddTodo(ctx, req.Todo.Name)
	if err != nil {
		return &todo.TodoResponse{
			Message: err.Error(),
			Data:    &resp,
			Error:   true,
		}, err
	}

	return &todo.TodoResponse{
		Message: "success add todo",
		Data:    &resp,
		Error:   false,
	}, nil
}

func (adapter *TodoAdapter) UpdateTodo(ctx context.Context, req *todo.UpdateTodoRequest) (*todo.TodoResponse, error) {
	resp, err := adapter.services.UpdateTodo(ctx, req.Todo)
	if err != nil {
		return &todo.TodoResponse{
			Message: err.Error(),
			Data:    &resp,
			Error:   true,
		}, err
	}

	return &todo.TodoResponse{
		Message: "success update todo",
		Data:    &resp,
		Error:   false,
	}, nil
}
func (adapter *TodoAdapter) DeleteTodo(ctx context.Context, req *todo.DeleteTodoRequest) (*todo.TodoResponse, error) {
	resp, err := adapter.services.DeleteTodo(ctx, req.Id)
	if err != nil {
		return &todo.TodoResponse{
			Message: err.Error(),
			Data:    &resp,
			Error:   true,
		}, err
	}

	return &todo.TodoResponse{
		Message: "success delete todo",
		Data:    &resp,
		Error:   false,
	}, nil
}
