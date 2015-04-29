package urlquery

import (
	"github.com/samdolan/go-ordered-map"
	"net/url"
	"strings"
)

// QueryParams handles url query string paramaters respecting order
// We replicate the interface of the default net/url.Values
type QueryParams struct {
	*orderedmap.OrderedMap
	originalRawQuery string
}

// NewQueryParams takes a set of params, validates them and return the params to use
// Seriously, use this rather than constructing QueryParams directly. It'll keep your sanity
func NewQueryParams(params [][]string) (*QueryParams, error) {
	m, err := orderedmap.NewOrderedMap(params)
	if err != nil {
		return nil, err
	}
	queryParams := QueryParams{OrderedMap: m}
	return &queryParams, nil
}

// NewFromQueryStr creates a new QueryParams struct from a raw query string (e.g. "foo=bar&baz=taz")
func NewFromQueryStr(queryStr string) (*QueryParams, error) {
	params, err := parseQuery(queryStr)
	if err != nil {
		return nil, err
	}

	queryParams, err := NewQueryParams(params)
	if err != nil {
		return nil, err
	}
	queryParams.originalRawQuery = queryStr
	return queryParams, nil
}

// Encode encodes the params as a url encoded string
// This is just returns the string method
func (p QueryParams) Encode() string {
	return p.String()
}

// Escape escapes the encoded url using `url.QueryEscape`
func (p QueryParams) Escape() string {
	return url.QueryEscape(p.Encode())
}

// OriginalRawQuery returns the raw query that was passed in to the constructor
// If this instance was not created from a query, then this will be empty string `""`
func (p QueryParams) OriginalRawQuery() string {
	return p.originalRawQuery
}

// parseQuery takes a query string and converts it into an ordered QueryParams
// Originally from the net/url.parseQuery function with tweaks to fit our data structure
func parseQuery(rawQuery string) ([][]string, error) {
	var err error
	params := [][]string{}

	rawQuery = strings.Trim(rawQuery, " ")
	rawQuery = strings.TrimLeft(rawQuery, "?")
	if rawQuery == "" {
		return params, nil
	}

	for rawQuery != "" {
		key := rawQuery
		if i := strings.IndexAny(key, "&;"); i >= 0 {
			key, rawQuery = key[:i], key[i+1:]
		} else {
			rawQuery = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		key, err1 := url.QueryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value, err1 = url.QueryUnescape(value)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		params = append(params, []string{key, value})
	}
	return params, err
}
