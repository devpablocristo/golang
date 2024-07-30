package port

import (
	"context"

	"github.com/devpablocristo/nanlabs/domain"
)

type TrelloService interface {
	CreateIssueCard(context.Context, *domain.Task) error
	CreateBugCard(context.Context, *domain.Task) error
	CreateTaskCard(context.Context, *domain.Task) error
}
