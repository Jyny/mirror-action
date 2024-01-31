package config

type Config struct {
	Debug bool `env:"DEBUG" envDefault:"false"`

	SrcRemoteURL     string `env:"src_remote_url,required"`
	SrcSShKey        string `env:"src_ssh_key"`
	SrcKnownHost     string `env:"src_known_host"`
	SrcIgnoreHostKey bool   `env:"src_ignore_host_key"`
	SrcUsername      string `env:"src_user_name"`
	SrcPassword      string `env:"src_password"`

	DstRemoteURL     string `env:"dst_remote_url,required"`
	DstSShKey        string `env:"dst_ssh_key"`
	DstKnownHost     string `env:"dst_known_host"`
	DstIgnoreHostKey bool   `env:"dst_ignore_host_key"`
	DstUsername      string `env:"dst_user_name"`
	DstPassword      string `env:"dst_password"`
}
