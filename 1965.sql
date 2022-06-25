# Write your MySQL query statement below
select J.employee_id from
(select * from Employees e1 left join Salaries s1 USING(employee_id)
 union
 select * from Employees e2 right join Salaries s2 USING(employee_id)) as J
 where J.name is NULL or J.salary is NULL
 order by J.employee_id
