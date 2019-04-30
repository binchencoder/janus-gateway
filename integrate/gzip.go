package integrate

import (
	"bytes"
	"errors"
	"flag"
	"net/http"
	"strconv"
	"strings"

	"github.com/klauspost/compress/gzip"
)

var (
	gzipSize = flag.Int("gzip-size", 1, "If response data size > gzipSize*1024, response data will be compressed.")
)

// ResponseRecorder is an implementation of http.ResponseWriter.
type ResponseRecorder struct {
	Code        int           // the HTTP response code from WriteHeader
	HeaderMap   http.Header   // the HTTP response headers
	Body        *bytes.Buffer // if non-nil, the bytes.Buffer to append written data to
	Flushed     bool
	wroteHeader bool
}

// NewRecorder returns an initialized ResponseRecorder.
func NewRecorder() *ResponseRecorder {
	return &ResponseRecorder{
		HeaderMap: make(http.Header),
		Body:      new(bytes.Buffer),
		Code:      200,
	}
}

// Header returns the response headers.
func (rw *ResponseRecorder) Header() http.Header {
	m := rw.HeaderMap
	if m == nil {
		m = make(http.Header)
		rw.HeaderMap = m
	}
	return m
}

// writeHeader writes a header if it was not written yet and
// detects Content-Type if needed.
//
// bytes or str are the beginning of the response body.
// We pass both to avoid unnecessarily generate garbage
// in rw.WriteString which was created for performance reasons.
// Non-nil bytes win.
func (rw *ResponseRecorder) writeHeader(b []byte, str string) {
	if rw.wroteHeader {
		return
	}
	if len(str) > 512 {
		str = str[:512]
	}

	_, hasType := rw.HeaderMap["Content-Type"]
	hasTE := rw.HeaderMap.Get("Transfer-Encoding") != ""
	if !hasType && !hasTE {
		if b == nil {
			b = []byte(str)
		}
		rw.HeaderMap.Set("Content-Type", http.DetectContentType(b))
	}

	rw.WriteHeader(200)
}

// Write always succeeds and writes to rw.Body, because of bytes.Buffer.Write's
// error is always nil .
func (rw *ResponseRecorder) Write(buf []byte) (int, error) {
	rw.writeHeader(buf, "")
	return rw.Body.Write(buf)
}

// WriteString always succeeds and writes to rw.Body, because of
// bytes.Buffer.WriteString's error is always nil .
func (rw *ResponseRecorder) WriteString(str string) (int, error) {
	rw.writeHeader(nil, str)
	return rw.Body.WriteString(str)
}

// WriteHeader sets rw.Code.
func (rw *ResponseRecorder) WriteHeader(code int) {
	if !rw.wroteHeader {
		rw.Code = code
		rw.wroteHeader = true
	}
}

// Flush sets rw.Flushed to true.
func (rw *ResponseRecorder) Flush() {
	if !rw.wroteHeader {
		rw.WriteHeader(200)
	}
	rw.Flushed = true
}

// compress returns bytes len that param s after the compression.
func Compress(b *bytes.Buffer, s *bytes.Buffer) (int, error) {
	if s.Len() == 0 {
		return 0, errors.New("compress data size is zero")
	}
	g := gzip.NewWriter(b)
	if g == nil {
		return 0, errors.New("gzip compress error")
	}
	defer g.Close()
	return g.Write(s.Bytes())
}

// gzip 处理中间件
type GzipMiddleware struct {
	Handler http.Handler
}

func (m *GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := NewRecorder()
	// passing a ResponseRecorder instead of the original RW
	m.Handler.ServeHTTP(rec, r)
	// copy the original headers first
	for k, v := range rec.Header() {
		w.Header()[k] = v
	}
	// client support gizp and response data size>gzipSize*1024
	if ae := r.Header.Get("Accept-Encoding"); ae != "" && strings.Contains(ae, "gzip") && rec.Body.Len() > *gzipSize*1024 {
		var b bytes.Buffer
		if s, err := Compress(&b, rec.Body); err == nil {
			w.Header().Add("Content-Encoding", "gzip")
			r.Header.Set("Content-Length", strconv.Itoa(s))
			w.WriteHeader(rec.Code)
			// write out our data
			w.Write(b.Bytes())
			return
		}
	}
	// write out the original body
	w.WriteHeader(rec.Code)
	w.Write(rec.Body.Bytes())
}
