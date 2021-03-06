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


import ApiClient from '../ApiClient';
import CtlService from './CtlService';





/**
* The RestServiceCollection model module.
* @module model/RestServiceCollection
* @version 1.0
*/
export default class RestServiceCollection {
    /**
    * Constructs a new <code>RestServiceCollection</code>.
    * @alias module:model/RestServiceCollection
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>RestServiceCollection</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/RestServiceCollection} obj Optional instance to populate.
    * @return {module:model/RestServiceCollection} The populated <code>RestServiceCollection</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RestServiceCollection();

            
            
            

            if (data.hasOwnProperty('Services')) {
                obj['Services'] = ApiClient.convertToType(data['Services'], [CtlService]);
            }
            if (data.hasOwnProperty('Total')) {
                obj['Total'] = ApiClient.convertToType(data['Total'], 'Number');
            }
        }
        return obj;
    }

    /**
    * @member {Array.<module:model/CtlService>} Services
    */
    Services = undefined;
    /**
    * @member {Number} Total
    */
    Total = undefined;








}


