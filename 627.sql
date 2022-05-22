# Write your MySQL query statement below
Update salary set 
sex = CASE sex when 'm' then 'f' else 'm'
end;
