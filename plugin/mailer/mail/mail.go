package mailengine

import (
	"flag"
	"fmt"
	"github.com/leductoan3082004/go-sdk/appCommon"
	"github.com/leductoan3082004/go-sdk/logger"
	"net/smtp"
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
	flag.StringVar(&s.address, s.prefix+"-address", "", "address")
}

func (s *mailEngine) Configure() error {
	s.logger = logger.GetCurrent().GetLogger(s.prefix)
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

type Mail struct {
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	Receiver string `json:"receiver"`
}

func (s *mailEngine) SendMail(mail Mail) error {
	fmt.Println(s.address)
	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte("Subject:" + mail.Subject + "\n" + mime + mail.Body)

	err := smtp.SendMail(s.address, auth, s.username, []string{mail.Receiver}, msg)
	if err != nil {
		return appCommon.NewCustomError(err, "cannot send mail", "ErrCannotSendMail")
	}
	return nil
}
