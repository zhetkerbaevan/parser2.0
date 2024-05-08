# Parser on Go
The parser extracts data from the site https://kolesa.kz/cars/almaty/ and saves it to the MySQL database.
  
The data is saved according to the structure of Automobile, which has the fields:
* id
* model
* price
* year

## Install
1. Clone the repository:
```sh
git clone https://github.com/zhetkerbaevan/parser.git
```
2. Go to the project directory:
 ```sh
  cd parser
```
3. Install dependencies:
 ```sh
go mod tidy
```
4. You will also need a MySQL database.   
You need to create a parser_go database and create a table there, according to the structure of Automobile.  
Then replace the url to the database in the database.go file to your:
 ```sh
  db, err := sql.Open("mysql", "your_url/parser_go")
```
## Start
Go to the cmd/parser folder
 ```sh
cd cmd/parser
```
then run the application
 ```sh
go run main.go
```
