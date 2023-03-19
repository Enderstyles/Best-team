# Introduction to the service
We are the "Best Team" team developing e-commerce web applications, and so far we have implemented only the basis of the application, that is, only the basic attributes of the service such as: home page, user authorization, item search, post feed. In the future, with the progress of our site, we will gradually draw out the concept of our site.

# Team members description, contributions of each team member
## Medeubekov Sergazy ID:200103154 :
Implemented user registration and login functions, made html pages for them and for the homepage, corrected errors that appeared during development.

## Askhat Al-Aziz ID:200103463 :
Developer and fixer of "Search" and "CreateItem" functions. I have created pages for them and made them mobile adaptable. Fixer of login error.

# Explanation how to run the code:
Start by creating a MySQL database:
Using the XAMPP Control Panel, run the modules: Apache and MySQL, MySQL port must be: 3306
Press the "admin" button on MySQL and go to the phpmyadmin page and there:
Create database "golang_project" and in it create table users with this structure:
![32](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/userstable.png)
Then table items with the following structure:
![33](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/itemstable.png)
For the search page to work properly. You need to install xampp at "C:/" and create the folder "pictures" at "C:/xampp/htdocs/pictures "

Once you have created the database, you can run a terminal in main.go and use "go run main.go" to run the code and go to the localhost page: http://localhost:3000

# Explanation of the code:
## Connect function
 - Connect function is for connectiong to database of out project. Nothing special\
![1](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen001.jpg)
![3](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen003.jpg)

## Register
 - Register function if for register page and in this function contains every process in this page. 
First of all we cheking the type of request of page. We only accepting "Post" type, if its not then we basicly reloading the page then we cheking out form from html page for errors.\
![2](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen002.jpg)

 - Then we taking data from form\
![4](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen004.jpg)

 - Cheking for blank fields, regex and length\
![5](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen005.jpg)
![6](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen006.jpg)

 - Cheking requirement for password\
![7](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen007.jpg)

 - Crypting password and cheking for error\
![8](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen008.jpg)

 - Then finally inserting everything in "Users" table and redirecting to login page\
![9](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen009.jpg)
![10](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen010.jpg)


## The login page
![27](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen027.jpg)

 - In login page first steps is basicly the same, so no need for explanation\
![28](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen028.jpg)

 - Next after all cheking and getting information from form we finding the username which was send by user and if there one, we getting password of this username\
![11](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen011.jpg)

 - Then we cheking if hashed password from the db and one which was send by user are the same and if success we redirecting to homepage\
![12](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen012.jpg)
![13](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen013.jpg)


## The homepage itself
![29](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen029.jpg)
![14](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen014.jpg)


## Search
 - First of all we getting the query itself from form (in home page or in "search" page) then getting the list of items from "search" function\
![15](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen015.jpg)

 - In "search" function first of all we building the query for db by spliting words and making an array of them,  then we connecting them into one string and sending query to our db\
![16](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen016.jpg)

 - After that we getting the result of the query and building the list of items which we will show in html pageafter building the list, we reversing it to show latest items first, then finally returning the list of items\
![17](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen017.jpg)

 - The recieved list of items we sending to "search.html"\
![18](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen018.jpg)


 - In html page we showing the full list of items using {{range .}} and {{end}}\
![19](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen019.jpg)

 - The result of query\
![30](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen030.jpg)


## All items page or Feed
 - In "feed" page all process are the same as in "search" except we showing all items(Yes we know that it could be optimized because its doing the same stuff that "search" does. We will fix that later)\
![20](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen020.jpg)

## Itemmaker page
- The "createItem" function if for making items using html page\
![31](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen031.jpg)
![21](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen021.jpg)

 - After cheking and getting information from form we getting img form it(now I realise there is a mistake. Its about "what if I don't want to upload img. We will fix it later")\
![22](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen022.jpg)

 - Then we copying img to file\
![23](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen023.jpg)

 - Creating file in server directory\
![25](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen025.jpg)\

 - After that we inserting the path of the image to db and if succesful we redirecting to "feed"\
![24](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen024.jpg)

## MAIN
 - And in "main" we connecting to db, creating mux and handling redirections\
![26](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen026.jpg)
