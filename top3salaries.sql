SELECT  d.Name AS Department, e1.Name AS Employee , e1.Salary 
FROM Employee AS e1, Employee as e2, Department AS d
WHERE e.DepartmentId = d.Id
AND e2.DepartmentId = e1.DepartmentId
AND e2.Salary >= e1.Salary 
GROUP BY e1.Id HAVING COUNT(DISTINCT e2.Salary) <= 3;

# faster solution 

select d.Name as 'Department', e1.Name as 'Employee', e1.Salary from Employee e1
join
Department d on e1.DepartmentId = d.id
where (
    select count(distinct e2.Salary)
    from Employee e2
    where e2.Salary > e1.Salary and e1.DepartmentId = e2.DepartmentId
) < 3
