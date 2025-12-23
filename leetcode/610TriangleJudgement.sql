SELECT 
  a.*, 
  (CASE WHEN (a.x + a.y > a.z AND a.x + a.z > a.y AND a.y + a.z > a.x) THEN 'Yes' ELSE 'No' END) AS triangle
FROM Triangle a;
