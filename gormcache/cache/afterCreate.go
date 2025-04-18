package cache

import (
	"github.com/ocean386/common/gormcache/config"
	"github.com/ocean386/common/gormcache/util"
	"gorm.io/gorm"
)

func AfterCreate(cache *Gorm2Cache) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		if db.RowsAffected == 0 {
			return // no rows affected, no need to invalidate cache
		}

		tableName := ""
		if db.Statement.Schema != nil {
			tableName = db.Statement.Schema.Table
		} else {
			tableName = db.Statement.Table
		}
		ctx := db.Statement.Context

		if db.Error == nil && cache.Config.InvalidateWhenUpdate && util.ShouldCache(tableName, cache.Config.Tables) {
			if cache.Config.CacheLevel == config.CacheLevelAll || cache.Config.CacheLevel == config.CacheLevelOnlySearch {
				invalidSearchCache := func() {
					// We invalidate search cache here,
					// because any newly created objects may cause search cache results to be outdated and invalid.
					cache.Logger.CtxInfo(ctx, "[AfterCreate] now start to invalidate search cache for table: %s", tableName)
					err := cache.InvalidateSearchCache(ctx, tableName)
					if err != nil {
						cache.Logger.CtxError(ctx, "[AfterCreate] invalidating search cache for table %s error: %v",
							tableName, err)
						return
					}
					cache.Logger.CtxInfo(ctx, "[AfterCreate] invalidating search cache for table: %s finished.", tableName)
				}
				if cache.Config.AsyncWrite {
					go invalidSearchCache()
				} else {
					invalidSearchCache()
				}
			}
		}
	}
}
