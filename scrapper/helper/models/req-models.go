package models

type Data struct {
	BatchComplete string      `json:"batchcomplete"`
	Warnings      interface{} `json:"warnings"`
	Query         QueryData   `json:"query"`
}

type QueryData struct {
	Normalized interface{} `json:"normalized"`
	Pages      PagesData   `json:"pages"`
}

type PagesData struct {
	Page Page `json:"25928"`
}

type Page struct {
	PageID    int         `json:"pageid"`
	Ns        int         `json:"ns"`
	Title     string      `json:"title"`
	Revisions []Revisions `json:"revisions"`
}

type Revisions struct {
	Html string `json:"*"`
}
