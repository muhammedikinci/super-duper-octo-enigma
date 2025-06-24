package lsp

type DiagnosticsNotification struct {
	Notification
	Params DiagnosticsNotificationParams `json:"params"`
}

type DiagnosticsNotificationParams struct {
	URI         string       `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
	Range    Range  `json:"range"`
	Severity int    `json:"severity"`
	Source   string `json:"source"`
	Message  string `json:"message"`
}
