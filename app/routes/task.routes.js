module.exports = app => {
    const tasks = require("../controllers/task.controller.js");

    //create new task
    app.post("/task", tasks.create);

    app.put("/task/:taskId", tasks.update);

};