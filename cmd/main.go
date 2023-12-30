package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/blugelabs/bluge"
	"github.com/dustin/go-humanize"
	"github.com/konidev20/rindex"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	indexer, err := rindex.New("test/tmp", "/Users/srigovindnayak/test-repo", "root")
	if err != nil {
		fmt.Printf("could not initialize the indexer: %v", err)
	}

	statsChannel := make(chan rindex.IndexStats)

	stats, err := indexer.Index(ctx, rindex.DefaultIndexOptions, statsChannel)
	if err != nil {
		fmt.Printf("could not index repository : %v", err)
	}

	fmt.Printf("%v", stats)

	// indexer, err := rindex.NewOffline("test/tmp", "D:/test", "test123")
	// if err != nil {
	// 	fmt.Printf("could not initialize the indexer: %v", err)
	// }

	// fVisitor := func(field string, value []byte) bool {
	// 	printMetadata(field, value)
	// 	return true
	// }

	// srVisitor := func() bool {
	// 	fmt.Println("=====================================")
	// 	return true
	// }

	// query := `filename:/.*\.(zip|ttf)/`

	// count, err := indexer.Search(query, fVisitor, srVisitor)
	// if err != nil {
	// 	fmt.Printf("could not search the index: %v", err)
	// }

	// fmt.Println("number of matches: ", count)
}

func printMetadata(field string, value []byte) {
	var f string
	if field == "_id" {
		f = "ID"
	} else if field == "repository_id" {
		f = "Repository ID"
	} else {
		f = strings.Title(strings.ReplaceAll(field, "_", " "))
	}

	v := ""
	switch field {
	case "mtime":
		t, err := bluge.DecodeDateTime(value)
		if err != nil {
			v = "error"
		} else {
			v = t.Format("2006-1-2")
		}
	case "updated":
		t, err := bluge.DecodeDateTime(value)
		if err != nil {
			v = "error"
		} else {
			v = t.Format("2006-1-2")
		}
	case "size":
		t, err := bluge.DecodeNumericFloat64(value)
		if err != nil {
			v = "error"
		} else {
			v = humanize.Bytes(uint64(t))
		}
	case "blobs":
		v = ""
	default:
		v = string(value)
	}

	if v != "" && field == "filename" || field == "mtime" || field == "path" {
		printRow(f, v)
	}
}

func printRow(header, value string) {
	fmt.Printf("%s %s\n", header+":", value)
}

func booleanQuery() *bluge.BooleanQuery {
	bq := bluge.NewBooleanQuery()
	bq.SetMinShould(1)
	return bq
}
