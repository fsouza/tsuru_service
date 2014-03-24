package main

import (
	"labix.org/v2/mgo/bson"
	. "launchpad.net/gocheck"
)

func (s *S) TestCreateInstance(c *C) {
	instanceName := "postgres-instance1"
	instance := &Instance{Name: instanceName}
	err := instance.Create()
	c.Assert(err, IsNil)
	c.Assert(instance.Host, NotNil)
	var result Instance
	coll := s.conn.Instances()
	err = coll.Find(bson.M{"name": instanceName}).One(&result)
	c.Assert(err, IsNil)
	c.Assert(result.Host, Not(Equals), "")
	c.Assert(result.Name, Equals, instanceName)
}
