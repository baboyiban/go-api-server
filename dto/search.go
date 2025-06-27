package dto

import "net/url"

func ExtractAllowedParams(values url.Values, allowedFields map[string]bool) map[string]string {
	params := map[string]string{}
	for key, vals := range values {
		if allowedFields[key] && len(vals) > 0 && vals[0] != "" {
			params[key] = vals[0]
		}
	}
	return params
}
