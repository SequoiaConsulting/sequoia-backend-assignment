const sql = require("./db.js");

const Task = function (task) {
    this.title = task.title;
    this.status = task.status; //admin, user
    this.description = task.description;
    this.due_date = task.due_date;
    this.project_id = task.project_id;
    this.assignee_email = task.assignee_email;
    this.assigner_email = task.assigner_email;
};

Task.create = (newTask,result) => {
    sql.query("INSERT INTO tasks SET ?",newTask, (err, res) => {
        if (err) {
            console.log("error: ", err);
            result(err, null);
            return;
        }

        console.log("created task: ", { id: res.insertId, ...newTask});
        result(null, { id: res.insertId, ...newTask });
    });
};

Task.updateById = (id, task, result) => {
    sql.query(
        "UPDATE tasks SET title = ?, status = ?, description = ?, assignee_email = ? WHERE id = ?",
        [task.title, task.status, task.description, task.assignee_email ,id],
        (err, res) => {
            if (err) {
                console.log("error: ", err);
                result(null, err);
                return;
            }

            if (res.affectedRows == 0) {
                // not found Customer with the id
                result({ kind: "not_found" }, null);
                return;
            }

            console.log("updated task: ", { id: id, ...task });
            result(null, { id: id, ...task });
        }
    );
};
module.exports = Task;