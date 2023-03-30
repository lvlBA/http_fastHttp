package app

type Config struct {
	ListenAddress  string `json:"listen_address"   yaml:"listen_address"   env:"LISTEN_ADDRESS"  envdefault:"8080"`
	ListenAddress2 string `json:"listen_address2" yaml:"listen_address2" env:"LISTEN_ADDRESS2" envdefault:"80"`
}
