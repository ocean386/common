package zorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
	SlowThreshold time.Duration
	mode          string
}

func NewGormLogger(c string) *GormLogger {
	return &GormLogger{
		SlowThreshold: 200 * time.Millisecond, // 一般超过200毫秒就算慢查所以不使用配置进行更改
		mode:          c,
	}
}

var _ logger.Interface = (*GormLogger)(nil)

func (l *GormLogger) LogMode(lev logger.LogLevel) logger.Interface {
	return &GormLogger{}
}
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Infof(msg, data)
}
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Errorf(msg, data)
}
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Errorf(msg, data)
}
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// 获取运行时间
	elapsed := time.Since(begin)
	// 获取 SQL 语句和返回条数
	sql, rows := fc()
	// 通用字段
	logFields := []logx.LogField{
		logx.Field("sql", sql),
		logx.Field("time", MicrosecondsStr(elapsed)),
		logx.Field("rows", rows),
	}
	// Gorm 错误
	if err != nil {
		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logx.WithContext(ctx).Infow("Database ErrRecordNotFound", logFields...)
		} else {
			// 其他错误使用 error 等级
			logFields = append(logFields, logx.Field("catch error", err))
			logx.WithContext(ctx).Errorw("Database Error", logFields...)
		}
	}
	// 慢查询日志
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		logx.WithContext(ctx).Sloww("Database Slow Log", logFields...)
	}
	// 非生产模式下，记录所有 SQL 请求
	if l.mode != service.ProMode && sql != "SHOW STATUS" {
		logx.WithContext(ctx).Infow("Database Query", logFields...)
	}
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}
