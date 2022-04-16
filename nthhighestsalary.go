CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  DECLARE offsetvalue INT;
  SET offsetvalue = N - 1;
  
  RETURN (
    select distinct 
        salary 
    from 
        Employee
    order by salary desc
    limit offsetvalue, 1
  );
END
