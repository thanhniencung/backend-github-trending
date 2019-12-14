package model

import "time"

type GithubRepo struct {
	Name        string    `json:"name" db:"name,omitempty"`
	Url         string    `json:"url" db:"url,omitempty"`
	Description string    `json:"description" db:"description,omitempty"`
	Color       string    `json:"color" db:"color,omitempty"`
	Lang        string    `json:"lang" db:"lang,omitempty"`
	Fork        string    `json:"fork" db:"fork,omitempty"`
	Stars       string    `json:"stars" db:"stars,omitempty"`
	StarsToday  string    `json:"starsToday" db:"stars_today,omitempty"`
	BuildBy     string    `json:"buildBy" db:"build_by,omitempty"`
	CreatedAt   time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at,omitempty"`
}
