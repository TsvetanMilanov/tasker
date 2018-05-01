"use strict";

const ApiVersion = "v1";
const TaskerEnvVarsPrefix = "TASKER_";

const getEnvVars = () => {
    return Object.keys(process.env)
        .reduce((prev, curr) => {
            if (curr.startsWith(TaskerEnvVarsPrefix)) {
                prev[curr.replace(TaskerEnvVarsPrefix, "")] = process.env[curr];
            }

            return prev;
        }, {});
};

module.exports = () => {
    const common = {
        apiVersion: ApiVersion,
        env: getEnvVars()
    };

    return common;
};
