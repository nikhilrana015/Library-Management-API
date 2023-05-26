# **Library Management API CRUD Operations**
Created a Library Management API performing all CRUD operations. Using golang as a programming language and [gorilla-mux](https://github.com/gorilla/mux) as golang-web framework. Using sql as a database, sql-workbench as visual tool for database and [GORM](https://gorm.io/docs/index.html) as Object Relational Mapping (ORM) library. Used [Postman](https://www.postman.com/) to design, build and test API's crud operations.

</br>

## **1. Add a new book**

- ## *a) In-Postman*
![Add new book-1](./screenshots/Add%20a%20new%20book%20-%20Postman.png)</br></br>
![Add new book-3](./screenshots/Add%20a%20new%20book-3-Postman.png)

- ## *b) Sql-workbench*
![Add new book-1](./screenshots/Add%20a%20new%20book-Sql.png)</br></br>
![Add new book-3](./screenshots/Add%20a%20new%20book-3-Sql.png)

## **2. Get book with book-id**

- ## *a) In-Postman*
![get book with book-id](./screenshots/Get-a-book-by-id-1-Postman.png)</br></br>
![book-id not found](./screenshots/Get-a-book-by-id-Not-Found-Postman.png)

## **3. Update book with book-id**

- ## *a) In-Postman*
![Update book with book-id-1](./screenshots/update-book-with-bookId-1-Postman.png)

- ## *b) Sql-workbench*
![Update book with book-id-1](./screenshots/update-book-with-bookId-1-Sql.png)


## **4. Get all books**

- ## *a) In-Postman*
![get all books](./screenshots/GetAllBooks-1-Postman.png)

- ## *b) Sql-workbench*
![get all books](./screenshots/GetAllBooks-1-Sql.png)


## **5. Delete book**

- ## *a) In-Postman*
![delete book with book-id](./screenshots/delete-bookBy-Id-1-Postman.png)</br></br>
![book not found](./screenshots/delete-bookBy-Id-NotFound-Postman.png)

- ## *b) Sql-workbench*
![delete book with book-id](./screenshots/delete-bookBy-Id-1-Sql.png)


## **6. Check-out Book**

- ## *a) In-Postman*
![check-out book with book-id](./screenshots/check-out-book-1-Postman.png)</br></br>
![check-out book Quantity-Zero](./screenshots/check-out-book-QuantityZero-Postman.png)

- ## *b) Sql-workbench*
![check-out book with book-id](./screenshots/check-out-book-1-Sql.png)

## **7. Check-in Book**

- ## *a) In-Postman*
![check-in book with book-id](./screenshots/check-in-book-1-Postman.png)

- ## *b) Sql-workbench*
![check-in book with book-id](./screenshots/check-in-book-1-Sql.png)


