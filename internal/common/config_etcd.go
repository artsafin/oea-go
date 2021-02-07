package common

import (
	"log"
)

func boolStr(val bool, trueVal, falseVal string) string {
	if val {
		return trueVal
	}
	return falseVal
}

func FillConfigFromEtcd(cfg *Config, etcd *EtcdService) error {
	conn, connErr := etcd.ConnectAndPing()
	if connErr != nil {
		return connErr
	}
	defer conn.Close()

	cfg.ForeachKey(func(name string, val interface{}, tags ConfigTags) (interface{}, bool) {
		keyResp, keyErr := conn.GetNotEmpty(tags.Name)
		if keyErr != nil {
			log.Printf("config: (%v) %v: %v\n", boolStr(tags.IsMandatory, "mandatory", "optional"), tags.Name, keyErr)
			return nil, !tags.IsMandatory
		}

		newVal := string(keyResp.Kvs[0].Value)

		return cfg.CastStringValueToConfigType(name, newVal), true
	})

	return cfg.PrepareAndValidate()
}
