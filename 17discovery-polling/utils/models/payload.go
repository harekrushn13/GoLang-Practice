package models

type RequestPayload struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	EventName    string `json:"event_name"`
	PluginEngine string `json:"plugin_engine"`
}
