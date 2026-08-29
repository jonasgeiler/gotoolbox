package gotoolbox

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Download(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	return resp, nil
}

// DownloadProgressPrinter is an io.Writer that counts the number of bytes
// written to it, and prints a write percentage to the console once the
// write process goes on for more than a second.
type DownloadProgressPrinter struct {
	ExpectedBytes             int64
	expectedBytesFormatted    string
	expectedBytesFormattedLen int
	writtenBytes              int64

	LogPrefix string

	hasPrintedWithExpectedBytes bool
	startOrLastPrintTime        time.Time
}

var _ io.Writer = (*DownloadProgressPrinter)(nil)

func (d *DownloadProgressPrinter) Write(p []byte) (int, error) {
	// Initialize start time. This cause download progress messages to only
	// appear for slower/bigger downloads.
	if d.startOrLastPrintTime.IsZero() {
		d.startOrLastPrintTime = time.Now()
	}

	// Count up the number of bytes written.
	n := len(p)
	d.writtenBytes += int64(n)

	// If more than a second has passed since last print OR we previously
	// printed and now reached 100%, print another download progress message.
	if time.Since(d.startOrLastPrintTime) >= time.Second || (d.hasPrintedWithExpectedBytes && d.writtenBytes >= d.ExpectedBytes) {
		if d.ExpectedBytes > 0 {
			// Cache formatted representation of expected bytes.
			if d.expectedBytesFormatted == "" {
				d.expectedBytesFormatted = formatBytes(d.ExpectedBytes)
				d.expectedBytesFormattedLen = len(d.expectedBytesFormatted)
			}

			writtenBytesFormatted := formatBytes(d.writtenBytes)
			fmt.Printf(
				"%sProgress: %6.2f%% (%s%s / %s)\n",
				d.LogPrefix,
				(float64(d.writtenBytes)/float64(d.ExpectedBytes))*100,
				strings.Repeat(
					" ",
					max(
						0,
						d.expectedBytesFormattedLen-len(writtenBytesFormatted),
					),
				),
				writtenBytesFormatted,
				d.expectedBytesFormatted,
			)
			d.hasPrintedWithExpectedBytes = true
		} else {
			fmt.Printf(
				"%sProgress: N/A (%s)\n",
				d.LogPrefix, formatBytes(d.writtenBytes),
			)
		}
		d.startOrLastPrintTime = time.Now()
	}

	return n, nil
}

func formatBytes(n int64) string {
	const unit = 1024
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}

	div, exp := int64(unit), 0
	for n >= unit*div && exp < 4 {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %ciB", float64(n)/float64(div), "KMGT"[exp])
}
