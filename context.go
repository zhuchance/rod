package rod

import (
	"context"
	"time"
)

// Context creates a clone with a context that inherits the previous one
func (b *Browser) Context(ctx context.Context) *Browser {
	if ctx == b.ctx {
		return b
	}

	newObj := *b
	newObj.ctx = ctx
	return &newObj
}

// GetContext returns the current context
func (b *Browser) GetContext() context.Context {
	return b.ctx
}

// Timeout for chained sub-operations
func (b *Browser) Timeout(d time.Duration) *Browser {
	ctx, cancel := context.WithTimeout(b.ctx, d)
	b.timeoutCancel = cancel
	return b.Context(ctx)
}

// CancelTimeout context
func (b *Browser) CancelTimeout() *Browser {
	b.timeoutCancel()
	return b
}

// Context creates a clone with a context that inherits the previous one
func (p *Page) Context(ctx context.Context) *Page {
	if ctx == p.ctx {
		return p
	}

	newObj := *p
	newObj.ctx = ctx
	return &newObj
}

// GetContext returns the current context
func (p *Page) GetContext() context.Context {
	return p.ctx
}

// Timeout for chained sub-operations
func (p *Page) Timeout(d time.Duration) *Page {
	ctx, cancel := context.WithTimeout(p.ctx, d)
	p.timeoutCancel = cancel
	return p.Context(ctx)
}

// CancelTimeout context
func (p *Page) CancelTimeout() *Page {
	p.timeoutCancel()
	return p
}

// Context creates a clone with a context that inherits the previous one
func (el *Element) Context(ctx context.Context) *Element {
	if ctx == el.ctx {
		return el
	}

	newObj := *el
	newObj.ctx = ctx
	return &newObj
}

// GetContext returns the current context
func (el *Element) GetContext() context.Context {
	return el.ctx
}

// Timeout for chained sub-operations
func (el *Element) Timeout(d time.Duration) *Element {
	ctx, cancel := context.WithTimeout(el.ctx, d)
	el.timeoutCancel = cancel
	return el.Context(ctx)
}

// CancelTimeout context
func (el *Element) CancelTimeout() *Element {
	el.timeoutCancel()
	return el
}
