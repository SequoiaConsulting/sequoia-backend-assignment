# Trello Board Application

### Built a "headless" simple Trello board application to manage users and their tasks.	

#### Assumptions (Developed somewhat similar to real Trello/JIRA Application)
 - Application admin and Project admin are two different terms. Application admin is the one who is owner/developer/CTO of the whole application/company. Project admin is any authorized user who has created the project board.
 - Application admin has all the permissions to all the API endpoints.
 - User can be uniquely identified by username or email id.
 - All requests must be authorised via token except login/signup requests. This token can be retrieved by requesting on login API endpoint by sending credentials in the request.
 - Project admin can deactivate any user(project admin/non-admin users) from the project except himself/herself.
 - Only project admin can update the project board's details. 

 #### Features (Implemented via API endpoints)
  - There can be mulitple project admins. Only admin can make other user as an admin of that project.
  - Once an user is deactivated it will automatically become a non-admin as well (if user was admin earlier).
  - At any time there will always be atleat one admin of a project board.
  - Any number of statuses can be added for a particular project board.
  - User(admin/non-admin) associated to n projects can view all users of those n projects. Can also view users of a specific project.
  - Application Admin can view all users of all projects whether associated with project or not. Can also view it from Admin Panel.
  - To add/retrieve/edit/remove a task, one must be associated to the task's project board.

#### Other Implementations
 - Generated [API documentation](https://www.getpostman.com/collections/676f38b3884d762d4f9d) via Postman Collection
 - Done with API testing.
 - Tried dockerizing the code. Made Docker Files but the code is not dockerized due to few issues.
 - Built admin panel using Django admin panel. Checked all the functionalities via admin panel.
 - Readme describes how to build and run the code.


#### This Application has three entities:

  - **Users**: Users of the application. 
    - [x] Each user is uniquely identified by his/her email address.
    - [x] There will be two roles: Admin and User. 
      - [x] Admin will be able to create/archive/unarchive Project Boards
      - [x] Admin will be able to view all users of all project boards
      - [x] Admin will be able to add new user or deactivate existing user from Project Board
      - [x] Non-Admin users will be able to view all the users of Project Boards that they're assigned to.
				
  - **Project Boards**: 
   - [x] Represents a Project which will have tasks defined under it. A non-admin user needs to be assigned to a particular Project Board in order to view it. 
		
  - **Tasks**: An atomic entity that defines the objective. 
    - [x] A Task is assignable to/by a user of the system. 
    - [x] A Task belongs to a particular Project Board.
    - [x] A Task should have minimum of the following properties: Title, Description, Assignee, Assigner, Due Date, Status.
    - [x] A Task can be added/removed/edited
    - [x] A Task can be in a particular Status. For ex: "Backlog", "InProgress", "Done".(Optional Task)
    - [x] User can create any number of new statuses or remove existing Status in a Project Board and once defined, A Task can be assigned to one of these statuses.(Optional Task)
	


#### Technology stack with which you have to write with:
  - All the above CRUD operations should be exposed as REST endpoints
  - _Programming language_: 
    - [ ] NodeJS
	- [ ] Golang
	- [x] Python
  - _Database_: 
    - [ ] MySQL
	- [x] Postgresql
	- [ ] MariaDB
	- [ ] SQLite
	- [ ] MS SQL
  - _Framework_: We don’t care as long it serves the above three pointers. We would love to see if you can implement in Erlang or Rust. 


##  We are not looking for UI implmentation of the board 

#### Our expectations when you say your code is ready:
  - [x] Write all the APIs visualizing if there is a Trello board UI for this. We are not expecting any UI components
  - [x] Quality code standards. Hope you have done lint check before pushing the code. No one likes to hear, "I didn't have time so..".
  - [x] Apt input validations. Think from the end user perspective(username contains only alphanumerics, email satisfies the standard regex pattern, etc.,). They say "Something went wrong, please try again" shows laziness of a developer, so don't be one. More your code breaks, more we lose trust on your quality.
  - [x] (Optional Task) Tests. Be it Unit or Integration or API tests. At least one of them because only your tests can assure your code is working. 
  - Update documentation in README.md file for us which should have the following
		○ [x] How to build and run your code
		○ [x] What are the assumptions you have made during development
		○ [x] Checkmark these expectations when you have finished them
	

#### We'd be really impressed if you include at-least one of the following below along with fulfilling our above expectations:
  - [ ] Dockerize your code.
  - [x] Generate Open API documentation using Swagger or related(Postman collection, etc.,)
  - [x] Covered 99.99% possible cases without errors and introduce new use-cases wherever "necessary" -- True traits of 10x developer :P
  - [x] Built admin panel UI for it - Django Admin Panel.
  


#### Interested enough? Steps to go ahead about this assignment:
  - Fork this Github repository. [Help link](https://guides.github.com/activities/forking) if needed
  - Keep committing your changes on this forked repository regularly, we prefer if you are comitting several small changes instead of one large commit. Dont worry, only your final commit will be considered for evaluation.
  - Make sure you keep editing this README file on your forked repository and mark checkboxes above on the things you completed ([Help link](https://www.markdownguide.org/extended-syntax/#task-lists) to mark a checkbox in this README markdown)
  - Once you have finalised, create a Pull Request to this original repository. We'll review it and get back to you with some news.
  

 
_In case of any queries, mail to Karthikeyan NG <karthikeyan.ng@sequoia.com> or Indrajeet Kumar <indrajeet@sequoia.com>. We'll revert to you with the clarifications_
 
 

