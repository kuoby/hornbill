package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

type process struct {
	cli *client.Client
	id  string
}

func (p *process) Running(ctx context.Context) (bool, error) {
	c, err := p.cli.ContainerInspect(ctx, p.id)
	if err != nil {
		return false, err
	}
	return c.State.Running, nil
}

func (p *process) Exists(ctx context.Context) (bool, error) {
	if _, err := p.cli.ContainerInspect(ctx, p.id); err != nil {
		if client.IsErrNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (p *process) Configure(ctx context.Context) error {
	_, err := p.cli.ContainerInspect(ctx, p.id)
	if err == nil {
		return nil
	}
	if !client.IsErrNotFound(err) {
		return errors.Wrap(err, "could not inspect container")
	}

	cfg := &container.Config{}
	hostCfg := &container.HostConfig{}
	if _, err := p.cli.ContainerCreate(ctx, cfg, hostCfg, nil, nil, p.id); err != nil {
		errors.Wrap(err, "could not create container")
	}
	return nil
}

func (p *process) Start(ctx context.Context) error {
	return nil
}
