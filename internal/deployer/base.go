package deployer

import (
	"context"
	cm "sail/common"
)

type DeployersItf interface {
	SetInChan(inchan <-chan cm.IMessage)
	SetOutChan(outchan chan<- cm.IMessage)

	Install(context.Context, cm.IMessage) error
	Uninstall(cm.IMessage) error

	Run(context.Context)
}

type BaseDeployer struct {
	inChan  <-chan cm.IMessage
	outChan chan<- cm.IMessage
}

func (b *BaseDeployer) SetInChan(inchan <-chan cm.IMessage) {
	b.inChan = inchan
}

func (b *BaseDeployer) SetOutChan(outchan chan<- cm.IMessage) {
	b.outChan = outchan
}
