select user_id as buyer_id, join_date, count(order_id) as orders_in_2019
from Users u left join Orders o on u.user_id = o.buyer_id 
and YEAR(order_date) = '2019'
group by user_id
order by user_id
