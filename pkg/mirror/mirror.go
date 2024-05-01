package mirror

import (
	"errors"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/jyny/mirror-action/pkg/logger"
)

const (
	MirrorRemoteName = "mirror"
)

type Mirror struct {
	srcRemote string
	srcAuth   transport.AuthMethod
	dstRemote string
	dstAuth   transport.AuthMethod
	refSpec   config.RefSpec
	logger    logger.Logger
}

func (m *Mirror) Run() error {
	m.logger.Info("Start mirroring...")
	defer m.logger.Info("End mirroring...")

	m.logger.Info("Cloning source repository...")
	srcRepo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:      m.srcRemote,
		Auth:     m.srcAuth,
		Progress: m.logger,
		Mirror:   true,
	})
	if err != nil {
		m.logger.Error("failed to clone source repository", "err", err)
		return err
	}

	_, err = srcRepo.CreateRemote(&config.RemoteConfig{
		Name: MirrorRemoteName,
		URLs: []string{m.dstRemote},
	})
	if err != nil {
		m.logger.Error("failed to create remote", "err", err)
		return err
	}

	m.logger.Info("Pushing to destination repository...")
	err = srcRepo.Push(&git.PushOptions{
		RemoteName: MirrorRemoteName,
		Auth:       m.dstAuth,
		Progress:   m.logger,
		RefSpecs:   []config.RefSpec{m.refSpec},
	})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		m.logger.Error("failed to push", "err", err)
		return err
	}

	m.logger.Debug("Successfully mirrored.")
	return nil
}
