package cache

import (
	"github.com/117503445/markdown-translate/pkg/translator"

	badger "github.com/dgraph-io/badger/v4"
	"github.com/rs/zerolog/log"
)

type BadgerCache struct {
	innerProvider translator.Provider
	db            *badger.DB
}

type BadgerConfig struct {
	Dir string
}

func NewBadgerWithConfig(p translator.Provider, cfg *BadgerConfig) *BadgerCache {
	const DEFAULT_BADGER_DIR = "./data/badger"
	if cfg.Dir == "" {
		cfg.Dir = DEFAULT_BADGER_DIR
	}

	db, err := badger.Open(badger.DefaultOptions(cfg.Dir))
	if err != nil {
		log.Warn().Err(err).Msg("failed to open badger db")
	}
	return &BadgerCache{innerProvider: p, db: db}
}

func NewBadgerCache(p translator.Provider) *BadgerCache {
	return NewBadgerWithConfig(p, &BadgerConfig{})
}

func (b *BadgerCache) Translate(source string) (string, error) {
	var result string

	b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(source))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			result = string(val)
			return nil
		})
		return err
	})
	if result != "" {
		log.Debug().Str("source", source).Str("result", result).Msg("cache hit")
		return result, nil
	}

	translated, err := b.innerProvider.Translate(source)

	if err == nil {
		log.Debug().Str("source", source).Str("result", translated).Msg("set cache")
		err = b.db.Update(func(txn *badger.Txn) error {
			return txn.Set([]byte(source), []byte(translated))
		})
	}

	return translated, err
}
