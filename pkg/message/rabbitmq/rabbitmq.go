package rabbitmq

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"strings"
	"sync"
)

type Client struct {
	con  *amqp.Connection
	chn  []*amqp.Channel
	lock sync.Mutex

	// connection info
	host     string
	port     uint16
	username string
	password string
	vhost    string
}

func CreateRabbitmqClient(host string, port uint16, username, password, vhost string) (*Client, error) {

	c := &Client{
		host:     host,
		port:     port,
		username: username,
		password: password,
		vhost:    vhost,
	}
	if err := c.connect(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) connect() error {
	v := c.vhost
	if strings.HasPrefix(v, "/") {
		v = strings.TrimPrefix(v, "/")
	}
	conStr := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", c.username, c.password, c.host, c.port, v)
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		cond := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", c.username, "*******", c.host, c.port, v)
		logrus.Debugf("Rabbitmq的链接信息为: %s", cond)
	}
	con, err := amqp.Dial(conStr)
	if err != nil {
		return err
	}
	c.con = con
	return nil
}

func (c *Client) GetChannel() (chn *amqp.Channel, err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if !c.IsConnected() {
		if err := c.connect(); err != nil {
			return nil, err
		}
	}
	chn, err = c.con.Channel()
	if err != nil {
		return
	}
	c.chn = append(c.chn, chn)
	return
}

func (c *Client) Close() {
	c.con.Close()
	for _, c := range c.chn {
		c.Close()
	}
}

func (c *Client) IsConnected() bool {
	return !c.con.IsClosed()
}
