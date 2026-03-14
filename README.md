# Shopping-calculator

Accepts input from the user:
    The name of item
    quantity of the item
    Price of the Item 
    They are saved in a data base

From the use of this app, consumer is able to receive suggestions on the names of items to buys and give an average of the price based on the previous purchases

the structure

- server: 
        - main.go
web :
        -index.html
        -style.css
        -web.js
int :- 
        - int: 
                - src : 
                    - calc.go            // it is incharge of calculating
                    - dbpopulator.go    // fills the db
                    - input.go         // process input from the user
                    - suggestion.go   // checks the items that the user usualy purchase 
                - main.go
        - helper

        - api:
                - ItemHandle.go     // hangles the name of the items
                - pricesHandler.go // handle the prices of each iitem
                - nameHandler. go // handles the name of the customer
        - storage:
                - mysql

mostry handled

The app should also be able to suggest the cheapest supper market a consumer can use. checking the shopping of different suppermarket and return the cheapest.


