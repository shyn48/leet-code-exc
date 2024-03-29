select customer_number from (
    select count(order_number) as order_count, customer_number from
    orders
    group by customer_number
    order by order_count desc
    limit 1
) as max_orders


# better solution

SELECT
    customer_number
FROM
    orders
GROUP BY customer_number
ORDER BY COUNT(*) DESC
LIMIT 1
;

#second try

select customer_number from orders
group by customer_number
having count(order_number) = (
    select max(count) from ( 
        select count(order_number) as count, customer_number from orders group by customer_number
    ) max_count
)
