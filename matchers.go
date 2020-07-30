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

type Expectation struct {
	Request  RequestMatcher      `json:"httpRequest,omitempty"`
	Response ExpectationResponse `json:"httpResponse,omitempty"`
	Priority int32               `json:"priority,omitempty"`
}

type ExpectationResponse struct {
	Body       map[string]interface{} `json:"body,omitempty"`
	Headers    map[string][]string    `json:"headers,omitempty"`
	StatusCode int32                  `json:"statusCode,omitempty"`
}

func (m ExpectationResponse) NewResponseOK() ExpectationResponse {
	m.StatusCode = 200
	return m
}

func (m ExpectationResponse) WithJsonBody(json map[string]interface{}) ExpectationResponse {
	m.Body = json
	return m
}

func (m ExpectationResponse) WithHeader(key, value string) ExpectationResponse {
	if m.Headers == nil {
		m.Headers = make(map[string][]string)
	}
	m.Headers[key] = []string{value}
	return m
}
