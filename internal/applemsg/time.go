package applemsg

import "time"

// ParseAppleTime converts Apple's Messages timestamp
// into Go time.
//
// Current macOS versions usually store nanoseconds since Unix epoch.
// This function keeps the conversion isolated so future format changes
// only affect this file.
func ParseAppleTime(v int64) time.Time {

	if v == 0 {
		return time.Time{}
	}

	// Current Messages format:
	// nanoseconds since Unix epoch
	if v > 100000000000000000 {
		return time.Unix(
			v/1_000_000_000,
			v%1_000_000_000,
		)
	}

	// microseconds fallback
	if v > 100000000000000 {
		return time.UnixMicro(v)
	}

	// seconds fallback
	return time.Unix(v, 0)
}
