package deployers

import (
	"context"
	cm "sail/common"
)

type DeployersItf interface {
	SetInChan(inchan <-chan cm.Message)
	SetOutChan(outchan chan<- cm.Message)

	Install(context.Context, cm.Message) error
	Uninstall(cm.Message) error

	Run(context.Context)
}

type BaseDeployer struct {
	InChan  <-chan cm.Message
	OutChan chan<- cm.Message
}

func (b *BaseDeployer) SetInChan(inchan <-chan cm.Message) {
	b.InChan = inchan
}

func (b *BaseDeployer) SetOutChan(outchan chan<- cm.Message) {
	b.OutChan = outchan
}
