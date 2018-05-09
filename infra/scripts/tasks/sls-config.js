"use strict";

const utils = require("../utils");

const TasksTableName = "tasks";

module.exports = () => {
    return {
        tasksTableName: utils.getResourceNameWithStageAndApiVersion(TasksTableName)
    };
};
