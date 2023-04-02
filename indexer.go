package main

import (
	"os"

	"github.com/blugelabs/bluge/analysis"
	"github.com/konidev20/restic-index/blugeindex"
	"github.com/rs/zerolog"
)

var log = zerolog.New(os.Stderr).With().Timestamp().Logger()

type IndexOptions struct {
	Filter Filter

	// Batching improves indexing speed at the cost of using
	// some more memory. Dramatically improves indexing speed
	// for large number of files.
	// 0 or 1 disables batching. Defaults to 1 (no batching) if not set.
	BatchSize uint

	// If set to true, all the repository snapshots and files will be scanned and re-indexed.
	Reindex bool

	// DocumentBuilder is responsible of creating the Bluge document that will be indexed
	DocumentBuilder blugeindex.DocumentBuilder

	filenameAnalyzer *analysis.Analyzer
}

type Repo struct {
	Location string
	Password string
}

type Indexer struct {
}
