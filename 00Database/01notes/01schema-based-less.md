## **1. Schema-Based Databases**

These databases **require a predefined schema** — you define tables, columns, data types, and relationships before
inserting data. They ensure data consistency and structure.

### **Relational Databases (SQL)**

- **Structure:** Tables (rows & columns) with strict schema rules.
- **Query Language:** SQL (Structured Query Language).
- **Use Cases:** Financial systems, CRM, e-commerce, anything needing ACID compliance (Atomicity, Consistency,
  Isolation, Durability).
- **Examples:** MySQL, PostgreSQL, Oracle, SQL Server.

**Example Schema Definition:**

```sql
CREATE TABLE users
(
    id         INT PRIMARY KEY,
    name       VARCHAR(100),
    email      VARCHAR(100) UNIQUE,
    created_at TIMESTAMP
);
```

---

### **Columnar Databases (for Analytics)**

- **Structure:** Columns instead of rows — optimized for fast reading of large datasets.
- **Use Cases:** Big data analytics, event tracking, reporting.
- **Examples:** ClickHouse, Apache Cassandra, Amazon Redshift.

**Example Schema Definition (ClickHouse):**

```sql
CREATE TABLE logs
(
    timestamp     DateTime,
    service       String,
    response_time Float64
) ENGINE = MergeTree()
ORDER BY timestamp;
```

---

## **2. Schema-Less Databases (NoSQL)**

These are more flexible — data can have different structures in the same collection. Ideal for fast-changing data
models.

---

### **Document Databases**

- **Structure:** Stores data as JSON or BSON documents (like nested objects).
- **Use Cases:** Content management, user profiles, catalogs.
- **Examples:** MongoDB, Couchbase, Firebase Firestore.

**Example Document (MongoDB):**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "preferences": {
    "theme": "dark",
    "notifications": true
  }
}
```

---

### **Key-Value Databases**

- **Structure:** Simple key-value pairs.
- **Use Cases:** Caching, real-time session management, configurations.
- **Examples:** Redis, DynamoDB, Riak.

**Example (Redis):**

```bash
SET user:1001 "John Doe"
GET user:1001
```

---

### **Wide-Column Stores**

- **Structure:** Like rows and columns, but each row can have different columns.
- **Use Cases:** Time-series data, recommendation engines.
- **Examples:** Apache Cassandra, HBase, ScyllaDB.

---

### **Graph Databases**

- **Structure:** Nodes (entities) connected by edges (relationships).
- **Use Cases:** Social networks, fraud detection, recommendation engines.
- **Examples:** Neo4j, ArangoDB, Amazon Neptune.

**Example (Neo4j):**

```cypher
CREATE (p:Person {name: "John"})-[:FRIENDS_WITH]->(q:Person {name: "Jane"})
```

---

## **3. Summary: Schema-Based vs Schema-Less**

| Feature               | Schema-Based (SQL) | Schema-Less (NoSQL)               |
|-----------------------|--------------------|-----------------------------------|
| **Schema definition** | Strict, predefined | Flexible, dynamic                 |
| **Scalability**       | Vertical scaling   | Horizontal scaling                |
| **Query language**    | SQL                | Varies (JSON, Key-Value, etc.)    |
| **Data consistency**  | Strong consistency | Eventual consistency (most cases) |
| **Use case examples** | Banking, ERP, CRM  | IoT, social media, analytics      |

---