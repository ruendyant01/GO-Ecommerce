package database

import (
	"errors"
	"fmt"
	"go_config/utils"
	"time"
)

type Options struct {
	Host             string
	Port             int
	Username         string
	Password         string
	DatabaseName     string
	Protocol         string
	Param            string
	ConnMaxIdleConns int
	ConnMaxLifetime  time.Duration
	ConnMaxOpenConns int
}

type Database interface {
	Open(options Options)
	Close() error
	Get() interface{}
	Ping() error
}

func BuildDNS(options Options) (string, error) {
	handleError := func(msg string) (string, error) { return "", errors.New(msg) }
	if utils.IsBlank(options.Host) {
		return handleError("host is blank")
	}
	if utils.IsBlank(options.DatabaseName) {
		return handleError("database name is blank")
	}
	if utils.IsBlank(options.Username) {
		return handleError("username is blank")
	}
	if utils.IsBlank(options.Password) {
		return handleError("password is blank")
	}
	if options.Port < 0 {
		return handleError("port is blank")
	}

	if utils.IsBlank(options.Protocol) {
		options.Protocol = "tcp"
	}

	if utils.IsBlank(options.Param) {
		options.Param = "parseTime=true"
	}

	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s", options.Username, options.Password, options.Protocol, options.Host, options.Port, options.DatabaseName, options.Param), nil
}
