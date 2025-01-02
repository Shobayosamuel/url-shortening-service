package routes

import (
	"time"

)

type request struct {
	URL				string			`json:"url"`
	CustomShort		string			`json:"custom_short"`
	Expiry			time.Duration	`json:"expiry"`
}

type response struct {
	URL				string			`json:"url"`
	CustomShort		string			`json:"custom_url"`
	Expiry			time.Duration	`json:"expiry"`
	XRateRemaining	int				`json:"rate-remaining"`
	XRateLimitReset	time.Duration	`json:"rate-limit-reset"`
}