/**
 * @fileoverview gRPC-Web generated client stub for mealey.api.users.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.mealey = {};
proto.mealey.api = {};
proto.mealey.api.users = {};
proto.mealey.api.users.v1 = require('./users_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.mealey.api.users.v1.UsersServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.mealey.api.users.v1.UsersServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.mealey.api.users.v1.CreateUserRequest,
 *   !proto.mealey.api.users.v1.CreateUserResult>}
 */
const methodDescriptor_UsersService_CreateUser = new grpc.web.MethodDescriptor(
  '/mealey.api.users.v1.UsersService/CreateUser',
  grpc.web.MethodType.UNARY,
  proto.mealey.api.users.v1.CreateUserRequest,
  proto.mealey.api.users.v1.CreateUserResult,
  /**
   * @param {!proto.mealey.api.users.v1.CreateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.mealey.api.users.v1.CreateUserResult.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.mealey.api.users.v1.CreateUserRequest,
 *   !proto.mealey.api.users.v1.CreateUserResult>}
 */
const methodInfo_UsersService_CreateUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.mealey.api.users.v1.CreateUserResult,
  /**
   * @param {!proto.mealey.api.users.v1.CreateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.mealey.api.users.v1.CreateUserResult.deserializeBinary
);


/**
 * @param {!proto.mealey.api.users.v1.CreateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.mealey.api.users.v1.CreateUserResult)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.mealey.api.users.v1.CreateUserResult>|undefined}
 *     The XHR Node Readable Stream
 */
proto.mealey.api.users.v1.UsersServiceClient.prototype.createUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/mealey.api.users.v1.UsersService/CreateUser',
      request,
      metadata || {},
      methodDescriptor_UsersService_CreateUser,
      callback);
};


/**
 * @param {!proto.mealey.api.users.v1.CreateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.mealey.api.users.v1.CreateUserResult>}
 *     Promise that resolves to the response
 */
proto.mealey.api.users.v1.UsersServicePromiseClient.prototype.createUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/mealey.api.users.v1.UsersService/CreateUser',
      request,
      metadata || {},
      methodDescriptor_UsersService_CreateUser);
};


module.exports = proto.mealey.api.users.v1;

