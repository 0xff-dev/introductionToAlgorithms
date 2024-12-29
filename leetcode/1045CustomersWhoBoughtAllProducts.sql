select a.customer_id from 
(select customer_id, count(distinct(product_key)) c from Customer group by customer_id) a,
(select count(distinct(product_key)) c from Product ) b where a.c = b.c;
