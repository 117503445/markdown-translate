package cache

import (
	badger "github.com/dgraph-io/badger/v4"
	"github.com/rs/zerolog/log"
)

type BadgerCache struct {
	db *badger.DB
}

type BadgerConfig struct {
	Dir string
}

func NewBadgerWithConfig(cfg *BadgerConfig) *BadgerCache {
	const DEFAULT_BADGER_DIR = "./data/badger"
	if cfg.Dir == "" {
		cfg.Dir = DEFAULT_BADGER_DIR
	}

	db, err := badger.Open(badger.DefaultOptions(cfg.Dir))
	if err != nil {
		log.Warn().Err(err).Msg("failed to open badger db")
	}
	return &BadgerCache{db: db}
}

func NewBadgerCache() *BadgerCache {
	return NewBadgerWithConfig(&BadgerConfig{})
}

func (b *BadgerCache) Get(source string) string {
	var result string

	if err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(source))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			result = string(val)
			return nil
		})
		return err
	}); err != nil {
		log.Warn().Err(err).Str("source", source).Msg("failed to get cache")
	}

	if result != "" {
		log.Debug().Str("source", source).Str("result", result).Msg("cache hit")
		return result
	} else {
		log.Debug().Str("source", source).Msg("cache miss")
		return ""
	}
}

func (b *BadgerCache) Set(source string, result string) {
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(source), []byte(result))
	})
	if err != nil {
		log.Warn().Err(err).Str("source", source).Str("result", result).Msg("failed to set cache")
	}
}
