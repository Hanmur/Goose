package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimitInterface {
	return MethodLimiter{
		Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}

func (limiter MethodLimiter) Key(context *gin.Context) string {
	uri := context.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}

	return uri[:index]
}

func (limiter MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := limiter.limiterBuckets[key]
	return bucket, ok
}

func (limiter MethodLimiter) AddBuckets(rules ...LimitBucketRule) LimitInterface {
	for _, rule := range rules {
		if _, ok := limiter.limiterBuckets[rule.Key]; !ok {
			limiter.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}

	return limiter
}
