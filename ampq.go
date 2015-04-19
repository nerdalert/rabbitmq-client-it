package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
)

type AmqpHost struct {
	amqpAddr string
	uid      string
	pwd      string
	vhost    string
	queue    string
}

func (a *AmqpHost) GetAmqpAddr() string {
	return a.amqpAddr
}
func (a *AmqpHost) SetAmqpAddr(value string) {
	a.amqpAddr = value
}
func (a *AmqpHost) GetUid() string {
	return a.uid
}
func (a *AmqpHost) SetUid(value string) {
	a.uid = value
}
func (a *AmqpHost) GetPwd() string {
	return a.pwd
}
func (a *AmqpHost) SetPwd(value string) {
	a.pwd = value
}
func (a *AmqpHost) GetVhost() string {
	return a.vhost
}
func (a *AmqpHost) SetVhost(value string) {
	a.vhost = value
}
func (a *AmqpHost) GetQueue() string {
	return a.queue
}
func (a *AmqpHost) SetQueue(value string) {
	a.queue = value
}

func throwErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}