# filesProcessor
Api for processing transactions from a csv file

RUN PROJECT

- Install docker
- Change the Postgres credentials for your own credentials in dataBase/postgres/initDb.go
- Run the migration m1.sql in dataBase/postgres/migrations
- Execute .deploy.sh or type: "go run main.go" on the Terminal and enter

File Template

> action/fileProcess/docs/txns.csv