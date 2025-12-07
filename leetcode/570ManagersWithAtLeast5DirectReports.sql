# Write your MySQL query statement belo
select name from Employee where id in (select managerId from Employee where managerId is not null group by managerID having(count(*))>=5)
