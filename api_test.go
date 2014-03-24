package main

import (
	"io"
	. "launchpad.net/gocheck"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
)

type ApiSuite struct{}

func (s *ApiSuite) getTestData(p ...string) io.ReadCloser {
	p = append([]string{}, ".", "testdata")
	fp := path.Join(p...)
	f, _ := os.OpenFile(fp, os.O_RDONLY, 0)
	return f
}

var _ = Suite(&ApiSuite{})

func (s *ApiSuite) TestCreateService(c *C) {
	b := s.getTestData("createservice.json")
	request, err := http.NewRequest("POST", "/services", b)
	c.Assert(err, IsNil)
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	err = AddInstance(recorder, request)
	c.Assert(err, IsNil)
	c.Assert(recorder.Code, Equals, 201)
}
