SELECT 
    c.category,
    COUNT(a.income) AS accounts_count
FROM 
    (
        SELECT 'Low Salary' AS category
        UNION ALL
        SELECT 'Average Salary'
        UNION ALL
        SELECT 'High Salary'
    ) c
LEFT JOIN (
    SELECT 
        income,
        CASE
            WHEN income < 20000 THEN 'Low Salary'
            WHEN income >= 20000 AND income <= 50000 THEN 'Average Salary'
            WHEN income > 50000 THEN 'High Salary'
            ELSE '其他'
        END AS category
    FROM Accounts
) a ON c.category = a.category
GROUP BY c.category;
