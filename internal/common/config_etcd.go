package common

import (
	"go.uber.org/zap"
)

func boolStr(val bool, trueVal, falseVal string) string {
	if val {
		return trueVal
	}
	return falseVal
}

func FillConfigFromEtcd(cfg *Config, etcd *EtcdService, logger *zap.SugaredLogger) error {
	conn, connErr := etcd.ConnectAndPing()
	if connErr != nil {
		return connErr
	}
	defer conn.Close()

	cfg.ForeachKey(func(name string, val interface{}, tags ConfigTags) (interface{}, bool) {
		keyResp, keyErr := conn.GetNotEmpty(tags.Name)
		if keyErr != nil {
			logger.Errorw("config field error", "mandatory", boolStr(tags.IsMandatory, "mandatory", "optional"), "tag", tags.Name, "error", keyErr)
			return nil, !tags.IsMandatory
		}

		newVal := string(keyResp.Kvs[0].Value)

		return cfg.CastStringValueToConfigType(name, newVal), true
	})

	return cfg.PrepareAndValidate()
}
