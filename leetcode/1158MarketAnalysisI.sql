# Write your MySQL query statement below
select user_id as buyer_id, join_date, IFNULL(a.orders_in_2019,0) orders_in_2019 from Users left join (select count(*) as orders_in_2019, buyer_id from Orders where order_date >= '2019-01-01' && order_date <= '2019-12-31' group by buyer_id) a on Users.user_id=a.buyer_id
