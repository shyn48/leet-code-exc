select distinct l1.Num As ConsecutiveNums from Logs l1, Logs l2, Logs l3
where
l1.id = l2.id - 1
and l2.id = l3.id - 1
and l1.num = l2.num
and l2.num = l3.num

#if the id's aren't consecutive we can order by the desired column and index the rows using something similar to this:

SELECT @n := @n + 1 n,
       num, 
  FROM table, (SELECT @n := 0) m
  ORDER BY created_at


#second try 
select l1.num as ConsecutiveNums
from Logs l1, Logs l2, Logs l3 where 
l1.id = l2.id - 1
and l2.id = l3.id - 1
and l1.num = l2.num 
and l2.num = l3.num
group by l1.num

/# third try
SELECT distinct num ConsecutiveNums
FROM
(SELECT id, num,
lag(num) over (order by id) as before,
lead(num) over (order by id) as after

FROM logs) next_prev
WHERE num=before and before =after
