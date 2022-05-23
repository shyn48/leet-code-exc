# Write your MySQL query statement below

select product_id, product_name from Product 
where product_id in (
       select product_id
    from Sales
    group by product_id
    HAVING MIN(sale_date) >= '2019-01-01' AND MAX(sale_date) <= '2019-03-31'
)
