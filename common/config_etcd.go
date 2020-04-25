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

func FillConfigFromEtcd(cfg *Config, etcd *Etcd) error {
	connErr := etcd.Connect()
	if connErr != nil {
		return connErr
	}
	defer etcd.Close()

	log.Printf("Waiting for etcd at %v for %v seconds...\n", etcd.Client.Endpoints(), dialTimeout)

	membersList, membersErr := etcd.MemberList()
	if membersErr != nil {
		return membersErr
	}
	log.Printf("Connected to etcd: %+v\n", membersList.Members)

	cfg.ForeachKey(func(name string, val interface{}, tags ConfigTags) (interface{}, bool) {
		keyResp, keyErr := etcd.GetNotEmpty(tags.Name)
		if keyErr != nil {
			log.Printf("config: (%v) %v: %v\n", boolStr(tags.IsMandatory, "mandatory", "optional"), tags.Name, keyErr)
			return nil, !tags.IsMandatory
		}

		newVal := string(keyResp.Kvs[0].Value)

		return cfg.CastStringValueToConfigType(name, newVal), true
	})

	return cfg.Validate()
}
