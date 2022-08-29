select actor_id, director_id from ActorDirector
group by actor_id, director_id 
having count(*) >= 3 

#second try
select actor_id, director_id 
from ActorDirector 
group by actor_id, director_id
having count(*) >= 3
