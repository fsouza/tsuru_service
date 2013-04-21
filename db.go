package main

import (
        "github.com/globocom/config"
        "labix.org/v2/mgo"
)

type Storage struct {
        session *mgo.Session
        dbname string
}

func (s *Storage) Instances() *mgo.Collection {
        c := s.Collection("instances")
        return c
}

func (s *Storage) Collection(name string) *mgo.Collection {
        return s.session.DB(s.dbname).C(name)
}

func Conn() (*Storage, error) {
        url, _ := config.GetString("database:url")
        dbname, _ := config.GetString("database:name")
        return Open(url, dbname)
}

func Open(addr string, name string) (*Storage, error) {
        sess, err := mgo.DialWithTimeout(addr, 1e9)
        if err != nil {
                return nil, err
        }

        sessCopy := sess.Copy()
        storage := &Storage{session: sessCopy, dbname: name}
        return storage, nil
}
