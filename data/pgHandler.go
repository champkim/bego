package data

import (
	"database/sql" //모든 sql 에 공통되는 interface ,
	"time"

	_ "github.com/lib/pq" //sqlite 패키지도 import 해야한다. _ 는 암시적 사용 예시
)

//인터페이스를 만들기 위한 구조체 정의
type pgHandler struct {
	db *sql.DB
}

func (p *pgHandler)GetPages() []*Page {
	pages := []*Page{}

	rows, err := p.db.Query("SELECT pageindex, contents, updatedat FROM pages")
	if err != nil {
		panic(err)
	}
	defer rows.Close() //function 이 종료 되기전에 rows 를 Close 시켜라 

	for rows.Next() {
		var p Page
		rows.Scan(&p.Index, &p.Contents, &p.UpdatedAt)
		pages = append(pages, &p)
	}
	
	return pages
}

func (p *pgHandler)AddPage(page *Page) bool {	
	stmt, err := p.db.Prepare("INSERT INTO pages (pageindex, contents, updatedat) VALUES ($1, $2, NOW())")		
	if err != nil {
		panic(err)
	}	
	
 	rst, err := stmt.Exec(page.Index, page.Contents)
		if err != nil {
		panic(err)
	}
	
	page.UpdatedAt = time.Now() 
	cnt, _ := rst.RowsAffected()
	return (cnt > 0)	
}	

func (p *pgHandler)UpdatePage(page *Page) bool {

	stmt, err := p.db.Prepare("UPDATE pages SET contents=$1, updatedat=NOW() WHERE pageindex=$2")	
	if err != nil {
		panic(err)
	}	
 	rst, err := stmt.Exec(page.Contents, page.Index)
		if err != nil {
		panic(err)
	}
		
	page.UpdatedAt = time.Now() 
	cnt, _ := rst.RowsAffected()
	return (cnt > 0)	
}	

func (p *pgHandler)GetPage(index int) *Page {
	
	rows, err := p.db.Query("SELECT pageindex, contents, updatedat FROM pages WHERE pageindex=$1", index)	
	if err != nil {
		panic(err)
	}
	defer rows.Close() //function 이 종료 되기전에 rows 를 Close 시켜라 
	//log.Println(">>>>>>>>>> pageindex >>>>>> " +  strconv.Itoa(index))			
	if rows.Next() {
		var pg Page //= new(Page)	
		rows.Scan(&pg.Index, &pg.Contents, &pg.UpdatedAt)
		return &pg
	} else {
		return nil
	}
}

func (p *pgHandler)DeletePage() bool {

	stmt, err := p.db.Prepare("DELETE FROM pages")	
	if err != nil {
		panic(err)
	}

 	rst, err := stmt.Exec()
	if err != nil {
		panic(err)
	}	
	cnt, _ := rst.RowsAffected()
	//return true
	return cnt >= 0
}

func (p *pgHandler)Close() {
	p.db.Close()
}

func newPgHandler(dbConn string) DBHandler {
	database, err := sql.Open("postgres", dbConn)
	if err != nil {
		panic(err)
	}

	statement, err := database.Prepare(
		`CREATE TABLE IF NOT EXISTS pages (
			pageindex int PRIMARY KEY,
			contents  TEXT,			
			updatedat TIMESTAMP
		)`)

	if err != nil {
		panic(err)
	}

	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}
		
	return &pgHandler{db: database}	
}