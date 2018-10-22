package session

import (
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"

	"gitlab.i/pkg/web/model"
)

func dialWithDB(network, address, password, DB string) (redis.Conn, error) {
	c, err := redis.Dial(network, address)
	if err != nil {
		return nil, err
	}
	if _, err := c.Do("SELECT", DB); err != nil {
		c.Close()
		return nil, err
	}
	return c, err
}

func NewPool(cfg *model.RedisStorage) *redis.Pool {
	dbAddr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	p := &redis.Pool{
		MaxIdle:     cfg.Size,
		IdleTimeout: 120 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return dialWithDB("tcp", dbAddr, "", cfg.Database)
		},
	}
	return p
}

func NewSentinelPool(cfg *model.RedisStorage) *redis.Pool {
	dbAddr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	sntnl := &sentinel.Sentinel{
		Addrs:      []string{dbAddr},
		MasterName: cfg.Master,
		Dial: func(addr string) (redis.Conn, error) {
			timeout := 500 * time.Millisecond
			c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   cfg.Size,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			masterAddr, err := sntnl.MasterAddr()
			if err != nil {
				return nil, err
			}
			c, err := dialWithDB("tcp", masterAddr, "", cfg.Database)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if !sentinel.TestRole(c, "master") {
				return errors.New("Role check failed")
			} else {
				return nil
			}
		},
	}
}
