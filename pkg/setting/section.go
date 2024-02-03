package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MySQLSettingS struct {
	IP        string
	Port      string
	User      string
	Password  string
	Database  string
	Charset   string
	ParseTime bool
}

type SMTPSettingS struct {
	Host     string
	Port     int
	User     string
	Password string
}

type RedisSettingS struct {
	IP   string
	Port string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
