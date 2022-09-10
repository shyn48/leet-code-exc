select stock_name, SUM(
    if(operation = 'Buy', -price, price)
) as capital_gain_loss
from Stocks group by stock_name

select 
stock_name, 
sum(case when operation = 'Buy'
then -price
else price
end) as capital_gain_loss
from Stocks group by stock_name
