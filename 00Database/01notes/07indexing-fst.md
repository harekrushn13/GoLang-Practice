# **1. What is Indexing in Databases?**

Indexing is a **performance optimization technique** that helps the database:

- **Locate data faster** (without scanning the entire table).
- **Reduce I/O operations** (especially important for large datasets).
- **Speed up query filtering and sorting** (`WHERE`, `ORDER BY`).

**Think of an index like a book's table of contents** — instead of reading every page, you jump to the exact page you
need.

---

# **2. Types of Indexes**

Let’s break down key index types:

## **2.1. Primary Index**

- **Defines data order** (e.g., `ORDER BY`).
- Stored as **sparse marks** (ClickHouse reads only relevant parts).

Example:

```sql
CREATE TABLE orders
(
    order_id UInt32,
    product  String,
    price    Float32
) ENGINE = MergeTree()
ORDER BY order_id;
```

**Query:**

```sql
SELECT *
FROM orders
WHERE order_id = 123;  
```

**Only fetches the segment where `order_id=123` is present — fast!**

---

## **2.2. Secondary Index (Data Skipping Index)**

- Helps skip irrelevant data when the primary index isn’t enough.
- ClickHouse supports:
    - **Min-max index** — Tracks the min/max value of each segment.
    - **Set index** — Tracks distinct values (for `IN()` queries).
    - **Bloom filter index** — Good for substring searches.

Example (Bloom filter index):

```sql
CREATE TABLE users
(
    id    UInt32,
    name  String,
    email String
) ENGINE = MergeTree()
ORDER BY id
SETTINGS index_granularity = 8192;

ALTER TABLE users
    ADD INDEX email_index email TYPE bloom_filter(0.01) GRANULARITY 4;
```

**Query:**

```sql
SELECT *
FROM users
WHERE email = 'john@example.com';
```

**Faster because it skips irrelevant granules!**

---

---

# **3. What is an FST (Finite State Transducer)?**

An **FST (Finite State Transducer)** is a type of **finite-state machine** — but instead of just accepting/rejecting
input like a regular finite-state automaton (FSA), it **transforms input to output** efficiently.

**In databases, FSTs are powerful for indexing string data — especially in text search!**

---

## **3.1. How FST Works (Simple Breakdown)**

- **States** → Represent partial matches (prefixes).
- **Transitions** → Move between states based on input characters.
- **Outputs** → Store compressed pointers to data.

**Example:**  
For indexing words like `apple`, `apricot`, `banana`:

```
          (start)
            |
            a
           / \
         p    b
        / \    \
       p   r    a
      /     \     \
     l       i     n
    /         \      \
   e           c      a
                 \
                  o
                   \
                    t
```

**Why FST is powerful?**

- **Space-efficient** (shares common prefixes).
- **Fast lookups** (fewer transitions).
- **Supports range queries & autocomplete efficiently** (e.g., prefix-based searches).

---

---

# 🚀 **4. Indexing in ClickHouse (Detailed)**

Let’s now connect everything to **ClickHouse**:

---

## **4.1. Primary Key Index (Sparse)**

ClickHouse **doesn’t** store a traditional B-Tree index — instead, it stores:

- **Primary key values** every **`index_granularity`** rows (default 8192 rows).
- **Marks** (pointers to rows) — faster to locate segments.

**Example:**

```sql
CREATE TABLE logs
(
    timestamp  DateTime,
    user_id    UInt32,
    event_type String
) ENGINE = MergeTree()
ORDER BY (timestamp, user_id);
```

**Query:**

```sql
SELECT *
FROM logs
WHERE user_id = 123;
```

**Skips all unnecessary parts of the table!**

---

## **4.2. Data Skipping Index (Secondary Index)**

If your `WHERE` conditions **don’t align with the primary key**, secondary indexes help.

**Example:**

```sql
CREATE TABLE products
(
    id       UInt32,
    name     String,
    category String
) ENGINE = MergeTree()
ORDER BY id;

ALTER TABLE products
    ADD INDEX category_index category TYPE set(100) GRANULARITY 2;
```

**Query:**

```sql
SELECT *
FROM products
WHERE category = 'electronics';
```

**Faster filtering!**

---

## **4.3. Full-Text Search Index (FST-based)**

ClickHouse doesn’t have native FST support yet, but it can handle **Bloom filters** for substring search — which mimics
FST behavior for many use cases.

**Example:**

```sql
ALTER TABLE logs
    ADD INDEX event_search_index event_type TYPE bloom_filter(0.01) GRANULARITY 4;
```

**Query:**

```sql
SELECT *
FROM logs
WHERE event_type LIKE '%error%';
```

**Faster partial text matching!**

---

---

# **5. Performance Tips for ClickHouse Indexing**

1️ **Choose the right `ORDER BY` key** — Primary index is sparse, so plan carefully.  
2️ **Use secondary indexes wisely** — Avoid too many; they add overhead.  
3️ **Control granularity** — Smaller granules = faster lookups but larger indexes.

```sql
SET
index_granularity = 4096;
```

4️ **Use `LIMIT` with large tables** — Fetch fewer rows faster.  
5️ **Leverage parallel processing** — ClickHouse splits data into parts for multi-threaded reads.

---

---

# **6. Final Comparison: Indexing vs FST**

| Feature                       | Traditional Indexing | FST-based Indexing                   |
|-------------------------------|----------------------|--------------------------------------|
| **Space Efficiency**          | Moderate             | Very high (compressed structure)     |
| **Speed**                     | Fast (B-Tree/Hash)   | Very fast (state transitions)        |
| **Supports Prefix Search**    | Limited              | Excellent (e.g., autocomplete)       |
| **Supports Substring Search** | Weak (needs LIKE)    | Good (Bloom filters, approximations) |
| **Memory Usage**              | Higher               | Lower (compact automaton)            |

---
