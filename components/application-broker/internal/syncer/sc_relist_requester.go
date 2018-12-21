package syncer

import (
	"time"

	"github.com/sirupsen/logrus"
)

const maxSyncRetries = 5

//go:generate mockery -name=brokerSyncer -output=automock -outpkg=automock -case=underscore
type brokerSyncer interface {
	Sync(retries int) error
}

// RelistRequester informs the Service Catalog that given Service Broker should be relisted.
// Due to performance reason, many relist requests which happen during the period defined in `reListDurationWindow`
// result in a single Service Catalog relist trigger.
type RelistRequester struct {
	brokerSyncer         brokerSyncer
	reListDurationWindow time.Duration
	log                  logrus.FieldLogger

	timeAfterProvider TimeAfterProvider

	relistRequested chan struct{}
}

// NewRelistRequester returns new instance of RelistRequester
func NewRelistRequester(brokerSyncer brokerSyncer, reListDuration time.Duration, log logrus.FieldLogger) *RelistRequester {
	return &RelistRequester{
		brokerSyncer:         brokerSyncer,
		reListDurationWindow: reListDuration,
		log:                  log.WithField("service", "syncer:sc-relist-requester"),

		relistRequested: make(chan struct{}, 1),
	}
}

// RequestRelist informs the Service Catalog that Broker should be relisted.
func (r *RelistRequester) RequestRelist() {
	select {
	case r.relistRequested <- struct{}{}:
	default: // relist already requested, drop next request
	}
}

// Run runs worker which executing relist operation
func (r *RelistRequester) Run(stopCh <-chan struct{}) {
	r.log.Infof("Starting Broker relist worker with relist duration window %v", r.reListDurationWindow)
	for {
		select {
		case <-r.timeAfterProvider.After(r.reListDurationWindow):
		case <-stopCh:
			r.log.Infof("Shutting down Broker relist worker")
			return
		}
		select {
		case <-r.relistRequested:
			if err := r.brokerSyncer.Sync(maxSyncRetries); err != nil {
				r.log.Errorf("Error occurred when synchronizing ServiceBrokers [maxRetires=%d]: %v", maxSyncRetries, err)
			} else {
				r.log.Infof("Relist request for ServiceBrokers fulfilled")
			}
		default:
		}
	}
}

// TimeAfterProvider is a provider for time.After.
// If not initialised defaults to time.After implementation for stdlib.
// It's intended to facilitate testing without time dependency.
type TimeAfterProvider func(d time.Duration) <-chan time.Time

// After calls attached After implementation.
func (p TimeAfterProvider) After(d time.Duration) <-chan time.Time {
	if p == nil {
		return time.After(d)
	}
	return p(d)
}
