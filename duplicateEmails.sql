SELECT DISTINCT p1.Email
FROM Person p1, Person p2
WHERE p1.Email = p2.Email and p1.id != p2.id

#or

select distinct email from Person p1 where (
    select count(email) from Person p2
    where p1.email = p2.email
) > 1

# or

select Email
from Person
group by Email
having count(Email) > 1;
