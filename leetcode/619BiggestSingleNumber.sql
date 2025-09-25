# Write your MySQL query statement below
select max(a.num) num  from (select num from MyNumbers group by num having(count(*))=1) a
