package mailer

import (
	"flag"
	"github.com/leductoan3082004/go-sdk/logger"
)

type mailEngine struct {
	prefix   string
	name     string
	logger   logger.Logger
	host     string
	port     string
	address  string
	username string
	password string
}

func NewMailEngine(name, prefix string) *mailEngine {
	return &mailEngine{
		logger: logger.GetCurrent().GetLogger(prefix),
		prefix: prefix,
		name:   name,
	}
}

func (s *mailEngine) GetPrefix() string {
	return s.prefix
}

func (s *mailEngine) Get() interface{} {
	return s
}

func (s *mailEngine) Name() string {
	return s.name
}

func (s *mailEngine) InitFlags() {
	flag.StringVar(&s.host, s.prefix+"-host", "", "host for mail engine")
	flag.StringVar(&s.port, s.prefix+"-port", "", "port for mail engine")
	flag.StringVar(&s.username, s.prefix+"-username", "", "username or email of account")
	flag.StringVar(&s.password, s.prefix+"-password", "", "app password for this account")
}

func (s *mailEngine) Configure() error {
	return nil
}

func (s *mailEngine) Run() error {
	return s.Configure()
}

func (s *mailEngine) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
