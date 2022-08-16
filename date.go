package openbd

import "time"

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var lastErr error
	s := string(b)

	if s == `"0000-00-00 00:00:00"` {
		*d = Date{time.Time{}}
		return nil
	}

	v, err := time.Parse(`"2006-01-02 15:04:05"`, s)
	if err != nil {
		lastErr = err
	} else {
		*d = Date{v}
		return nil
	}
	v, err = time.Parse(`"2006-01-02"`, s)
	if err != nil {
		lastErr = err
	} else {
		*d = Date{v}
		return nil
	}

	if lastErr != nil {
		return lastErr
	}
	return nil
}
