# Introduction to the service
We are the "Best Team" developing a web application for e-commerce written in GoLang language for buying and selling various types of goods. On our web application, you can place your product in our feed, as well as give reviews and ratings to other products. Here you can find and sell products that you like!

# Team members description, contributions of each team member
## Medeubekov Sergazy ID:200103154 :
- Created a user rating system, now users can evaluate the product and the average rating of the product from all users will be visible.
- Created filtering by product rating, users can search for products depending on their rating
- Created a user profile page, after logging in, the user now sees his profile and can log out of it

## Askhat Al-Aziz ID:200103463 :
- Creator of filtering py price
- Creator of ItemPage where you can see detailed information about item
- Creator of Comments System
- Debugger and fixer of the project

# Some preparetions: Explanation how to run the code
Find "golang_project.sql" and run it on your D B
!!! We are using xampp and MySql !!!
For the search page to work properly. You need to install xampp at "C:/" and create the folder "pictures" at "C:/xampp/htdocs/pictures "
Once you have created the database, you can run a terminal in main.go and use "go run main.go" to run the code and go to the localhost page: http://localhost:3000


## MinMax function
 - Getting items from form
 - Searching in DB items hich is between this values
 - Getting items
 - Getting tags 
 - Parsing data to template
![1](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen032.png)
![2](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen036.png)

## Rate item
 - Getting item id from Mux's variables
 - Cheching for authorization
 - Checking the validity of rating
 - Insetring data to DB
 - Updating the Avg of rating
 - redirecting
![3](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen034.png)
![4](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen035.png)
![5](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen040.png)

## Item description
 - Gettign item id from page
 - Finding this item in DB
 - Initalizing item variable
 - Getting comments related to this item
 - Reversing order parsing item and comments 
 - Executing 
![6](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen037.png)
![7](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen038.png)
![8](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen036.png)

## Commenting 
 - Getting userId to find username in DB
 - Inserting comment to DB
![9](https://github.com/Enderstyles/Best-team/blob/main/data/pictures/screen039.png)
