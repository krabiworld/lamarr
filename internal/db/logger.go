package db

import (
	"context"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
	"time"
)

type Logger struct{}

func (l Logger) LogMode(_ logger.LogLevel) logger.Interface {
	return l
}

func (l Logger) Info(ctx context.Context, s string, i ...interface{}) {
	log.Ctx(ctx).Info().Msgf(s, i...)
}

func (l Logger) Warn(ctx context.Context, s string, i ...interface{}) {
	log.Ctx(ctx).Warn().Msgf(s, i...)
}

func (l Logger) Error(ctx context.Context, s string, i ...interface{}) {
	log.Ctx(ctx).Error().Msgf(s, i...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	log.Ctx(ctx).Trace().
		Err(err).
		Dur("begin", time.Since(begin)).
		Str("sql", sql).
		Int64("rows", rows).
		Send()
}
