# go backend test [![go](https://miro.medium.com/max/700/1*Ifpd_HtDiK9u6h68SZgNuA.png)](https://go.dev/)

> start go backend

I will fill in the future.

## Description
모자이크 화면의 데이터 정보를 저장 하고 업데이트 한다.

## Environment
* go lang 
* windows11
* vscode 

## Prerequisite

* [[router and dispatcher] gorilla mux](https://github.com/gorilla/mux)
* [[render] unrolled](https://github.com/unrolled/render)
* [[cors] rs/cors](https://github.com/rs/cors)
* [[lib] negroni](https://github.com/urfave/negroni)

* [[.env] godotenv](https://github.com/joho/godotenv)

* [[db] postgresq](https://https://github.com/lib/pq)

* [[tdd] goconvey](https://https://github.com/smartystreets/goconvey)
* [[tdd] assert](https://https://github.com/stretchr/testify/tree/master/assert)

## Files
* pages.go 
```
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
```
* 원하는 저장소 memory, db, file 등에 대하여 DBHandler interface 를 구현하면  다른 소스파일에 영향 없이 Adapter 처럼 변경 사용 가능 합니다.
  (memHandler.go, pgHandler.go) 
## Usage
* tdd
```
\bego\app>goconvey      

\bego\app>go test
```
