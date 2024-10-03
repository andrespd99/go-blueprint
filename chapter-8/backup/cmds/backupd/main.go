package backupd

import (
	"flag"
	"log"
	"time"

	"github.com/andrespd99/go-backup"
)

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()
	var (
		interval = flag.Duration("interval", 10*time.Second, "interval between checks")
		archive  = flag.String("archive", "archive", "path to archive location")
		dbpath   = flag.String("db", "./db", "path to filedb database")
	)
	flag.Parse()

	m := &backup.Monitor{
		Destination: *archive,
		Archiver:    backup.ZIP,
		Paths:       make(map[string]string),
	}
}
