/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal.js");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.pb = (function() {

    /**
     * Namespace pb.
     * @exports pb
     * @namespace
     */
    var pb = {};

    /**
     * EnumSex enum.
     * @name pb.EnumSex
     * @enum {number}
     * @property {number} Unknow=0 Unknow value
     * @property {number} Male=1 Male value
     * @property {number} Female=2 Female value
     */
    pb.EnumSex = (function() {
        var valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "Unknow"] = 0;
        values[valuesById[1] = "Male"] = 1;
        values[valuesById[2] = "Female"] = 2;
        return values;
    })();

    return pb;
})();

module.exports = $root;
