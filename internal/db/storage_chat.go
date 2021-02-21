package db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"oea-go/internal/common/config"
	"time"
)

const (
	connectTimeout = time.Second * 10
	rwTimeout      = time.Second * 5
)

type Storage struct {
	addr string
}

func NewStorage(addr string) *Storage {
	return &Storage{addr: addr}
}

func (s *Storage) conn() (redis.Conn, error) {
	return redis.Dial(
		"tcp",
		s.addr,
		redis.DialConnectTimeout(connectTimeout),
		redis.DialReadTimeout(rwTimeout),
		redis.DialWriteTimeout(rwTimeout),
		redis.DialClientName("oea"),
	)
}

func (s *Storage) getChatIdKey(account config.Email) string {
	return fmt.Sprintf("chatid:%v", account)
}

func (s *Storage) SetChatID(account config.Email, chatID int64) (err error) {
	rconn, err := s.conn()

	if err != nil {
		return err
	}
	defer rconn.Close()

	_, err = rconn.Do("SET", s.getChatIdKey(account), chatID)

	return err
}

func (s *Storage) GetChatID(account config.Email) (chatID int64, err error) {
	rconn, err := s.conn()

	if err != nil {
		return 0, err
	}

	defer rconn.Close()

	chatID, err = redis.Int64(rconn.Do("GET", s.getChatIdKey(account)))

	if err != nil {
		return 0, err
	}

	return chatID, nil
}

func (s *Storage) Keys() (keys []string, err error) {
	rconn, err := s.conn()

	if err != nil {
		return nil, err
	}

	defer rconn.Close()

	keys, err = redis.Strings(rconn.Do("KEYS", "*"))

	if err != nil {
		return nil, err
	}

	return keys, nil
}
