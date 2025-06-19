package utils

import "time"

func ParseTimePtr(str *string) *time.Time {
	if str == nil || *str == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, *str)
	if err != nil {
		return nil
	}
	return &t
}

func FormatTimePtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.Format(time.RFC3339)
	return &s
}
