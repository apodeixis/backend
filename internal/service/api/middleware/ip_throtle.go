package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/apodeixis/backend/internal/service/api/ctx"

	"github.com/juju/ratelimit"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

type Limiter struct {
	mu         sync.Mutex
	limiterMap map[string]*ratelimit.Bucket
	limit      int64
	interval   time.Duration
}

func NewLimiter(limit int64, interval time.Duration) *Limiter {
	return &Limiter{
		limiterMap: make(map[string]*ratelimit.Bucket),
		limit:      limit,
		interval:   interval,
	}
}

// IpThrottle is a middleware function that applies rate limiting to incoming requests
// on a per-IP address basis.
func IpThrottle(limiter *Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			limiter.mu.Lock()
			defer limiter.mu.Unlock()
			bucket, ok := limiter.limiterMap[ip]
			if !ok {
				bucket = ratelimit.NewBucketWithQuantum(limiter.interval, limiter.limit, limiter.limit)
				limiter.limiterMap[ip] = bucket
			}
			if bucket.TakeAvailable(1) <= 0 {
				ctx.Log(r).Error("too many requests")
				ape.RenderErr(w, problems.TooManyRequests())
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
