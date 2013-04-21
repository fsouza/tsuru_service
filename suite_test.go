package main

import (
        . "launchpad.net/gocheck"
        "testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct {
        conn *Storage
}

var _ = Suite(&S{})

func (s *S) SetUpSuite(c *C) {
        s.conn, _ = Conn()
}
