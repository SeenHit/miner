package memtxpool

import (
	"github.com/hacash/core/interfacev2"
)

func (p *MemTxPool) SubscribeOnAddTxSuccess(addtxCh chan interfacev2.Transaction) {
	p.addTxSuccess.Subscribe(addtxCh)
}

// 暂停事件订阅
func (p *MemTxPool) PauseEventSubscribe() {
	p.changeLock.Lock()
	defer p.changeLock.Unlock()

	p.isBanEventSubscribe = true
}

// 重开事件订阅
func (p *MemTxPool) RenewalEventSubscribe() {
	p.changeLock.Lock()
	defer p.changeLock.Unlock()

	p.isBanEventSubscribe = false
}
