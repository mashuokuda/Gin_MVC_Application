package notify

type NotifyJSON []struct {
	DiscussID int    `json:"discussID"`
	Hash      string `json:"hash"`
	Level     int    `json:"level"`
	Comment   string `json:"comment"`
}
