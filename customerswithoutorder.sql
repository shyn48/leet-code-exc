select name as 'Customers' from Customers 
where Customers.id not in (
    select customerid from orders
);

#SECOND TRY
select name AS 'Customers'
from Customers c left join Orders o
on c.id = o.customerId
where o.customerId is null
