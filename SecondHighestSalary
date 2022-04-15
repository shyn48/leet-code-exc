select max(salary)
from  Employee
where salary < (select max(salary) from Employee)

# Alternate solution using subquery:

select 
    (select distinct 
        salary 
    from 
        Employee
    order by salary desc
    limit 1 offset 1) as SecondHighestSalary
