# Secondary Index Lookup

## Objective

- Given that the Index and the pages are existing, write a program to lookup the index and fetch the records from the pages.
- The index here is a secondary index
- The index is against the last column in the records.<br>
Eg: 10,2010,Person 10,Rank3<br>
Assume this record is in the page3,<br>
The index will look like this: <br>
{ "Rank3" : ["page3"] }


## Assumptions
- The Rank column only has 6 values
>- Rank1
>- Rank2
>- Rank3
>- Rank4
>- Rank5
>- nil

- The index is a JSON containing the key value pairs of the rank column value and the array of pages in which the records containing the rank value is stored.  
eg: [ { "Rank1" : [ "Page1","Page2" ] } ] - The records with the Rank value "Rank1" is stored in the pages 1 and 2.

- The heap is in the form of pages. Each page is a seperate file that contains 10 records each.

- Each record in a page is separated by a new line character "/n" and each column in the record is seperated by a comma ",".
