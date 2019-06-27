package http

type ETagChecker interface {
	CheckETag(etag string) bool
}