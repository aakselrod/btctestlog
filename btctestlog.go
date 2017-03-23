// This file was originally written by Alex Akselrod. See UNLICENSE for info.

package btctestlog

import (
	"testing"

	"github.com/btcsuite/seelog"
)

// testLog is a custom receiver that writes messages to an instance of
// testing.T to allow tests using btclog to write logging output during tests
// for debugging.
type testLog struct {
	t *testing.T
}

// ReceiveMessage is required by the seelog.CustomReceiver interface.
func (l *testLog) ReceiveMessage(message string, level seelog.LogLevel,
	context seelog.LogContextInterface) error {
	l.t.Logf("[%s]: %s", level.String(), message)
	return nil
}

// AfterParse is required by the seelog.CustomReceiver interface.
func (*testLog) AfterParse(seelog.CustomReceiverInitArgs) error { return nil }

// Flush is required by the seelog.CustomReceiver interface.
func (*testLog) Flush() {}

// Close is required by the seelog.CustomReceiver interface.
func (*testLog) Close() error { return nil }

// NewTestLogger creates a seelog.LoggerInterface based on the above custom
// receiver.
func NewTestLogger(t *testing.T) (seelog.LoggerInterface, error) {
	return seelog.LoggerFromCustomReceiver(&testLog{t: t})
}
