package qbotservice

import "context"

type Bot interface {
	StartGame(ctx context.Context)
}
