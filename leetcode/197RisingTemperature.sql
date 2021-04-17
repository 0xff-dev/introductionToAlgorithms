-- talbe name Weather

select a.id from Weather as a, Weather as b where b.recordDate=date_sub(a.recordDate, interval 1 day) and a.Temperature > b.Temperature;