package application

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"strconv"

	"github.com/devpablocristo/nanlabs/application/port"
	"github.com/devpablocristo/nanlabs/domain"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

type TaskService struct {
	storage       port.Storage
	trelloService port.TrelloService
}

func NewTaskService(s port.Storage, t port.TrelloService) *TaskService {
	return &TaskService{
		storage:       s,
		trelloService: t,
	}
}

func (ps *TaskService) GetTasks(ctx context.Context) (map[string]*domain.Task, error) {
	Tasks := ps.storage.ListTasks(ctx)
	return Tasks, nil
}

func (ps *TaskService) GetTask(ctx context.Context, UUID string) (*domain.Task, error) {
	p, err := ps.storage.GetTask(ctx, UUID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *TaskService) CreateTask(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	task.UUID = uuid.New().String()
	err := ps.storage.SaveTask(ctx, task)
	if err != nil {
		return &domain.Task{}, err
	}
	return task, nil
}

func (ps *TaskService) UpdateTask(ctx context.Context, UUID string) error {
	return ps.storage.UpdateTask(ctx, UUID)
}

func (ps *TaskService) DeleteTask(ctx context.Context, UUID string) error {
	return ps.storage.DeleteTask(ctx, UUID)
}

func (ps *TaskService) CreateCard(ctx context.Context, task *domain.Task) error {

	task, err := ps.CreateTask(ctx, task)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return err
	}

	switch task.Type {
	case "issue":
		err := ps.trelloService.CreateIssueCard(ctx, task)
		if err != nil {
			return err
		}
		// // just memdb, not really necesary
		// err = ps.storage.SaveTask(ctx, task)
		// if err != nil {
		// 	return err
		// }
		return nil
	case "bug":
		createBugTitle(task)
		err := ps.trelloService.CreateBugCard(ctx, task)
		if err != nil {
			return err
		}
		return nil
	case "task":
		err := ps.trelloService.CreateTaskCard(ctx, task)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("incorrect task type")
	}

}

func createBugTitle(task *domain.Task) {
	titleWord := randstr.String(5, "abcdefghijklmnopqrstuvwxyz")
	titleNumber := strconv.Itoa(rand.Intn(999999))
	task.Title = "bug" + "-" + titleWord + "-" + titleNumber
}
