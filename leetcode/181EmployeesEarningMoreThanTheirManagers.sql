select A.name as Employee from Employee A inner join Employee as B on A.managerId=B.id and A.salary > B.salary;
