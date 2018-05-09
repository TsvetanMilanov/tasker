"use strict";

const constants = require("./constants");

const getEnvVars = () => {
    return Object.keys(process.env)
        .reduce((prev, curr) => {
            if (curr.startsWith(constants.TaskerEnvVarsPrefix)) {
                prev[curr.replace(constants.TaskerEnvVarsPrefix, "")] = process.env[curr];
            }

            return prev;
        }, {});
};

module.exports = () => {
    const common = {
        apiVersion: constants.ApiVersion,
        env: getEnvVars()
    };
    return common;
};
