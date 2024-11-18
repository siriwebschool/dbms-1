# Primary Index Lookup

## Objective

Given that the Index and the pages are existing, write a program to lookup the index and fetch the record from the page.

## Assumptions

- The index is a JSON containing the key value pairs of the primary key and the page in which the record is stored.  
eg: [{"1" : "Page1"}] - The record with primary key "1" is stored in the page 1.

- The heap is in the form of pages. Each page is a seperate file that contains 10 records each.

- Each record in a page is separated by a new line character "/n" and each column in the record is seperated by a comma ",".

## Solution Usage

- make sure GO is installed on your PC.

- run the command "go run db.go --search  \* insert primary key \* "

## Solution Approach

Check out the comments in the db.go file