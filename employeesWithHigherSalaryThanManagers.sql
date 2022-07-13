select e1.name as 'Employee' from Employee e1 join Employee e2  on e1.managerId = e2.id where e1.salary > e2.salary

#second try

select e1.name as 'Employee'
from Employee e1, Employee e2 
    where e1.salary > e2.salary and e1.managerId = e2.Id
    
#faster solution
select a.Name
from Employee a inner join Employee b on a.ManagerId=b.Id
where a.Salary>b.Salary
