package random_stress_utils

import (
	"math/rand"
	"net/http"
	"time"
)

func RandomSleep(min, max time.Duration) {
	duration := time.Duration(rand.Int63n(int64(max-min+1)) + int64(min))
	time.Sleep(duration)
}

func RandomHttpStatus() int {
	commonStatusCodes := []int{
		// 2xx Success
		http.StatusOK,        // 200
		http.StatusCreated,   // 201
		http.StatusNoContent, // 204
		// 3xx Redirection
		http.StatusFound,       // 302
		http.StatusNotModified, // 304
		// 4xx Client Error
		http.StatusBadRequest,       // 400
		http.StatusUnauthorized,     // 401
		http.StatusForbidden,        // 403
		http.StatusNotFound,         // 404
		http.StatusMethodNotAllowed, // 405
		http.StatusTooManyRequests,  // 429
		// 5xx Server Error
		http.StatusInternalServerError, // 500
		http.StatusNotImplemented,      // 501
		http.StatusServiceUnavailable,  // 503
	}

	return commonStatusCodes[rand.Intn(len(commonStatusCodes))]
}
