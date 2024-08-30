select name from SalesPerson s where s.sales_id not in (select sales_id from Company c, Orders o where o.com_id=c.com_id and c.name='RED');
