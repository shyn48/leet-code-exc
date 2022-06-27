select FirstName, LastName, City, State from Person left join Address on Person.PersonId = Address.PersonId;

# pretty easy question

# 2nd run:
select firstName, lastName, city, state
from Person left join Address on Person.personId = Address.personId
