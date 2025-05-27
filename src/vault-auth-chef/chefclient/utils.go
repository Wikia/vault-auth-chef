package chefclient

import (
	"fmt"
	"time"

	"github.com/hashicorp/vault/sdk/logical"
)

// SanitizeTTLStr parses a TTL string and applies the system's max/min TTLs
func SanitizeTTLStr(str string, sys logical.SystemView) (time.Duration, error) {
	if str == "" {
		return sys.DefaultLeaseTTL(), nil
	}
	ttl, err := time.ParseDuration(str)
	if err != nil {
		return 0, fmt.Errorf("invalid duration string: %w", err)
	}
	if ttl > sys.MaxLeaseTTL() {
		ttl = sys.MaxLeaseTTL()
	}
	return ttl, nil
}

