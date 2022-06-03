# Write your MySQL query statement below

select name, ifnull(sum(r.distance), 0) as travelled_distance
from Users u left join Rides r on u.id = r.user_id
group by r.user_id
order by travelled_distance desc, name asc

--using coorelated sub-query
select u.name, (select isnull(sum(r.distance),0)
                from rides r
                where u.id = r.user_id) as travelled_distance
from users u
order by travelled_distance desc, name asc

--using cte
with cte
as
(
    select r.user_id, sum(r.distance) as travelled_distance
    from rides r
    group by r.user_id
)
    select u.name, isnull(c.travelled_distance,0) as travelled_distance
    from users u left outer join cte c
    on u.id = c.user_id
    order by travelled_distance desc, name asc
