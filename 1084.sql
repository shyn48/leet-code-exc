# Write your MySQL query statement below

select product_id, product_name from Product 
where product_id in (
       select product_id
    from Sales
    group by product_id
    HAVING MIN(sale_date) >= '2019-01-01' AND MAX(sale_date) <= '2019-03-31'
)

#second try
select s.product_id, product_name 
from Sales s join Product p on s.product_id = p.product_id
group by s.product_id having min(s.sale_date) >= '2019-01-01' 
and max(s.sale_date) <= '2019-03-31'
