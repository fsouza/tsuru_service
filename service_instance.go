package main

import (
	"encoding/json"
	"errors"
	"github.com/tsuru/config"
	"os/exec"
)

type Service struct {
	Name string
}

type App struct {
	Name, Host string
}

type ServiceProvider interface {
	Create() error
	BindApp(app *App) error
	UnbindApp(app *App) error
}

type Instance struct {
	Name   string
	Host   string
	Status string
}

func runCmd(cmd string, args ...string) ([]byte, error) {
	command := exec.Command(cmd, args...)
	return command.CombinedOutput()
}

func ip(instance_id []byte) (string, error) {
	docker, err := config.GetString("docker:binary")
	if err != nil {
		return "", err
	}
	instance_json, err := runCmd("sudo", docker, "inspect", string(instance_id))
	var result map[string]interface{}
	err = json.Unmarshal(instance_json, &result)
	if err != nil {
		return "", err
	}
	networkSettings := result["NetworkSettings"].(map[string]interface{})
	instance_ip := networkSettings["IpAddress"].(string)
	if instance_ip != "" {
		return instance_ip, nil
	}
	return "", errors.New("instance ip not found")
}

func (s *Instance) Create() error {
	docker, err := config.GetString("docker:binary")
	if err != nil {
		return err
	}

	image, err := config.GetString("docker:image")
	if err != nil {
		return err
	}

	output, err := runCmd("sudo", docker, "run", "-d", image, "/usr/sbin/sshd", "-D")
	if err != nil {
		return err
	}
	s.Host, err = ip(output)
	if err != nil {
		return nil
	}
	conn, err := Conn()
	if err != nil {
		return err
	}
	return conn.Instances().Insert(s)
}
