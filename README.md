# appointyapi

This is instagram api build using Golang. 
<h2>Functions:</h2>
<ul>
<li>Create users</li>
<li>Get users by id</li>
<li>Create Posts</li>
<li>Get posts by id</li>
<li>Get all posts of the user</li>
</ul>
<h3>Create users<h3>
This function creates a user with id,name,email,password in database on POST request. endpoint:"/users".
<hr>
<h3>Get users by ids<h3>
This function retrives a user data in the database with id on GET request.  endpoint:"/users/id".
<hr>
<h3>Create Posts<h3>
This function creates a post with id,userid,caption,imageurl,postedtimestamp in database on POST request. endpoint:"/posts".
<hr>
<h3>Get posts by id<h3>
This function retrives a post with id on GET request. endpoint:"/posts/id".
<hr>
<h3>Get all posts of the user<h3>
This function retrives all the posts with userid in database on GET request. endpoint:"/posts/userid".
