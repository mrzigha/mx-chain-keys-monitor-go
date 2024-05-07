package executors

import "errors"

var (
	errNilRatingsChecker             = errors.New("nil ratings checker instance")
	errNilOutputNotifier             = errors.New("nil output notifier")
	errNilValidatorStatisticsQuerier = errors.New("nil output notifier")
	errNilStatusHandler              = errors.New("nil status handler")
	errNilBLSKeysFetcher             = errors.New("nil BLS Keys fetcher")
)
