package model

type AccessLog struct {
	BaseModel
	RemoteAddr           string `json:"remote_addr" gorm:"type:varchar(50);index;not null;default:''"`
	ClientIP             string `json:"client_ip" gorm:"type:varchar(50);index;not null;default:''"`
	IPLocation           string `json:"ip_location"`
	RequestTimeFloat     int64  `json:"request_time_float" gorm:"not null;default:0"`
	RequestTime          int64  `json:"request_time" gorm:"not null;default:0"`
	RequestID            string `json:"request_id" gorm:"type:varchar(100);unique_index;not null;default:''"`
	Method               string `json:"method" gorm:"type:varchar(10);index;not null;default:''"`
	RequestContentType   string `json:"request_content_type"`
	RequestContentLength int64  `json:"request_content_length" gorm:"not null;default:0"`
	Url                  string `json:"url" gorm:"type:varchar(4096)"`
	Path                 string `json:"path"`
	QueryString          string `json:"query_string" gorm:"type:varchar(4096)"`
	Referer              string `json:"referer" gorm:"type:varchar(4096)"`
	StatusCode           int    `json:"status_code" gorm:"index;not null;default:0"`
	ContentLength        int    `json:"content_length" gorm:"not null;default:0"`
	ContentLengthFormat  string `json:"content_length_format"`
	ResponseTimeFloat    int64  `json:"response_time_float" gorm:"not null;default:0"`
	ResponseTime         int64  `json:"response_time" gorm:"not null;default:0"`
	ResponseContentType  string `json:"response_content_type"`
	Duration             int64  `json:"duration" gorm:"not null;default:0"`
	DurationFormat       string `json:"duration_format"`
	Group                string `json:"group" gorm:"type:varchar(50);index;not null;default:''"`
	UserAgent            string `json:"user_agent"`
	ClientOs             string `json:"client_os"`
	ClientBrowser        string `json:"client_browser"`
}
