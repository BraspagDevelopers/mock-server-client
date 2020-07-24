package mockserver

type RequestMatcher struct {
	Method  string              `json:"method,omitempty"`
	Path    string              `json:"path,omitempty"`
	Body    BodyMatcher         `json:"body,omitempty"`
	Headers map[string][]string `json:"headers,omitempty"`
}

func (m RequestMatcher) WithHeader(key, value string) RequestMatcher {
	if m.Headers == nil {
		m.Headers = make(map[string][]string)
	}
	m.Headers[key] = []string{value}
	return m
}

type BodyMatcher struct {
	Type      string                 `json:"type,omitempty"`
	JSON      map[string]interface{} `json:"json,omitempty"`
	MatchType string                 `json:"matchType,omitempty"`
}

const (
	MatchBodyJSON = "JSON"
)
const (
	StrictMatch   = "STRICT"
	TolerantMatch = "ONLY_MATCHING_FIELDS"
)

func (m RequestMatcher) WithJsonFields(json map[string]interface{}) RequestMatcher {
	m.Body = BodyMatcher{
		Type:      MatchBodyJSON,
		JSON:      json,
		MatchType: TolerantMatch,
	}
	return m
}

type ResponseMatcher struct {
	Body       BodyMatcher         `json:"body,omitempty"`
	Headers    map[string][]string `json:"headers,omitempty"`
	StatusCode int32               `json:"statusCode,omitempty"`
}

func (m ResponseMatcher) WithJsonFields(json map[string]interface{}) ResponseMatcher {
	m.Body = BodyMatcher{
		JSON: json,
	}
	return m
}

func (m ResponseMatcher) WithHeader(key, value string) ResponseMatcher {
	if m.Headers == nil {
		m.Headers = make(map[string][]string)
	}
	m.Headers[key] = []string{value}
	return m
}
