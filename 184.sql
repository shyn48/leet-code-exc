select d.Name as Department, e.Name as Employee, e.salary
from Employee e join Department as d on e.departmentId = d.Id
    join (
        select max(salary) as Salary, departmentId
        from Employee
        group by departmentId
    ) as maxsalary
    on e.Salary = maxsalary.Salary and e.departmentId = maxsalary.departmentId
