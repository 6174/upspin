// Package gcptest implements simple types and utility functions to help test users of GCP client.
package gcptest

import (
	"errors"

	"upspin.googlesource.com/upspin.git/cloud/gcp"
)

// DummyGCP is a dummy version of gcp.Interface that does nothing.
type DummyGCP struct {
}

var _ gcp.Interface = (*DummyGCP)(nil)

func (m *DummyGCP) PutLocalFile(srcLocalFilename string, ref string) (refLink string, error error) {
	return "", nil
}

func (m *DummyGCP) Get(ref string) (link string, error error) {
	return "", nil
}

func (m *DummyGCP) Download(ref string) ([]byte, error) {
	return nil, nil
}

func (m *DummyGCP) Put(ref string, contents []byte) (refLink string, error error) {
	return "", nil
}

func (m *DummyGCP) List(prefix string) (name []string, link []string, err error) {
	return []string{}, []string{}, nil
}

func (m *DummyGCP) Delete(ref string) error {
	return nil
}

func (m *DummyGCP) Connect() {
}

// ExpectGetGCP is a DummyGCP that expects Get will be called with a
// given ref and when it does, it replies with the preset link.
type ExpectGetGCP struct {
	DummyGCP
	Ref  string
	Link string
}

func (e *ExpectGetGCP) Get(ref string) (link string, error error) {
	if ref == e.Ref {
		return e.Link, nil
	}
	return "", errors.New("not found")
}

// ExpectDownloadCapturePutGCP inspects all calls to Download with the
// given Ref and if it matches, it returns Data. It also captures all
// Put requests.
type ExpectDownloadCapturePutGCP struct {
	DummyGCP
	// Expectation for calls to Download
	Ref  string
	Data []byte
	// Storage for calls to Put
	PutRef      []string
	PutContents [][]byte
}

func (e *ExpectDownloadCapturePutGCP) Download(ref string) ([]byte, error) {
	if ref == e.Ref {
		return e.Data, nil
	}
	return nil, errors.New("not found")
}

func (c *ExpectDownloadCapturePutGCP) Put(ref string, contents []byte) (refLink string, error error) {
	c.PutRef = append(c.PutRef, ref)
	c.PutContents = append(c.PutContents, contents)
	return "", nil
}