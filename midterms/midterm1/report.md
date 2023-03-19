Explanation of the code:

Connect function is for connectiong to database of out project. Nothing special.
![1](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen001.jpg)

![3](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen003.jpg)
Register function if for register page and in this function contains every process in this page.
First of all we cheking the type of request of page. We only accepting "Post" type, 
if its not then we basicly reloading the page
then we cheking out form from html page for errors.
![2](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen002.jpg)

Then we taking data from form
![4](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen004.jpg)

Cheking for blank fields, regex and length

Cheking requirement for password

Crypting password and cheking for error

Then finally inserting everything in "Users" table and redirecting to login page


In login page first steps is basicly the same, so no need for explanation

Next after all cheking and getting information from form 
we finding the username which was send by user and if there one, 
we getting password of this username

Then we cheking if hashed password from the db and one which was send
by user are the same and if success we redirecting to homepage


The homepage itself


Next is "searchitems" function. 
First of all we getting the query itself from form 
then getting the list of items from "search" function



In "search" function first of all we building the query for db
by spliting words and making an array of them, 
then we connecting them into one string and sending query to our db


After that we getting the result of the query and building the list of items 
which we will show in html page
after building the list, we reversing it to show latest items first, 
then finally returning the list of items 

The recieved list of items we sending to "search.html"


In html page we showing the full list of items using {{range .}} and {{end}}

In "feed" page all process are the same as in "search" 
except we showing all items



The "createItem" function if for making items using html page


After cheking and getting information from form we getting img form it
(now I realise there is a mistake. 
Its about "what if I don't want to upload img. We will fix it later")




Then we copying img to file

Creating file in local server directory

After that we inserting the path of the image to db 
and if succesful we redirecting to "feed"



And in "main" we connecting to db, creating mux and handling redirections










