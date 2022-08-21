# Write your MySQL query statement below

select s.name from SalesPerson s
where s.name not in (
    select s1.name from Orders o join Company c on o.com_id = c.com_id
join SalesPerson s1 on s1.sales_id = o.sales_id 
    where c.name = "RED"
)


# alternate

SELECT
    s.name
FROM
    salesperson s
WHERE
    s.sales_id NOT IN (SELECT
            o.sales_id
        FROM
            orders o
                LEFT JOIN
            company c ON o.com_id = c.com_id
        WHERE
            c.name = 'RED')
;

# Write your MySQL query statement below
select name from SalesPerson s where s.sales_id not in (
    select o.sales_id
        FROM
            orders o
        JOIN
            company c ON o.com_id = c.com_id
        where c.name = 'RED'
) 
