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

var _DocstoreDocument = require('./DocstoreDocument');

var _DocstoreDocument2 = _interopRequireDefault(_DocstoreDocument);

/**
* The DocstorePutDocumentRequest model module.
* @module model/DocstorePutDocumentRequest
* @version 1.0
*/

var DocstorePutDocumentRequest = (function () {
    /**
    * Constructs a new <code>DocstorePutDocumentRequest</code>.
    * @alias module:model/DocstorePutDocumentRequest
    * @class
    */

    function DocstorePutDocumentRequest() {
        _classCallCheck(this, DocstorePutDocumentRequest);

        this.StoreID = undefined;
        this.DocumentID = undefined;
        this.Document = undefined;
    }

    /**
    * Constructs a <code>DocstorePutDocumentRequest</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/DocstorePutDocumentRequest} obj Optional instance to populate.
    * @return {module:model/DocstorePutDocumentRequest} The populated <code>DocstorePutDocumentRequest</code> instance.
    */

    DocstorePutDocumentRequest.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DocstorePutDocumentRequest();

            if (data.hasOwnProperty('StoreID')) {
                obj['StoreID'] = _ApiClient2['default'].convertToType(data['StoreID'], 'String');
            }
            if (data.hasOwnProperty('DocumentID')) {
                obj['DocumentID'] = _ApiClient2['default'].convertToType(data['DocumentID'], 'String');
            }
            if (data.hasOwnProperty('Document')) {
                obj['Document'] = _DocstoreDocument2['default'].constructFromObject(data['Document']);
            }
        }
        return obj;
    };

    /**
    * @member {String} StoreID
    */
    return DocstorePutDocumentRequest;
})();

exports['default'] = DocstorePutDocumentRequest;
module.exports = exports['default'];

/**
* @member {String} DocumentID
*/

/**
* @member {module:model/DocstoreDocument} Document
*/