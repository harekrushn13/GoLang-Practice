## **1. What is Partitioning?**

Partitioning is the process of **dividing a large table into smaller, more manageable pieces (partitions)** while still
treating it as one table in your queries.

**Why?**

- **Improved performance** — Queries can scan only relevant partitions instead of the entire table.
- **Faster reads and writes** — Smaller data chunks mean quicker access.
- **Easier maintenance** — Backups, archiving, and data purging are simpler.
- **Load distribution** — Helps balance data across servers in distributed databases.

---

## **2. Types of Partitioning**

Let’s break this down into **horizontal** and **vertical** partitioning:

---

### **2.1. Horizontal Partitioning (Sharding)**

Here, rows are divided into multiple partitions based on a condition — each partition has the same columns but different
rows.

🔹 **Types of Horizontal Partitioning:**

1️ **Range Partitioning**

- Data is divided based on value ranges in a column (e.g., dates, IDs).
- Example:
   ```sql
   CREATE TABLE sales (
       sale_id INT,
       sale_date DATE,
       amount DECIMAL
   ) 
   PARTITION BY RANGE (YEAR(sale_date)) (
       PARTITION p2023 VALUES LESS THAN (2024),
       PARTITION p2024 VALUES LESS THAN (2025)
   );
   ```

**Use case:** Time-series data (logs, sales data, etc.).

---

2️ **List Partitioning**

- Rows are partitioned based on a list of values.
- Example:
   ```sql
   CREATE TABLE orders (
       order_id INT,
       region VARCHAR(20)
   )
   PARTITION BY LIST(region) (
       PARTITION east VALUES IN ('NY', 'NJ', 'PA'),
       PARTITION west VALUES IN ('CA', 'NV', 'OR')
   );
   ```

**Use case:** Data categorized by region, product type, etc.

---

3️ **Hash Partitioning**

- Rows are distributed using a **hash function** on a column’s value (e.g., `MOD(id, 4)` for 4 partitions).
- Example:
   ```sql
   CREATE TABLE users (
       user_id INT,
       name VARCHAR(100)
   )
   PARTITION BY HASH(user_id) PARTITIONS 4;
   ```

**Use case:** Uniform distribution when ranges or lists aren’t predictable (user IDs, sensor data).

---

4️ **Composite Partitioning**

- Combines two or more strategies (e.g., Range + Hash).
- Example:
   ```sql
   CREATE TABLE logs (
       log_id INT,
       log_date DATE,
       server_id INT
   )
   PARTITION BY RANGE (YEAR(log_date))
   SUBPARTITION BY HASH(server_id) SUBPARTITIONS 4 (
       PARTITION p2023 VALUES LESS THAN (2024),
       PARTITION p2024 VALUES LESS THAN (2025)
   );
   ```

**Use case:** Multi-dimensional data (time + server, region + category).

---

### **2.2. Vertical Partitioning**

Here, columns are split into multiple tables based on how frequently they’re used or how large they are.

**Example:**

- Table 1 → Frequently queried data (`id`, `name`, `email`).
- Table 2 → Rarely queried data (`profile_picture`, `preferences`).

```sql
CREATE TABLE user_core
(
    user_id INT PRIMARY KEY,
    name    VARCHAR(100),
    email   VARCHAR(100)
);

CREATE TABLE user_details
(
    user_id         INT,
    profile_picture BLOB,
    preferences     JSON
);
```

**Use case:** Large, rarely used columns (JSON data, BLOBs), performance tuning.

---

## **3. When to Use Partitioning?**

Partitioning isn’t always the right choice — it works best when:

- **Huge tables** (millions/billions of rows).
- **Frequent range-based queries** (e.g., `WHERE date BETWEEN ...`).
- **Need faster data archiving/purging** (drop old partitions instead of `DELETE`).
- **Parallel processing or distributed setup** (e.g., ClickHouse, PostgreSQL, MySQL partitioned tables).

---

## **4. Challenges & Trade-offs**

Partitioning isn’t free — here’s what to watch out for:

🔹 **Query Complexity:**

- Queries must match the partition key to benefit (e.g., `WHERE sale_date = '2024-01-01'`).
- Queries without the key may scan all partitions.

🔹 **Index Limitations:**

- Some databases restrict indexes on partitioned tables.

🔹 **Data Skew:**

- Uneven partitioning (e.g., one region has 90% of data) can kill performance.

🔹 **Inserts and Updates:**

- Wrong partitioning can slow down `INSERT` or `UPDATE` due to partition checks.

---

## **5. Partitioning in Popular Databases**

| Database       | Supported Partitioning Types                  | Notes                                          |
|----------------|-----------------------------------------------|------------------------------------------------|
| **PostgreSQL** | Range, List, Hash, Composite                  | Declarative partitioning since v10.            |
| **MySQL**      | Range, List, Hash, Key (auto-hash), Composite | Limited index support on partitions.           |
| **ClickHouse** | Range, Hash (by key), Partition by Date or ID | Great for time-series and analytics workloads. |
| **MongoDB**    | Sharding (Hash, Range, Zone-based)            | Built-in horizontal scaling.                   |
| **Oracle**     | Range, List, Hash, Composite                  | One of the most advanced partitioning systems. |

---

## **6. Final Takeaways**

- **Horizontal Partitioning** (sharding) divides rows; ideal for scaling.
- **Vertical Partitioning** divides columns; ideal for large, wide tables.
- **Range/List/Hash/Composite** strategies handle different workloads.
- **Carefully choose the partition key** — it affects performance directly.

---

# **2. Partitioning Strategies in ClickHouse**

ClickHouse supports partitioning mainly for performance and scalability. Let’s dive into its strategies:

---

### **2.1. Partitioning Basics in ClickHouse**

👉 In ClickHouse, **partitions** are more like **subdirectories** storing data parts, while **primary keys** handle
sorting within each partition.

### **Syntax Example:**

```sql
CREATE TABLE sales
(
    order_id  UInt32,
    sale_date Date,
    region    String,
    amount    Float32
) ENGINE = MergeTree()
PARTITION BY toYYYYMM(sale_date)
ORDER BY (region, sale_date);
```

---

### **2.2. Partitioning Strategies**

Let’s explore the most common strategies:

---

1️ **Time-based Partitioning (Range-like)**

- **Best for:** Logs, events, time-series data.
- **Example:** Partition by month:
   ```sql
   PARTITION BY toYYYYMM(sale_date)
   ORDER BY (region, sale_date);
   ```
- **Why?** Queries like `WHERE sale_date BETWEEN '2024-01-01' AND '2024-01-31'` hit only the January 2024 partition.

---

2️ **Category-based Partitioning (List-like)**

- **Best for:** Region-based data, product categories.
- **Example:** Partition by region:
   ```sql
   PARTITION BY region
   ORDER BY sale_date;
   ```
- **Why?** Queries like `WHERE region = 'Europe'` scan only the relevant partition.

---

3️ **Hash-based Partitioning (Uniform distribution)**

- **Best for:** Evenly distributing data (e.g., user IDs).
- **Example:** Partition by hash of user_id:
   ```sql
   PARTITION BY cityHash64(user_id) % 4
   ORDER BY user_id;
   ```
- **Why?** Ensures balanced data across 4 partitions.

---

4️ **No Partitioning (Single partition)**

- **Best for:** Small datasets.
- **Example:**
   ```sql
   PARTITION BY tuple()
   ORDER BY (order_id);
   ```

---

## **3. Final Takeaways**

| Feature               | Horizontal Partitioning       | Vertical Partitioning            |
|-----------------------|-------------------------------|----------------------------------|
| **Splits Data By**    | Rows (row subsets)            | Columns (column subsets)         |
| **Best For**          | Large datasets, regional data | Wide tables, rarely used columns |
| **Performance Boost** | Selects only relevant rows    | Skips unused columns             |
| **Trade-offs**        | Complex joins, data skew      | Join overhead, schema complexity |

---