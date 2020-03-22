package common

type Config struct {
	BaseUri    string `mapstructure:"base_uri"`
	OutDir     string `mapstructure:"out_dir"`
	ApiTokenOf string `mapstructure:"api_token_of"`
	ApiTokenEm string `mapstructure:"api_token_em"`
	DocIdOf    string `mapstructure:"doc_id_of"`
	DocIdEm    string `mapstructure:"doc_id_em"`
}

func (c Config) validate() {
	if c.ApiTokenOf == "" || c.BaseUri == "" || c.DocIdOf == "" {
		panic("Config parameters are not set")
	}
}
