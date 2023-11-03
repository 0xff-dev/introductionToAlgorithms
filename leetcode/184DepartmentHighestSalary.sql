select d.name as Department,e.name Employee, e.salary as Salary 
  from Employee as e, Department as d, (select departmentId as d , max(salary) as b from Employee group by departmentId) as c
  where e.departmentId=d.id and c.d=d.id and c.b=e.salary;
