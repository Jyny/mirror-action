package config

type Config struct {
	Debug   bool   `env:"DEBUG" envDefault:"false"`
	RefSpec string `env:"REF_SPEC" envDefault:"+refs/heads/*:refs/heads/*"`

	SrcRemoteURL     string `env:"INPUT_SRC_REMOTE_URL,required"`
	SrcSShKey        string `env:"INPUT_SRC_SSH_KEY"`
	SrcKnownHost     string `env:"INPUT_SRC_KNOWN_HOST"`
	SrcIgnoreHostKey bool   `env:"INPUT_SRC_IGNORE_HOST_KEY"`
	SrcUsername      string `env:"INPUT_SRC_USERNAME"`
	SrcPassword      string `env:"INPUT_SRC_PASSWORD"`

	DstRemoteURL     string `env:"INPUT_DST_REMOTE_URL,required"`
	DstSShKey        string `env:"INPUT_DST_SSH_KEY"`
	DstKnownHost     string `env:"INPUT_DST_KNOWN_HOST"`
	DstIgnoreHostKey bool   `env:"INPUT_DST_IGNORE_HOST_KEY"`
	DstUsername      string `env:"INPUT_DST_USERNAME"`
	DstPassword      string `env:"INPUT_DST_PASSWORD"`
}
