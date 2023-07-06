package models

type ShortURLStruct struct {
	HashURL string   `json:"hashUrl"`
	URL     string   `json:"url"`
	UserIDs []string `json:"userIds"`
}

func (s *ShortURLStruct) IsEmpty() bool {
	if s.HashURL == "" && s.URL == "" && len(s.UserIDs) == 0 {
		return true
	}

	return false
}

func (s *ShortURLStruct) UserExist(userID string) bool {
	exist := false

	for _, v := range s.UserIDs {
		if v == userID {
			exist = true
			break
		}
	}

	return exist
}

type AllURLsStruct struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type UserStruct struct {
	UserID   string
	UserURLs []string
}
