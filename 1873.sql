select 
    employee_id,
    case 
        when employee_id%2=1 and SUBSTRING(name, 1,1) != 'M' then salary
        else 0
    end as bonus
from Employees
order by 
    Employee_id ASC

1873. select 
    employee_id,
    case 
        when employee_id%2=1 and name not like 'M%' then salary
        else 0
    end as bonus
from Employees
order by 
    Employee_id ASC

