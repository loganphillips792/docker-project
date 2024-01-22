package services

import (
	"context"
	"fmt"
	"github.com/loganphillips792/kubernetes-project/db"
	"log/slog"
)

type Counts struct {
	Global int
}

func NewCountService(log *slog.Logger, cs *db.CountStore) CountServices {
	return CountServices{
		Log:        log,
		CountStore: cs,
	}
}

type CountServices struct {
	Log        *slog.Logger
	CountStore *db.CountStore
}

func (cs CountServices) Increment(ctx context.Context) {
	fmt.Println("INCREMENT")
}
