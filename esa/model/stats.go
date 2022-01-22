package model

type Stats struct {
	Members            int `json:"members"`
	Posts              int `json:"posts"`
	PostsWip           int `json:"posts_wip"`
	PostsShipped       int `json:"posts_shipped"`
	Comments           int `json:"comments"`
	Stars              int `json:"stars"`
	DailyActiveUsers   int `json:"daily_active_users"`
	WeeklyActiveUsers  int `json:"weekly_active_users"`
	MonthlyActiveUsers int `json:"monthly_active_users"`
}
