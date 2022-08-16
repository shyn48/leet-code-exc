select distinct t1.*
from stadium t1, stadium t2, stadium t3
where t1.people >= 100 and t2.people >= 100 and t3.people >= 100
and
(
    (t1.id - t2.id = 1 and t1.id - t3.id = 2 and t2.id - t3.id =1) 
    or
    (t2.id - t1.id = 1 and t2.id - t3.id = 2 and t1.id - t3.id =1) 
    or
    (t3.id - t2.id = 1 and t2.id - t1.id =1 and t3.id - t1.id = 2) 
)
order by t1.id

#second try

select distinct r.id, r.visit_date, r.people
from stadium r,
(select a.id as FROM_ID, a.id+2 as TO_ID
from stadium a, stadium b, stadium c
where a.id+1 = b.id
and b.id+1 = c.id
and a.people >= 100
and b.people >= 100
and c.people >= 100) b
where r.id between b.FROM_ID and b.TO_ID
