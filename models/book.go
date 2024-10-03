package models

type Book struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	EditionNumber string `json:"edition_number"`
}

type PickUpSchedule struct {
	BookTitle  string `json:"book_title"`
	PickUpTime string `json:"pick_up_time"`
}

type ResponseBook struct {
	Key          string   `json:"key"`
	Title        string   `json:"title"`
	AuthorName   []string `json:"author_name"`
	EditionCount int      `json:"edition_count"`
}
