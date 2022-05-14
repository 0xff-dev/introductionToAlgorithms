select distinct(a.num) as  ConsecutiveNums  from Logs as a inner join Logs as b on a.id+1=b.id and a.num=b.num inner join Logs as c on b.id+1=c.id and b.num=c.num;
