package http

import "time"

type LastModifiedChecker interface {
	CheckLastModified(lastModified time.Time) bool
}
