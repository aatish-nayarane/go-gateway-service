package routing

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// handle http client tiimeout

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:       100,
		IdleConnTimeout:    90 * time.Second,
		DisableCompression: false,
	},
}

// var  maxretries = 3
// backoff := 200 * time.Second

func doWithRetry(req *http.Request, maxRetries int, backoff time.Duration) (*http.Response, error) {
    var resp *http.Response
    var err error

    for attempt := 1; attempt <= maxRetries; attempt++ {
        // Clone the request body (only works for certain readers)
        reqCopy := req.Clone(req.Context())

        resp, err = httpClient.Do(reqCopy)
        if err == nil {
            return resp, nil // success
        }

        if attempt < maxRetries {
            fmt.Printf("Retry %d failed: %v. Retrying in %v...\n", attempt, err, backoff)
            time.Sleep(backoff)
            backoff *= 2 // exponential backoff
        }
    }

    return nil, fmt.Errorf("all retries failed: %v", err)
}


func ProxyRequest(target string) echo.HandlerFunc {
	fmt.Println("this is your input: ", target)
	return func(c echo.Context) error {
		// Route target url
		fullURL := target + c.Request().URL.Path
		req, err := http.NewRequestWithContext(c.Request().Context(), c.Request().Method, fullURL, c.Request().Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to create request")
		}
		req.Header = c.Request().Header
		// ✅ Add timeout to http.Client.
		// resp, err := httpClient.Do(req)
		// if err != nil {
		// 	return c.String(http.StatusBadGateway, "Backend service unavailable")
		// }

		// ✅ Add retries if backend service is down. 
        resp, err := doWithRetry(req, 3, 200*time.Millisecond)
        if err != nil {
            return c.String(http.StatusBadGateway, "Backend service failed after retries")
        }
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return c.Blob(resp.StatusCode, resp.Header.Get("Content-Type"), body)
	}
}
