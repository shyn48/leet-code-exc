select w1.id as 'Id' from Weather w1 join Weather w2 
where DATEDIFF(w1.recordDate, w2.recordDate) = 1
and w1.temperature > w2.temperature 

#second try
select w2.id
from Weather w1, Weather w2 
where datediff(w2.recordDate, w1.recordDate) = 1 and w2.temperature > w1.temperature
