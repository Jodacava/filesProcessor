# filesProcessor
Api for processing transactions from a csv file

RUN PROJECT

- Install docker
- Change the Postgres credentials for your own credentials in dataBase/postgres/initDb.go
- Run the migration m1.sql in dataBase/postgres/migrations
- Execute .deploy.sh or type: "go run main.go" on the Terminal and enter

CURL FOR LOCAL PURPOSE

curl --location 'http://localhost:8080/api/files-processor-api/action/file/process' \
--form 'file=@"/<route to the CSV file>"' \
--form 'user-email="<email destination>"'

FILE TEMPLATE

> action/fileProcess/docs/txns.csv