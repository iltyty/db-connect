# DB Connect

## How to run

1. Initialize the mariadb/mysql database, it will listen on localhost:3306.

2. Run the backend first, it will listen on localhost:8000.

3. Then run the frontend, it will work on localhost:5173.



## Database-Mariadb/Mysql

### Setup

1. [Install and setup the database.](https://mariadb.com/kb/en/a-mariadb-primer/)

2. Create database `db_connect`:

   ```mysql
   create database db_connect;
   ```



## Backend-Gin

### How to run

1. [Install the go environment](https://go.dev/doc/install)
2. Run the following commands:

```shell
cd backend
go run main.go
```



## Frontend-Vue

### How to run

```shell
cd frontend
npm install
npm run dev
```

Then access http://localhost:5173 in the browser.
