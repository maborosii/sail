package validate

import (
	"fmt"
	"strconv"
	"time"
)

func IsExpired(sample int, duration time.Duration) (bool, error) {
	strSample := strconv.Itoa(sample)
	if len(strSample) != 10 {
		return false, fmt.Errorf("timestamp type is not supported")
	}

	now := int(time.Now().Unix())
	if (sample < (now - int(duration.Seconds()))) || (sample > (now + int(duration.Seconds()))) {
		return false, fmt.Errorf("sample timestamp is expired")
	}

	return true, nil
}
