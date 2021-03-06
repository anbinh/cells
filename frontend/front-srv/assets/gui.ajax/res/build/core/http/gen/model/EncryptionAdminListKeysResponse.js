/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */

'use strict';

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

var _EncryptionKey = require('./EncryptionKey');

var _EncryptionKey2 = _interopRequireDefault(_EncryptionKey);

/**
* The EncryptionAdminListKeysResponse model module.
* @module model/EncryptionAdminListKeysResponse
* @version 1.0
*/

var EncryptionAdminListKeysResponse = (function () {
    /**
    * Constructs a new <code>EncryptionAdminListKeysResponse</code>.
    * @alias module:model/EncryptionAdminListKeysResponse
    * @class
    */

    function EncryptionAdminListKeysResponse() {
        _classCallCheck(this, EncryptionAdminListKeysResponse);

        this.Keys = undefined;
    }

    /**
    * Constructs a <code>EncryptionAdminListKeysResponse</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/EncryptionAdminListKeysResponse} obj Optional instance to populate.
    * @return {module:model/EncryptionAdminListKeysResponse} The populated <code>EncryptionAdminListKeysResponse</code> instance.
    */

    EncryptionAdminListKeysResponse.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new EncryptionAdminListKeysResponse();

            if (data.hasOwnProperty('Keys')) {
                obj['Keys'] = _ApiClient2['default'].convertToType(data['Keys'], [_EncryptionKey2['default']]);
            }
        }
        return obj;
    };

    /**
    * @member {Array.<module:model/EncryptionKey>} Keys
    */
    return EncryptionAdminListKeysResponse;
})();

exports['default'] = EncryptionAdminListKeysResponse;
module.exports = exports['default'];
