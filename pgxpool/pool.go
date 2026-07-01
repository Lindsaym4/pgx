package pgxpool

import (
	"context"
)

func (p *Pool) Acquire(ctx context.Context) (*Conn, error) {
	conn, err := p.acquireRaw(ctx)
	if err != nil {
		return nil, err
	}

	if p.config.BeforeAcquire != nil {
		if !p.config.BeforeAcquire(ctx, conn) {
			conn.Close()
			return p.Acquire(ctx) // Retry acquisition
		}
	}

	return conn, nil
}