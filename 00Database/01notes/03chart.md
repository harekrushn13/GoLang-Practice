## Histogram Query
```sybase
SELECT 
    toStartOfInterval(timestamp, INTERVAL 5 minute) AS bucket_start,
    AVG(value) AS avg_value_per_bucket
FROM sensor_readings
WHERE sensor_id = 1 
  AND timestamp BETWEEN '2024-03-01 10:00:00' AND '2024-03-01 10:59:00'
GROUP BY bucket_start
ORDER BY bucket_start;

```