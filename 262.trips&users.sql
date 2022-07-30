select request_at as day,
round(sum(case when status != 'completed' then 1 else 0 end)/count(*), 2) 
as 'cancellation rate' 
from trips t join users rider on t.client_id = rider.users_id 
join users driver on t.driver_id = driver.users_id
where t.request_at between '2013-10-01' and '2013-10-03'
and rider.banned = 'No'
and driver.banned = 'No'
group by request_at

#second try

select request_at as 'Day',
round(sum(case when status != 'completed' then 1 else 0 end)/ count(*), 2) as 'Cancellation Rate'
from trips t join users c on t.client_id = c.users_id
join users d on t.driver_id = d.users_id 
where c.banned = 'No' 
and
d.banned = 'No'
and t.request_at between '2013-10-01' and '2013-10-03'
group by request_at
