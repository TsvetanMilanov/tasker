"use strict";

const constants = require("./constants");

const getOptions = () => {
    let result = {};

    for (let i = 0; i < process.argv.length; i++) {
        const currentOption = process.argv[i];
        if (currentOption.startsWith("--")) {
            const optionName = currentOption.substr(2);
            let optionValue = process.argv[i + 1];

            if (optionValue === undefined || optionValue.startsWith("--")) {
                optionValue = true;
            }

            result[optionName] = optionValue;
        }
    }

    if (!result.stage) {
        throw new Error("Please provide --stage option.");
    }

    return result;
};

const getResourceNameWithStageAndApiVersion = (resource) => {
    const opts = getOptions();
    return `${opts.stage}-${resource}-${constants.ApiVersion}`;
};

module.exports = {
    getOptions,
    getResourceNameWithStageAndApiVersion
};
