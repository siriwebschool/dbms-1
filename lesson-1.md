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
