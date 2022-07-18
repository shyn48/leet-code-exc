select d.Name as Department, e.Name as Employee, e.salary
from Employee e join Department as d on e.departmentId = d.Id
    join (
        select max(salary) as Salary, departmentId
        from Employee
        group by departmentId
    ) as maxsalary
    on e.Salary = maxsalary.Salary and e.departmentId = maxsalary.departmentId


#second try

select d.name as "Department", e.name as "Employee", e.Salary
from Employee e 
join Department d on e.departmentId = d.id 
where (e.DepartmentId, e.Salary) in (
    select d1.id, max(salary) from  Employee e1
    join Department d1 on e1.departmentId = d1.id
    group by d1.id
)
