package validation

import (
	"fmt"
	neturl "net/url"
	"strings"
)

func ValidateFlags(urlStr string, requests, concurrency int) error {
	var validationErrors []string

	validateUrl(urlStr, &validationErrors)
	validateRequests(requests, &validationErrors)
	validateConcurrency(concurrency, requests, &validationErrors)

	if len(validationErrors) > 0 {
		return fmt.Errorf("flags internal failed:\n- %s",
			strings.Join(validationErrors, "\n- "))
	}

	return nil
}

func validateUrl(url string, validationErrors *[]string) {
	trimmedUrl := strings.TrimSpace(url)
	if trimmedUrl == "" {
		*validationErrors = append(*validationErrors, "url cannot be empty")
		return
	}
	parsedUrl, err := neturl.Parse(trimmedUrl)
	if err != nil {
		*validationErrors = append(*validationErrors, fmt.Sprintf("invalid url: %v", err))
		return
	}
	if parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		*validationErrors = append(*validationErrors, "url must include scheme and host, e.g: https://example.com")
		return
	}
	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		*validationErrors = append(*validationErrors, "url scheme must be http or https")
	}
}

func validateRequests(requests int, validationErrors *[]string) {
	if requests <= 0 {
		*validationErrors = append(*validationErrors, "requests must be a positive number")
	}
}

func validateConcurrency(concurrency int, requests int, validationErrors *[]string) {
	if concurrency <= 0 {
		*validationErrors = append(*validationErrors, "concurrency must be a positive number")
		return
	}
	if requests > 0 && concurrency > requests {
		*validationErrors = append(*validationErrors, "concurrency cannot exceed requests")
	}
}
