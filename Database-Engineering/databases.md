# Understanding Databases
## Introduction
This article will teach green horns in backend web development about databases. Prior knowledge about databases are not required
to read and comprehend the lessons in this article, as it seeks to provide the reader with that knowledge.

**To decide whether you should read this article, you can test your current knowledge against the contents in this article at the end of this 
page.**
## What are Databases
A database is a set of related information. A database can be something as simple as a telephone book, or as complex as a fintech or
banking system. Database systems were among the first computer applications created because of the need to 
replace the tedious usage of paper records.

The first electronic databases were first hierarchical and network database systems, but were replaced by the relational model in 1970. Relational database systems
make use of tables representing linked data records. This model is adopted by all Relational Database Management Systems (RDBMS) and the programming language
used to write them is called the Structured Query Language (SQL). There are alternatives to RDBMS which are written with no tabular structure (NoSQL), but we will look at 
those later.

When working with database systems in a web application, you will most likely be engaging with terms like _Customers_, _Accounts_, and _Transactions._ These
terms are common among fintech solutions which require expert knowledge of database engineering. Let's learn about the properties of relational database systems.
## Relational Database Properties
Relational databases are made up of tables with any convenient number of rows and columns. The number of columns is dependent on the SQL server being used
but it is so large that it may not be an issue. The number of rows is dependent on the storage capacity and ease of maintainability the database engineer
can provide.

The first column of the tables in a relational database is often the _primary key_ used to identify a row in each table. It could be a serial number or a name, depending
on the name of the table or what data is represented in it. It is advisable to use a customer identification number (often written as `CustomerId` in code), however, smaller 
systems can make use of a _compound key_ (a primary key consisting of two or more columns like `firstName` and `lastName`.) if you are certain that no two
customers of the bank can bear the same first and last names. In database engineering terms, customer IDs are called _surrogate keys_ and the first/last names are called
_natural keys_.

Some columns make use of the primary key; they are known as _foreign keys_ and are used to identify what data is related to what customer. In database engineering
terms, this process is known as a _join_.

The general definitions for common database engineering terminologies include:
- entity: a property in the database e.g. customers, parts, addresses, etc.
- column: an individual piece of data stored in a table
- row: a set of columns that describe an entity (called a record)
- table: a set of rows held in memory (non-persistent) or on permanent storage (persistent)
- result set: another name for a non-persistent table, the result of an SQL query
- primary key: column(s) that is used as a unique identifier for each row in the table
- foreign key: column(s) that is used together to identify a single row in another table

## SQL in Relational Database Model
SQL works perfectly with the relational databases model because the result of an SQL query (called the result set) is a table. Hence, storing the result set of the 
query creates a new table in a relational database. SQL does not mean "Structured Query Language." The SQL language has several distinct parts. The important parts in 
database engineering include:
- SQL schema statements: used to define the data structures stored in a database. An example of this include using the `create table` statement to create a new table.
- SQL data statements: used to manipulate the data structures defined by the schema statements. An example is the process of populating a new table.
- SQL transaction statements: used to begin, end, and roll back transactions

## Using SQL STATEMENTS
Let's look at an example where these statements are adopted. To create a table, the statement is similar to:
```sql
CREATE TABLE  company
(
    company_id SMALLINT ,
    company_name VARCHAR(30),
    CONSTRAINT pk_company PRIMARY KEY (company_id)
);
```
Next, let us uses a data statement to insert a row into the table above:
```sql
INSERT INTO company (company_id, company_name)
VALUES (0007, 'MacKie Groups');
```
Finally, let us retrieve data from the table created:
```sql
mysql < SELECT company_name
    -> FROM company
    -> WHERE company_id = 0007;
```





# Test Your Knowledge
1. What kind of databases are the most complex?
2. What is a surrogate key?
3. 