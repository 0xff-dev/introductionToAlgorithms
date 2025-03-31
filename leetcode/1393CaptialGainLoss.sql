# Write your MySQL query statement below
select b.stock_name stock_name, IFNULL(b.p, 0) - IFNULL(a.p, 0) as capital_gain_loss
from 
 (select stock_name,sum(price) p from Stocks where operation='Sell' group by stock_name) b
 join (select stock_name,sum(price) p from Stocks where operation='Buy' group by stock_name) a
on a.stock_name=b.stock_name

