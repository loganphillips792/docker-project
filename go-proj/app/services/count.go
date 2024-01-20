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

func NewCount(log *slog.Logger, cs *db.CountStore) Count {
	return Count{
		Log:        log,
		CountStore: cs,
	}
}

type Count struct {
	Log        *slog.Logger
	CountStore *db.CountStore
}

func (cs Count) Increment(ctx context.Context) {
	fmt.Println("INCREMENT")
}
