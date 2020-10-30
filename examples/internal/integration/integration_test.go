package integration_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/binchencoder/ease-gateway/examples/internal/proto/examplepb"
	"github.com/binchencoder/ease-gateway/gateway/runtime"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

var marshaler = &runtime.JSONPb{}

func TestEcho(t *testing.T) {
	if testing.Short() {
		t.Skip()
		return
	}

	for _, apiPrefix := range []string{"v1", "v2"} {
		t.Run(apiPrefix, func(t *testing.T) {
			testEcho(t, 8088, apiPrefix, "application/json")
			testEchoOneof(t, 8088, apiPrefix, "application/json")
			testEchoOneof1(t, 8088, apiPrefix, "application/json")
			testEchoOneof2(t, 8088, apiPrefix, "application/json")
			testEchoBody(t, 8088, apiPrefix)
			// Use SendHeader/SetTrailer without gRPC server https://github.com/grpc-ecosystem/grpc-gateway/issues/517#issuecomment-684625645
			testEchoBody(t, 8089, apiPrefix)
		})
	}
}

func testEcho(t *testing.T, port int, apiPrefix string, contentType string) {
	apiURL := fmt.Sprintf("http://localhost:%d/%s/example/echo/myid", port, apiPrefix)
	resp, err := http.Post(apiURL, "application/json", strings.NewReader("{}"))
	if err != nil {
		t.Errorf("http.Post(%q) failed with %v; want success", apiURL, err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if apiPrefix == "v2" {
		if got, want := resp.StatusCode, http.StatusNotFound; got != want {
			t.Errorf("resp.StatusCode = %d; want %d", got, want)
			t.Logf("%s", buf)
		}
	} else {
		if got, want := resp.StatusCode, http.StatusOK; got != want {
			t.Errorf("resp.StatusCode = %d; want %d", got, want)
			t.Logf("%s", buf)
		}
	}

	msg := new(examplepb.UnannotatedSimpleMessage)
	if err := marshaler.Unmarshal(buf, msg); err != nil {
		t.Errorf("marshaler.Unmarshal(%s, msg) failed with %v; want success", buf, err)
		return
	}
	if got, want := msg.Id, "myid"; got != want {
		t.Errorf("msg.Id = %q; want %q", got, want)
	}

	if value := resp.Header.Get("Content-Type"); value != contentType {
		t.Errorf("Content-Type was %s, wanted %s", value, contentType)
	}
}

func testEchoOneof(t *testing.T, port int, apiPrefix string, contentType string) {
	apiURL := fmt.Sprintf("http://localhost:%d/%s/example/echo/myid/10/golang", port, apiPrefix)
	resp, err := http.Get(apiURL)
	if err != nil {
		t.Errorf("http.Get(%q) failed with %v; want success", apiURL, err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if apiPrefix == "v2" {
		if got, want := resp.StatusCode, http.StatusNotFound; got != want {
			t.Errorf("resp.StatusCode = %d; want %d", got, want)
			t.Logf("%s", buf)
		}
	} else {
		if got, want := resp.StatusCode, http.StatusOK; got != want {
			t.Errorf("resp.StatusCode = %d; want %d", got, want)
			t.Logf("%s", buf)
		}
	}

	msg := new(examplepb.UnannotatedSimpleMessage)
	if err := marshaler.Unmarshal(buf, msg); err != nil {
		t.Errorf("marshaler.Unmarshal(%s, msg) failed with %v; want success", buf, err)
		return
	}
	if got, want := msg.GetLang(), "golang"; got != want {
		t.Errorf("msg.GetLang() = %q; want %q", got, want)
	}

	if value := resp.Header.Get("Content-Type"); value != contentType {
		t.Errorf("Content-Type was %s, wanted %s", value, contentType)
	}
}

func testEchoOneof1(t *testing.T, port int, apiPrefix string, contentType string) {
	apiURL := fmt.Sprintf("http://localhost:%d/%s/example/echo1/myid/10/golang", port, apiPrefix)
	resp, err := http.Get(apiURL)
	if err != nil {
		t.Errorf("http.Get(%q) failed with %v; want success", apiURL, err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if apiPrefix == "v2" {
		if got, want := resp.StatusCode, http.StatusNotFound; got != want {
			t.Errorf("resp.StatusCode = %d; want %d", got, want)
			t.Logf("%s", buf)
		}
	} else {
		if got, want := resp.StatusCode, http.StatusOK; got != want {
			t.Errorf("resp.StatusCode = %d; want %d", got, want)
			t.Logf("%s", buf)
		}
	}

	msg := new(examplepb.UnannotatedSimpleMessage)
	if err := marshaler.Unmarshal(buf, msg); err != nil {
		t.Errorf("marshaler.Unmarshal(%s, msg) failed with %v; want success", buf, err)
		return
	}
	if got, want := msg.GetStatus().GetNote(), "golang"; got != want {
		t.Errorf("msg.GetStatus().GetNote() = %q; want %q", got, want)
	}

	if value := resp.Header.Get("Content-Type"); value != contentType {
		t.Errorf("Content-Type was %s, wanted %s", value, contentType)
	}
}

func testEchoOneof2(t *testing.T, port int, apiPrefix string, contentType string) {
	apiURL := fmt.Sprintf("http://localhost:%d/%s/example/echo2/golang", port, apiPrefix)
	resp, err := http.Get(apiURL)
	if err != nil {
		t.Errorf("http.Get(%q) failed with %v; want success", apiURL, err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if apiPrefix == "v2" {
		if got, want := resp.StatusCode, http.StatusNotFound; got != want {
			t.Errorf("resp.StatusCode = %d; want %d", got, want)
			t.Logf("%s", buf)
		}
	} else {
		if got, want := resp.StatusCode, http.StatusOK; got != want {
			t.Errorf("resp.StatusCode = %d; want %d", got, want)
			t.Logf("%s", buf)
		}
	}

	msg := new(examplepb.UnannotatedSimpleMessage)
	if err := marshaler.Unmarshal(buf, msg); err != nil {
		t.Errorf("marshaler.Unmarshal(%s, msg) failed with %v; want success", buf, err)
		return
	}
	if got, want := msg.GetNo().GetNote(), "golang"; got != want {
		t.Errorf("msg.GetNo().GetNote() = %q; want %q", got, want)
	}

	if value := resp.Header.Get("Content-Type"); value != contentType {
		t.Errorf("Content-Type was %s, wanted %s", value, contentType)
	}
}

func testEchoBody(t *testing.T, port int, apiPrefix string) {
	sent := examplepb.UnannotatedSimpleMessage{Id: "example"}
	payload, err := marshaler.Marshal(&sent)
	if err != nil {
		t.Fatalf("marshaler.Marshal(%#v) failed with %v; want success", payload, err)
	}

	apiURL := fmt.Sprintf("http://localhost:%d/%s/example/echo_body", port, apiPrefix)
	resp, err := http.Post(apiURL, "", bytes.NewReader(payload))
	if err != nil {
		t.Errorf("http.Post(%q) failed with %v; want success", apiURL, err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if apiPrefix == "v2" {
		if got, want := resp.StatusCode, http.StatusNotFound; got != want {
			t.Errorf("resp.StatusCode = %d; want %d", got, want)
			t.Logf("%s", buf)
		}
	} else {
		if got, want := resp.StatusCode, http.StatusOK; got != want {
			t.Errorf("resp.StatusCode = %d; want %d", got, want)
			t.Logf("%s", buf)
		}
	}

	var received examplepb.UnannotatedSimpleMessage
	if err := marshaler.Unmarshal(buf, &received); err != nil {
		t.Errorf("marshaler.Unmarshal(%s, msg) failed with %v; want success", buf, err)
		return
	}
	if diff := cmp.Diff(&received, &sent, protocmp.Transform()); diff != "" {
		t.Errorf(diff)
	}

	if got, want := resp.Header.Get("Grpc-Metadata-Foo"), "foo1"; got != want {
		t.Errorf("Grpc-Metadata-Foo was %q, wanted %q", got, want)
	}
	if got, want := resp.Header.Get("Grpc-Metadata-Bar"), "bar1"; got != want {
		t.Errorf("Grpc-Metadata-Bar was %q, wanted %q", got, want)
	}

	if got, want := resp.Trailer.Get("Grpc-Trailer-Foo"), "foo2"; got != want {
		t.Errorf("Grpc-Trailer-Foo was %q, wanted %q", got, want)
	}
	if got, want := resp.Trailer.Get("Grpc-Trailer-Bar"), "bar2"; got != want {
		t.Errorf("Grpc-Trailer-Bar was %q, wanted %q", got, want)
	}
}
