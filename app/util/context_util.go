package util

import (
	"context"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

type CtxKey string

const (
	QueriesKey CtxKey = "queries"
)

func WithQueries(ctx context.Context, q *dbgen.Queries) context.Context {
	return context.WithValue(ctx, QueriesKey, q)
}

func GetQueries(ctx context.Context) *dbgen.Queries {
	queries, ok := ctx.Value(QueriesKey).(*dbgen.Queries)
	if !ok {
		return nil
	}
	return queries
}
