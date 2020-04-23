// +build amd64,linux

package webdav

import (
	"context"
	"net/http"
	"os"
	"syscall"
	"time"
)

func findCreationDate(ctx context.Context, fs FileSystem, ls LockSystem, name string, fi os.FileInfo) (string, error) {
	stat := fi.Sys().(*syscall.Stat_t)
	return time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec).UTC().Format(http.TimeFormat), nil
}
