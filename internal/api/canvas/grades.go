package canvas

import (
	"net/http"
	"time"
)

type Enrollment struct {
	Type                 string   `json:"type"`
	Role                 string   `json:"role"`
	ComputedCurrentScore *float64 `json:"computed_current_score"`
	ComputedFinalScore   *float64 `json:"computed_final_score"`
	ComputedCurrentGrade *string  `json:"computed_current_grade"`
	ComputedFinalGrade   *string  `json:"computed_final_grade"`
}

type Course struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	CourseCode  string       `json:"course_code"`
	Enrollments []Enrollment `json:"enrollments"`
}

type Client struct {
	domain     string
	token      string
	httpClient *http.Client
}

func NewClient(domain string, token string) *Client {
	return &Client{
		domain: domain,
		token:  token,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

}
