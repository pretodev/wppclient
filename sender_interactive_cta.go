package wpp

func (s *Sender) SendCallToActionURL(phoneNumber, body, displayText, url string, opts ...intrOpt) (string, error) {
	intr := newintr("cta_url", body, map[string]any{
		"action": map[string]any{
			"name": "cta_url",
			"parameters": map[string]any{
				"url":          url,
				"display_text": displayText,
			},
		},
	})
	return s.sendRequest(intr.Data(phoneNumber))
}