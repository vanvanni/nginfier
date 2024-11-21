package config

type ProxyConfig struct {
	Domain      string `json:"domain"`
	Port        int    `json:"port"`
	Destination string `json:"destination"`
	EnableSSL   bool   `json:"enable_ssl"`
}

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ValidationError struct {
	MissingFields []string `json:"missing_fields"`
}
