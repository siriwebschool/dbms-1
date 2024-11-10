# 6 Oct 2024, 13 Oct 2024

## Transactions
What is a transaction? begin, write, commit

## ACID

### Atomicity
  All transactions should either succeed or not succeed at all.

### Consistency
- Consistency in data
- Consistency in read
- Eventual Consistency

### Isolation

**Read Phenomenons**

- Dirty Reads
- Non Repeatable Reads
- Phantom Reads
- Lost Updates

**Isolation Levels**

- Read uncommitted
- Read committed
- Repeatable reads (locks)
- Snapshots
- Serialisable

### Durability
  - Write ahead log
  - Asynchronous Snapshot

# 20 Oct 2024
How is data stored in the db?
 - Page
   - Seperator bits
 - Heap
 - Indexing
   - b-tree (algorithm)
 - Row based db
 - Column based db

# 27 Oct 2024

## Indexing
- Primary Indexing
- Secondary Indexing

# 03 Nov 2024
- Recap + Test

# 10 Nov 2024
- Build your own database. Using the dummy database given above, write a script which when run like the following: `node index.js 26` will fetch the details of the 26th employee in the database.
