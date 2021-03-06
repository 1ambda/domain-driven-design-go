/// <reference path="./custom.d.ts" />
// tslint:disable
/**
 * Gateway
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


import * as url from "url";
import * as portableFetch from "portable-fetch";
import { Configuration } from "./configuration";

const BASE_PATH = "http://localhost/api".replace(/\/+$/, "");

/**
 *
 * @export
 */
export const COLLECTION_FORMATS = {
    csv: ",",
    ssv: " ",
    tsv: "\t",
    pipes: "|",
};

/**
 *
 * @export
 * @interface FetchAPI
 */
export interface FetchAPI {
    (url: string, init?: any): Promise<Response>;
}

/**
 *  
 * @export
 * @interface FetchArgs
 */
export interface FetchArgs {
    url: string;
    options: any;
}

/**
 * 
 * @export
 * @class BaseAPI
 */
export class BaseAPI {
    protected configuration: Configuration;

    constructor(configuration?: Configuration, protected basePath: string = BASE_PATH, protected fetch: FetchAPI = portableFetch) {
        if (configuration) {
            this.configuration = configuration;
            this.basePath = configuration.basePath || this.basePath;
        }
    }
};

/**
 * 
 * @export
 * @class RequiredError
 * @extends {Error}
 */
export class RequiredError extends Error {
    name: "RequiredError"
    constructor(public field: string, msg?: string) {
        super(msg);
    }
}

/**
 * 
 * @export
 * @interface AuthResponse
 */
export interface AuthResponse {
    /**
     * 
     * @type {string}
     * @memberof AuthResponse
     */
    uid?: string;
}

/**
 * 
 * @export
 * @interface Empty
 */
export interface Empty {
}

/**
 * 
 * @export
 * @interface Exception
 */
export interface Exception {
    /**
     * 
     * @type {string}
     * @memberof Exception
     */
    timestamp?: string;
    /**
     * 
     * @type {number}
     * @memberof Exception
     */
    code?: number;
    /**
     * 
     * @type {string}
     * @memberof Exception
     */
    message?: string;
    /**
     * 
     * @type {string}
     * @memberof Exception
     */
    type?: Exception.TypeEnum;
}

/**
 * @export
 * @namespace Exception
 */
export namespace Exception {
    /**
     * @export
     * @enum {string}
     */
    export enum TypeEnum {
        BadRequest = <any> 'BadRequest',
        Unauthorized = <any> 'Unauthorized',
        Forbidden = <any> 'Forbidden',
        NotFound = <any> 'NotFound',
        InternalServer = <any> 'InternalServer'
    }
}

/**
 * 
 * @export
 * @interface InlineResponse200
 */
export interface InlineResponse200 {
    /**
     * 
     * @type {Array&lt;Product&gt;}
     * @memberof InlineResponse200
     */
    rows?: Array<Product>;
    /**
     * 
     * @type {Pagination}
     * @memberof InlineResponse200
     */
    pagination?: Pagination;
}

/**
 * 
 * @export
 * @interface InlineResponse2001
 */
export interface InlineResponse2001 {
    /**
     * 
     * @type {Product}
     * @memberof InlineResponse2001
     */
    product?: Product;
    /**
     * 
     * @type {Array&lt;ProductOption&gt;}
     * @memberof InlineResponse2001
     */
    options?: Array<ProductOption>;
}

/**
 * 
 * @export
 * @interface LoginRequest
 */
export interface LoginRequest {
    /**
     * 
     * @type {string}
     * @memberof LoginRequest
     */
    uid?: string;
    /**
     * 
     * @type {string}
     * @memberof LoginRequest
     */
    password?: string;
}

/**
 * 
 * @export
 * @interface Pagination
 */
export interface Pagination {
    /**
     * 
     * @type {number}
     * @memberof Pagination
     */
    itemCountPerPage: number;
    /**
     * 
     * @type {number}
     * @memberof Pagination
     */
    currentPageOffset: number;
    /**
     * 
     * @type {number}
     * @memberof Pagination
     */
    totalItemCount: number;
}

/**
 * 
 * @export
 * @interface Product
 */
export interface Product {
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    price?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    onSale?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    categoryID?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    categoryDisplayName?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    categoryPath?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    imageID?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    imageType?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    imagePath?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    createdAt?: string;
    /**
     * 
     * @type {string}
     * @memberof Product
     */
    updatedAt?: string;
}

/**
 * 
 * @export
 * @interface ProductOption
 */
export interface ProductOption {
    /**
     * 
     * @type {string}
     * @memberof ProductOption
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof ProductOption
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof ProductOption
     */
    price?: string;
    /**
     * 
     * @type {string}
     * @memberof ProductOption
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof ProductOption
     */
    onSale?: string;
}

/**
 * 
 * @export
 * @interface RegisterRequest
 */
export interface RegisterRequest {
    /**
     * 
     * @type {string}
     * @memberof RegisterRequest
     */
    uid?: string;
    /**
     * 
     * @type {string}
     * @memberof RegisterRequest
     */
    email?: string;
    /**
     * 
     * @type {string}
     * @memberof RegisterRequest
     */
    password?: string;
}


/**
 * AuthApi - fetch parameter creator
 * @export
 */
export const AuthApiFetchParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @param {LoginRequest} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        login(body?: LoginRequest, options: any = {}): FetchArgs {
            const localVarPath = `/auth/login`;
            const localVarUrlObj = url.parse(localVarPath, true);
            const localVarRequestOptions = Object.assign({ method: 'POST' }, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            localVarHeaderParameter['Content-Type'] = 'application/json';

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);
            const needsSerialization = (<any>"LoginRequest" !== "string") || localVarRequestOptions.headers['Content-Type'] === 'application/json';
            localVarRequestOptions.body =  needsSerialization ? JSON.stringify(body || {}) : (body || "");

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {Empty} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        logout(body?: Empty, options: any = {}): FetchArgs {
            const localVarPath = `/auth/logout`;
            const localVarUrlObj = url.parse(localVarPath, true);
            const localVarRequestOptions = Object.assign({ method: 'POST' }, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            localVarHeaderParameter['Content-Type'] = 'application/json';

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);
            const needsSerialization = (<any>"Empty" !== "string") || localVarRequestOptions.headers['Content-Type'] === 'application/json';
            localVarRequestOptions.body =  needsSerialization ? JSON.stringify(body || {}) : (body || "");

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {RegisterRequest} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        register(body?: RegisterRequest, options: any = {}): FetchArgs {
            const localVarPath = `/auth/register`;
            const localVarUrlObj = url.parse(localVarPath, true);
            const localVarRequestOptions = Object.assign({ method: 'POST' }, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            localVarHeaderParameter['Content-Type'] = 'application/json';

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);
            const needsSerialization = (<any>"RegisterRequest" !== "string") || localVarRequestOptions.headers['Content-Type'] === 'application/json';
            localVarRequestOptions.body =  needsSerialization ? JSON.stringify(body || {}) : (body || "");

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        whoami(options: any = {}): FetchArgs {
            const localVarPath = `/auth/whoami`;
            const localVarUrlObj = url.parse(localVarPath, true);
            const localVarRequestOptions = Object.assign({ method: 'GET' }, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * AuthApi - functional programming interface
 * @export
 */
export const AuthApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @param {LoginRequest} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        login(body?: LoginRequest, options?: any): (fetch?: FetchAPI, basePath?: string) => Promise<AuthResponse> {
            const localVarFetchArgs = AuthApiFetchParamCreator(configuration).login(body, options);
            return (fetch: FetchAPI = portableFetch, basePath: string = BASE_PATH) => {
                return fetch(basePath + localVarFetchArgs.url, localVarFetchArgs.options).then((response) => {
                    if (response.status >= 200 && response.status < 300) {
                        return response.json();
                    } else {
                        throw response;
                    }
                });
            };
        },
        /**
         * 
         * @param {Empty} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        logout(body?: Empty, options?: any): (fetch?: FetchAPI, basePath?: string) => Promise<Empty> {
            const localVarFetchArgs = AuthApiFetchParamCreator(configuration).logout(body, options);
            return (fetch: FetchAPI = portableFetch, basePath: string = BASE_PATH) => {
                return fetch(basePath + localVarFetchArgs.url, localVarFetchArgs.options).then((response) => {
                    if (response.status >= 200 && response.status < 300) {
                        return response.json();
                    } else {
                        throw response;
                    }
                });
            };
        },
        /**
         * 
         * @param {RegisterRequest} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        register(body?: RegisterRequest, options?: any): (fetch?: FetchAPI, basePath?: string) => Promise<AuthResponse> {
            const localVarFetchArgs = AuthApiFetchParamCreator(configuration).register(body, options);
            return (fetch: FetchAPI = portableFetch, basePath: string = BASE_PATH) => {
                return fetch(basePath + localVarFetchArgs.url, localVarFetchArgs.options).then((response) => {
                    if (response.status >= 200 && response.status < 300) {
                        return response.json();
                    } else {
                        throw response;
                    }
                });
            };
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        whoami(options?: any): (fetch?: FetchAPI, basePath?: string) => Promise<AuthResponse> {
            const localVarFetchArgs = AuthApiFetchParamCreator(configuration).whoami(options);
            return (fetch: FetchAPI = portableFetch, basePath: string = BASE_PATH) => {
                return fetch(basePath + localVarFetchArgs.url, localVarFetchArgs.options).then((response) => {
                    if (response.status >= 200 && response.status < 300) {
                        return response.json();
                    } else {
                        throw response;
                    }
                });
            };
        },
    }
};

/**
 * AuthApi - factory interface
 * @export
 */
export const AuthApiFactory = function (configuration?: Configuration, fetch?: FetchAPI, basePath?: string) {
    return {
        /**
         * 
         * @param {LoginRequest} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        login(body?: LoginRequest, options?: any) {
            return AuthApiFp(configuration).login(body, options)(fetch, basePath);
        },
        /**
         * 
         * @param {Empty} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        logout(body?: Empty, options?: any) {
            return AuthApiFp(configuration).logout(body, options)(fetch, basePath);
        },
        /**
         * 
         * @param {RegisterRequest} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        register(body?: RegisterRequest, options?: any) {
            return AuthApiFp(configuration).register(body, options)(fetch, basePath);
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        whoami(options?: any) {
            return AuthApiFp(configuration).whoami(options)(fetch, basePath);
        },
    };
};

/**
 * AuthApi - object-oriented interface
 * @export
 * @class AuthApi
 * @extends {BaseAPI}
 */
export class AuthApi extends BaseAPI {
    /**
     * 
     * @param {} [body] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AuthApi
     */
    public login(body?: LoginRequest, options?: any) {
        return AuthApiFp(this.configuration).login(body, options)(this.fetch, this.basePath);
    }

    /**
     * 
     * @param {} [body] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AuthApi
     */
    public logout(body?: Empty, options?: any) {
        return AuthApiFp(this.configuration).logout(body, options)(this.fetch, this.basePath);
    }

    /**
     * 
     * @param {} [body] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AuthApi
     */
    public register(body?: RegisterRequest, options?: any) {
        return AuthApiFp(this.configuration).register(body, options)(this.fetch, this.basePath);
    }

    /**
     * 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AuthApi
     */
    public whoami(options?: any) {
        return AuthApiFp(this.configuration).whoami(options)(this.fetch, this.basePath);
    }

}

/**
 * ProductApi - fetch parameter creator
 * @export
 */
export const ProductApiFetchParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @param {number} [itemCountPerPage] 
         * @param {number} [currentPageOffset] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        findAll(itemCountPerPage?: number, currentPageOffset?: number, options: any = {}): FetchArgs {
            const localVarPath = `/product`;
            const localVarUrlObj = url.parse(localVarPath, true);
            const localVarRequestOptions = Object.assign({ method: 'GET' }, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            if (itemCountPerPage !== undefined) {
                localVarQueryParameter['itemCountPerPage'] = itemCountPerPage;
            }

            if (currentPageOffset !== undefined) {
                localVarQueryParameter['currentPageOffset'] = currentPageOffset;
            }

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} productID 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        findOneWithOptions(productID: string, options: any = {}): FetchArgs {
            // verify required parameter 'productID' is not null or undefined
            if (productID === null || productID === undefined) {
                throw new RequiredError('productID','Required parameter productID was null or undefined when calling findOneWithOptions.');
            }
            const localVarPath = `/product/{productID}`
                .replace(`{${"productID"}}`, encodeURIComponent(String(productID)));
            const localVarUrlObj = url.parse(localVarPath, true);
            const localVarRequestOptions = Object.assign({ method: 'GET' }, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * ProductApi - functional programming interface
 * @export
 */
export const ProductApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @param {number} [itemCountPerPage] 
         * @param {number} [currentPageOffset] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        findAll(itemCountPerPage?: number, currentPageOffset?: number, options?: any): (fetch?: FetchAPI, basePath?: string) => Promise<InlineResponse200> {
            const localVarFetchArgs = ProductApiFetchParamCreator(configuration).findAll(itemCountPerPage, currentPageOffset, options);
            return (fetch: FetchAPI = portableFetch, basePath: string = BASE_PATH) => {
                return fetch(basePath + localVarFetchArgs.url, localVarFetchArgs.options).then((response) => {
                    if (response.status >= 200 && response.status < 300) {
                        return response.json();
                    } else {
                        throw response;
                    }
                });
            };
        },
        /**
         * 
         * @param {string} productID 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        findOneWithOptions(productID: string, options?: any): (fetch?: FetchAPI, basePath?: string) => Promise<InlineResponse2001> {
            const localVarFetchArgs = ProductApiFetchParamCreator(configuration).findOneWithOptions(productID, options);
            return (fetch: FetchAPI = portableFetch, basePath: string = BASE_PATH) => {
                return fetch(basePath + localVarFetchArgs.url, localVarFetchArgs.options).then((response) => {
                    if (response.status >= 200 && response.status < 300) {
                        return response.json();
                    } else {
                        throw response;
                    }
                });
            };
        },
    }
};

/**
 * ProductApi - factory interface
 * @export
 */
export const ProductApiFactory = function (configuration?: Configuration, fetch?: FetchAPI, basePath?: string) {
    return {
        /**
         * 
         * @param {number} [itemCountPerPage] 
         * @param {number} [currentPageOffset] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        findAll(itemCountPerPage?: number, currentPageOffset?: number, options?: any) {
            return ProductApiFp(configuration).findAll(itemCountPerPage, currentPageOffset, options)(fetch, basePath);
        },
        /**
         * 
         * @param {string} productID 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        findOneWithOptions(productID: string, options?: any) {
            return ProductApiFp(configuration).findOneWithOptions(productID, options)(fetch, basePath);
        },
    };
};

/**
 * ProductApi - object-oriented interface
 * @export
 * @class ProductApi
 * @extends {BaseAPI}
 */
export class ProductApi extends BaseAPI {
    /**
     * 
     * @param {} [itemCountPerPage] 
     * @param {} [currentPageOffset] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ProductApi
     */
    public findAll(itemCountPerPage?: number, currentPageOffset?: number, options?: any) {
        return ProductApiFp(this.configuration).findAll(itemCountPerPage, currentPageOffset, options)(this.fetch, this.basePath);
    }

    /**
     * 
     * @param {} productID 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ProductApi
     */
    public findOneWithOptions(productID: string, options?: any) {
        return ProductApiFp(this.configuration).findOneWithOptions(productID, options)(this.fetch, this.basePath);
    }

}

