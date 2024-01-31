package mirror

import (
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/jyny/mirror-action/pkg/logger"
	xssh "golang.org/x/crypto/ssh"
)

const (
	rexGitRemoteURL = `^(([A-Za-z0-9]+@|http(|s)\:\/\/)|(http(|s)\:\/\/[A-Za-z0-9]+@))([A-Za-z0-9.]+(:\d+)?)(?::|\/)([\d\/\w.-]+?)(\.git){1}$`
)

type MirrorConfig struct {
	// RemoteURL is the URL of the repository to mirror.
	// It must be a valid git URL.
	// e.g. git@github.com:foo/bar.git
	// e.g. https://github.com/foo/bar.git
	RemoteURL string

	SSHKey        string
	HostKey       string
	IgnoreHostKey bool

	Username string
	Password string
}

func New(src, dst MirrorConfig, logger logger.Logger) (*Mirror, error) {
	logger.Debug("Source:", "remote url", src.RemoteURL)
	if !validGitURL(src.RemoteURL) {
		logger.Error("validate source remote url", "err", ErrInvalidGitRemoteURL)
		return nil, ErrInvalidGitRemoteURL
	}

	logger.Debug("Destination:", "remote url", dst.RemoteURL)
	if !validGitURL(dst.RemoteURL) {
		logger.Error("validate destination remote url", "err", ErrInvalidGitRemoteURL)
		return nil, ErrInvalidGitRemoteURL
	}

	return &Mirror{
		srcRemote: src.RemoteURL,
		srcAuth:   newAuthMethodOrNil(&src),
		dstRemote: dst.RemoteURL,
		dstAuth:   newAuthMethodOrNil(&dst),
		logger:    logger,
	}, nil
}

func validGitURL(url string) bool {
	return regexp.MustCompile(rexGitRemoteURL).MatchString(url)
}

func newAuthMethodOrNil(cfg *MirrorConfig) transport.AuthMethod {
	if cfg == nil {
		return nil
	}

	switch {
	case strings.HasPrefix(cfg.RemoteURL, "git@"):
		return newSSHAuthMethod(cfg)
	case strings.HasPrefix(cfg.RemoteURL, "https://"):
		return newHTTPAuthMethod(cfg)
	default:
		return nil
	}
}

func newSSHAuthMethod(cfg *MirrorConfig) ssh.AuthMethod {
	pk, err := ssh.NewPublicKeys("git", []byte(cfg.SSHKey), "")
	if err != nil {
		return nil
	}

	if cfg.IgnoreHostKey {
		pk.HostKeyCallback = xssh.InsecureIgnoreHostKey()
		return pk
	}

	hostKey, _, _, _, err := xssh.ParseAuthorizedKey([]byte(cfg.HostKey))
	if err != nil {
		return nil
	}
	pk.HostKeyCallback = xssh.FixedHostKey(hostKey)

	return pk
}

func newHTTPAuthMethod(cfg *MirrorConfig) http.AuthMethod {
	return &http.BasicAuth{
		Username: cfg.Username,
		Password: cfg.Password,
	}
}
