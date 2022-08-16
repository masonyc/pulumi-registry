"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
exports.__esModule = true;
exports.resourceOptsDefaults = exports.getVersion = exports.getEnvNumber = exports.getEnvBoolean = exports.getEnv = void 0;
function getEnv() {
    var vars = [];
    for (var _i = 0; _i < arguments.length; _i++) {
        vars[_i] = arguments[_i];
    }
    for (var _a = 0, vars_1 = vars; _a < vars_1.length; _a++) {
        var v = vars_1[_a];
        var value = process.env[v];
        if (value) {
            return value;
        }
    }
    return undefined;
}
exports.getEnv = getEnv;
function getEnvBoolean() {
    var vars = [];
    for (var _i = 0; _i < arguments.length; _i++) {
        vars[_i] = arguments[_i];
    }
    var s = getEnv.apply(void 0, vars);
    if (s !== undefined) {
        // NOTE: these values are taken from https://golang.org/src/strconv/atob.go?s=351:391#L1, which is what
        // Terraform uses internally when parsing boolean values.
        if (["1", "t", "T", "true", "TRUE", "True"].find(function (v) { return v === s; }) !== undefined) {
            return true;
        }
        if (["0", "f", "F", "false", "FALSE", "False"].find(function (v) { return v === s; }) !== undefined) {
            return false;
        }
    }
    return undefined;
}
exports.getEnvBoolean = getEnvBoolean;
function getEnvNumber() {
    var vars = [];
    for (var _i = 0; _i < arguments.length; _i++) {
        vars[_i] = arguments[_i];
    }
    var s = getEnv.apply(void 0, vars);
    if (s !== undefined) {
        var f = parseFloat(s);
        if (!isNaN(f)) {
            return f;
        }
    }
    return undefined;
}
exports.getEnvNumber = getEnvNumber;
function getVersion() {
    var version = require('./package.json').version;
    // Node allows for the version to be prefixed by a "v", while semver doesn't.
    // If there is a v, strip it off.
    if (version.indexOf('v') === 0) {
        version = version.slice(1);
    }
    return version;
}
exports.getVersion = getVersion;
/** @internal */
function resourceOptsDefaults() {
    return { version: getVersion() };
}
exports.resourceOptsDefaults = resourceOptsDefaults;
