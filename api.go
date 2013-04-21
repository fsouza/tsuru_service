package main

import (
        "encoding/json"
        "labix.org/v2/mgo/bson"
        "net/http"
)

// AddInstance add service instance
func AddInstance(w http.ResponseWriter, r *http.Request) error {
        defer r.Body.Close()
        name := r.FormValue("name")
        service := &Instance{Name: name}
        if err := service.Create(); err != nil {
                w.WriteHeader(http.StatusCreated)
                return nil
        }
        return nil
}

// Bind bind the service instance to app
func Bind(w http.ResponseWriter, r *http.Request) error {
        conn, err := Conn()
        if err != nil {
                return err
        }
        name := r.FormValue("name")
        var inst Instance
        err = conn.Instances().Find(bson.M{"name": name}).One(&inst)
        if err != nil {
                return err
        }
        ret := map[string]string{"HOST": inst.Host}
        out, err := json.Marshal(ret)
        if err != nil {
                return err
        }
        w.Write(out)
        return nil
}

func Unbind(w http.ResponseWriter, r *http.Request) error {
        return nil
}
