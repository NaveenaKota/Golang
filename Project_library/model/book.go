package model

type Books struct {
	book_id        int    `json:"book_id",bun:"book_id"`
	year_published string `json:"year_published",bun:"year_published"`
	book_title     string `json:"book_title",bun:"book_title"`
	author_name    string `json:"author_name",bun:"author_name"`
	book_category  string `json:"book_category",bun:"book_category"`
}

type Item struct {
	book_id    int    `json:"book_id",bun:"book_id"`
	start_date string `json:"start_date",bun:"start_date"`
	end_date   string `json:"end_date",bun:"end_date"`
	person_id  string `json:"person_id",bun:"person_id"`
}

type Member struct {
	Person_id   string    `json:"person_id",bun:"person_id"`
	Name        string `json:"name",bun:"name"`
	PhoneNumber int    `json:"phoneNumber",bun:"phoneNumber"`
	Email       string `json:"email",bun:"email"`
	Department  string `json:"department",bun:"department"`
}

type location struct {
	book_id  int `jsom:"book_id",bun:"book_id"`
	floor_no int `json:"floor_no",bun:"floor_no"`
	shelf_no int `json:"shelf_no",bun:"shelf_no"`
}
