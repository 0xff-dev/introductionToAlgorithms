CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  DECLARE my_int INT DEFAULT -1;
  SET my_int = N - 1;
  RETURN (
        SELECT distinct(salary) FROM Employee
        ORDER BY salary DESC
        LIMIT 1 OFFSET my_int
  );
END
