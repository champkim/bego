package data

import "time"

type Page struct {
	Index     int       `json:"index"`
	Contents  string    `json:"contents"`	
	UpdatedAt time.Time `json:"updated_at"`
}

type DBHandler interface {
	GetPages() []*Page
	AddPage(page *Page) bool	
	UpdatePage(page *Page) bool		
	GetPage(index int) *Page
	DeletePage() bool
	Close()
}

func NewDBHandler(dbConn string) DBHandler {
	//return newMemHandler()
	//return newSqliteHandler(filepath)
	return newPgHandler(dbConn)
}

