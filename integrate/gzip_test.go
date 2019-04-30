package integrate

import (
	"bytes"
	"testing"
)

func TestCompress(t *testing.T) {
	s := bytes.NewBufferString(`"in order to test gzip compress.in order to test gzip compress.
	                             in order to test gzip compress.in order to test gzip compress.
	                             in order to test gzip compress.in order to test gzip compress.
	                             in order to test gzip compress.in order to test gzip compress."`)
	var b bytes.Buffer
	t.Log(b.Len())
	size, err := Compress(&b, s)
	if err != nil {
		t.Error(err)
	}
	if size > 0 {
		t.Log(size, ",", s.Len(), ",", b.Len())
	} else {
		t.Error("compress error")
	}

}
