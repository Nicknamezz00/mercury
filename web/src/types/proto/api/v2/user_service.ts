/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../../google/protobuf/timestamp";
import { RowStatus, rowStatusFromJSON, rowStatusToJSON } from "./common";

export const protobufPackage = "mercury.api.v2";

export interface User {
  id: number;
  username: string;
  role: User_Role;
  email: string;
  nickname: string;
  avatarUrl: string;
  password: string;
  rowStatus: RowStatus;
  createTime?: Date | undefined;
  updateTime?: Date | undefined;
}

export enum User_Role {
  ROLE_UNSPECIFIED = 0,
  HOST = 1,
  ADMIN = 2,
  USER = 3,
  UNRECOGNIZED = -1,
}

export function user_RoleFromJSON(object: any): User_Role {
  switch (object) {
    case 0:
    case "ROLE_UNSPECIFIED":
      return User_Role.ROLE_UNSPECIFIED;
    case 1:
    case "HOST":
      return User_Role.HOST;
    case 2:
    case "ADMIN":
      return User_Role.ADMIN;
    case 3:
    case "USER":
      return User_Role.USER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return User_Role.UNRECOGNIZED;
  }
}

export function user_RoleToJSON(object: User_Role): string {
  switch (object) {
    case User_Role.ROLE_UNSPECIFIED:
      return "ROLE_UNSPECIFIED";
    case User_Role.HOST:
      return "HOST";
    case User_Role.ADMIN:
      return "ADMIN";
    case User_Role.USER:
      return "USER";
    case User_Role.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface GetUserRequest {
  username: string;
}

export interface GetUserResponse {
  user?: User | undefined;
}

export interface UpdateUserRequest {
  username: string;
  user?: User | undefined;
  updateMask: string[];
}

export interface UpdateUserResponse {
  user?: User | undefined;
}

export interface ListUserAccessTokensRequest {
  username: string;
}

export interface ListUserAccessTokensResponse {
  accessTokens: UserAccessToken[];
}

export interface CreateUserAccessTokenRequest {
  username: string;
  description: string;
  expiresAt?: Date | undefined;
}

export interface CreateUserAccessTokenResponse {
  accessToken?: UserAccessToken | undefined;
}

export interface DeleteUserAccessTokenRequest {
  username: string;
  accessToken: string;
}

export interface DeleteUserAccessTokenResponse {
}

export interface UserAccessToken {
  accessToken: string;
  description: string;
  issuedAt?: Date | undefined;
  expiresAt?: Date | undefined;
}

function createBaseUser(): User {
  return {
    id: 0,
    username: "",
    role: 0,
    email: "",
    nickname: "",
    avatarUrl: "",
    password: "",
    rowStatus: 0,
    createTime: undefined,
    updateTime: undefined,
  };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.username !== "") {
      writer.uint32(18).string(message.username);
    }
    if (message.role !== 0) {
      writer.uint32(24).int32(message.role);
    }
    if (message.email !== "") {
      writer.uint32(34).string(message.email);
    }
    if (message.nickname !== "") {
      writer.uint32(42).string(message.nickname);
    }
    if (message.avatarUrl !== "") {
      writer.uint32(50).string(message.avatarUrl);
    }
    if (message.password !== "") {
      writer.uint32(58).string(message.password);
    }
    if (message.rowStatus !== 0) {
      writer.uint32(64).int32(message.rowStatus);
    }
    if (message.createTime !== undefined) {
      Timestamp.encode(toTimestamp(message.createTime), writer.uint32(74).fork()).ldelim();
    }
    if (message.updateTime !== undefined) {
      Timestamp.encode(toTimestamp(message.updateTime), writer.uint32(82).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = reader.int32();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.username = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.role = reader.int32() as any;
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.email = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.nickname = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.avatarUrl = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.password = reader.string();
          continue;
        case 8:
          if (tag !== 64) {
            break;
          }

          message.rowStatus = reader.int32() as any;
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.createTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.updateTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): User {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      username: isSet(object.username) ? String(object.username) : "",
      role: isSet(object.role) ? user_RoleFromJSON(object.role) : 0,
      email: isSet(object.email) ? String(object.email) : "",
      nickname: isSet(object.nickname) ? String(object.nickname) : "",
      avatarUrl: isSet(object.avatarUrl) ? String(object.avatarUrl) : "",
      password: isSet(object.password) ? String(object.password) : "",
      rowStatus: isSet(object.rowStatus) ? rowStatusFromJSON(object.rowStatus) : 0,
      createTime: isSet(object.createTime) ? fromJsonTimestamp(object.createTime) : undefined,
      updateTime: isSet(object.updateTime) ? fromJsonTimestamp(object.updateTime) : undefined,
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.username !== undefined && (obj.username = message.username);
    message.role !== undefined && (obj.role = user_RoleToJSON(message.role));
    message.email !== undefined && (obj.email = message.email);
    message.nickname !== undefined && (obj.nickname = message.nickname);
    message.avatarUrl !== undefined && (obj.avatarUrl = message.avatarUrl);
    message.password !== undefined && (obj.password = message.password);
    message.rowStatus !== undefined && (obj.rowStatus = rowStatusToJSON(message.rowStatus));
    message.createTime !== undefined && (obj.createTime = message.createTime.toISOString());
    message.updateTime !== undefined && (obj.updateTime = message.updateTime.toISOString());
    return obj;
  },

  create(base?: DeepPartial<User>): User {
    return User.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<User>): User {
    const message = createBaseUser();
    message.id = object.id ?? 0;
    message.username = object.username ?? "";
    message.role = object.role ?? 0;
    message.email = object.email ?? "";
    message.nickname = object.nickname ?? "";
    message.avatarUrl = object.avatarUrl ?? "";
    message.password = object.password ?? "";
    message.rowStatus = object.rowStatus ?? 0;
    message.createTime = object.createTime ?? undefined;
    message.updateTime = object.updateTime ?? undefined;
    return message;
  },
};

function createBaseGetUserRequest(): GetUserRequest {
  return { username: "" };
}

export const GetUserRequest = {
  encode(message: GetUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.username = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserRequest {
    return { username: isSet(object.username) ? String(object.username) : "" };
  },

  toJSON(message: GetUserRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    return obj;
  },

  create(base?: DeepPartial<GetUserRequest>): GetUserRequest {
    return GetUserRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<GetUserRequest>): GetUserRequest {
    const message = createBaseGetUserRequest();
    message.username = object.username ?? "";
    return message;
  },
};

function createBaseGetUserResponse(): GetUserResponse {
  return { user: undefined };
}

export const GetUserResponse = {
  encode(message: GetUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserResponse {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: GetUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  create(base?: DeepPartial<GetUserResponse>): GetUserResponse {
    return GetUserResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<GetUserResponse>): GetUserResponse {
    const message = createBaseGetUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseUpdateUserRequest(): UpdateUserRequest {
  return { username: "", user: undefined, updateMask: [] };
}

export const UpdateUserRequest = {
  encode(message: UpdateUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.updateMask) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUserRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.username = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.updateMask.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateUserRequest {
    return {
      username: isSet(object.username) ? String(object.username) : "",
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
      updateMask: Array.isArray(object?.updateMask) ? object.updateMask.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: UpdateUserRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    if (message.updateMask) {
      obj.updateMask = message.updateMask.map((e) => e);
    } else {
      obj.updateMask = [];
    }
    return obj;
  },

  create(base?: DeepPartial<UpdateUserRequest>): UpdateUserRequest {
    return UpdateUserRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UpdateUserRequest>): UpdateUserRequest {
    const message = createBaseUpdateUserRequest();
    message.username = object.username ?? "";
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.updateMask = object.updateMask?.map((e) => e) || [];
    return message;
  },
};

function createBaseUpdateUserResponse(): UpdateUserResponse {
  return { user: undefined };
}

export const UpdateUserResponse = {
  encode(message: UpdateUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUserResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateUserResponse {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: UpdateUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  create(base?: DeepPartial<UpdateUserResponse>): UpdateUserResponse {
    return UpdateUserResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UpdateUserResponse>): UpdateUserResponse {
    const message = createBaseUpdateUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseListUserAccessTokensRequest(): ListUserAccessTokensRequest {
  return { username: "" };
}

export const ListUserAccessTokensRequest = {
  encode(message: ListUserAccessTokensRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListUserAccessTokensRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListUserAccessTokensRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.username = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListUserAccessTokensRequest {
    return { username: isSet(object.username) ? String(object.username) : "" };
  },

  toJSON(message: ListUserAccessTokensRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    return obj;
  },

  create(base?: DeepPartial<ListUserAccessTokensRequest>): ListUserAccessTokensRequest {
    return ListUserAccessTokensRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<ListUserAccessTokensRequest>): ListUserAccessTokensRequest {
    const message = createBaseListUserAccessTokensRequest();
    message.username = object.username ?? "";
    return message;
  },
};

function createBaseListUserAccessTokensResponse(): ListUserAccessTokensResponse {
  return { accessTokens: [] };
}

export const ListUserAccessTokensResponse = {
  encode(message: ListUserAccessTokensResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.accessTokens) {
      UserAccessToken.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListUserAccessTokensResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListUserAccessTokensResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.accessTokens.push(UserAccessToken.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListUserAccessTokensResponse {
    return {
      accessTokens: Array.isArray(object?.accessTokens)
        ? object.accessTokens.map((e: any) => UserAccessToken.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ListUserAccessTokensResponse): unknown {
    const obj: any = {};
    if (message.accessTokens) {
      obj.accessTokens = message.accessTokens.map((e) => e ? UserAccessToken.toJSON(e) : undefined);
    } else {
      obj.accessTokens = [];
    }
    return obj;
  },

  create(base?: DeepPartial<ListUserAccessTokensResponse>): ListUserAccessTokensResponse {
    return ListUserAccessTokensResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<ListUserAccessTokensResponse>): ListUserAccessTokensResponse {
    const message = createBaseListUserAccessTokensResponse();
    message.accessTokens = object.accessTokens?.map((e) => UserAccessToken.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreateUserAccessTokenRequest(): CreateUserAccessTokenRequest {
  return { username: "", description: "", expiresAt: undefined };
}

export const CreateUserAccessTokenRequest = {
  encode(message: CreateUserAccessTokenRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.expiresAt !== undefined) {
      Timestamp.encode(toTimestamp(message.expiresAt), writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateUserAccessTokenRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserAccessTokenRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.username = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.description = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.expiresAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateUserAccessTokenRequest {
    return {
      username: isSet(object.username) ? String(object.username) : "",
      description: isSet(object.description) ? String(object.description) : "",
      expiresAt: isSet(object.expiresAt) ? fromJsonTimestamp(object.expiresAt) : undefined,
    };
  },

  toJSON(message: CreateUserAccessTokenRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    message.description !== undefined && (obj.description = message.description);
    message.expiresAt !== undefined && (obj.expiresAt = message.expiresAt.toISOString());
    return obj;
  },

  create(base?: DeepPartial<CreateUserAccessTokenRequest>): CreateUserAccessTokenRequest {
    return CreateUserAccessTokenRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<CreateUserAccessTokenRequest>): CreateUserAccessTokenRequest {
    const message = createBaseCreateUserAccessTokenRequest();
    message.username = object.username ?? "";
    message.description = object.description ?? "";
    message.expiresAt = object.expiresAt ?? undefined;
    return message;
  },
};

function createBaseCreateUserAccessTokenResponse(): CreateUserAccessTokenResponse {
  return { accessToken: undefined };
}

export const CreateUserAccessTokenResponse = {
  encode(message: CreateUserAccessTokenResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accessToken !== undefined) {
      UserAccessToken.encode(message.accessToken, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateUserAccessTokenResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserAccessTokenResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.accessToken = UserAccessToken.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateUserAccessTokenResponse {
    return { accessToken: isSet(object.accessToken) ? UserAccessToken.fromJSON(object.accessToken) : undefined };
  },

  toJSON(message: CreateUserAccessTokenResponse): unknown {
    const obj: any = {};
    message.accessToken !== undefined &&
      (obj.accessToken = message.accessToken ? UserAccessToken.toJSON(message.accessToken) : undefined);
    return obj;
  },

  create(base?: DeepPartial<CreateUserAccessTokenResponse>): CreateUserAccessTokenResponse {
    return CreateUserAccessTokenResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<CreateUserAccessTokenResponse>): CreateUserAccessTokenResponse {
    const message = createBaseCreateUserAccessTokenResponse();
    message.accessToken = (object.accessToken !== undefined && object.accessToken !== null)
      ? UserAccessToken.fromPartial(object.accessToken)
      : undefined;
    return message;
  },
};

function createBaseDeleteUserAccessTokenRequest(): DeleteUserAccessTokenRequest {
  return { username: "", accessToken: "" };
}

export const DeleteUserAccessTokenRequest = {
  encode(message: DeleteUserAccessTokenRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.accessToken !== "") {
      writer.uint32(18).string(message.accessToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteUserAccessTokenRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteUserAccessTokenRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.username = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.accessToken = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteUserAccessTokenRequest {
    return {
      username: isSet(object.username) ? String(object.username) : "",
      accessToken: isSet(object.accessToken) ? String(object.accessToken) : "",
    };
  },

  toJSON(message: DeleteUserAccessTokenRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    message.accessToken !== undefined && (obj.accessToken = message.accessToken);
    return obj;
  },

  create(base?: DeepPartial<DeleteUserAccessTokenRequest>): DeleteUserAccessTokenRequest {
    return DeleteUserAccessTokenRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<DeleteUserAccessTokenRequest>): DeleteUserAccessTokenRequest {
    const message = createBaseDeleteUserAccessTokenRequest();
    message.username = object.username ?? "";
    message.accessToken = object.accessToken ?? "";
    return message;
  },
};

function createBaseDeleteUserAccessTokenResponse(): DeleteUserAccessTokenResponse {
  return {};
}

export const DeleteUserAccessTokenResponse = {
  encode(_: DeleteUserAccessTokenResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteUserAccessTokenResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteUserAccessTokenResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): DeleteUserAccessTokenResponse {
    return {};
  },

  toJSON(_: DeleteUserAccessTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<DeleteUserAccessTokenResponse>): DeleteUserAccessTokenResponse {
    return DeleteUserAccessTokenResponse.fromPartial(base ?? {});
  },

  fromPartial(_: DeepPartial<DeleteUserAccessTokenResponse>): DeleteUserAccessTokenResponse {
    const message = createBaseDeleteUserAccessTokenResponse();
    return message;
  },
};

function createBaseUserAccessToken(): UserAccessToken {
  return { accessToken: "", description: "", issuedAt: undefined, expiresAt: undefined };
}

export const UserAccessToken = {
  encode(message: UserAccessToken, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accessToken !== "") {
      writer.uint32(10).string(message.accessToken);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.issuedAt !== undefined) {
      Timestamp.encode(toTimestamp(message.issuedAt), writer.uint32(26).fork()).ldelim();
    }
    if (message.expiresAt !== undefined) {
      Timestamp.encode(toTimestamp(message.expiresAt), writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserAccessToken {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserAccessToken();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.accessToken = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.description = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.issuedAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.expiresAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UserAccessToken {
    return {
      accessToken: isSet(object.accessToken) ? String(object.accessToken) : "",
      description: isSet(object.description) ? String(object.description) : "",
      issuedAt: isSet(object.issuedAt) ? fromJsonTimestamp(object.issuedAt) : undefined,
      expiresAt: isSet(object.expiresAt) ? fromJsonTimestamp(object.expiresAt) : undefined,
    };
  },

  toJSON(message: UserAccessToken): unknown {
    const obj: any = {};
    message.accessToken !== undefined && (obj.accessToken = message.accessToken);
    message.description !== undefined && (obj.description = message.description);
    message.issuedAt !== undefined && (obj.issuedAt = message.issuedAt.toISOString());
    message.expiresAt !== undefined && (obj.expiresAt = message.expiresAt.toISOString());
    return obj;
  },

  create(base?: DeepPartial<UserAccessToken>): UserAccessToken {
    return UserAccessToken.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UserAccessToken>): UserAccessToken {
    const message = createBaseUserAccessToken();
    message.accessToken = object.accessToken ?? "";
    message.description = object.description ?? "";
    message.issuedAt = object.issuedAt ?? undefined;
    message.expiresAt = object.expiresAt ?? undefined;
    return message;
  },
};

export type UserServiceDefinition = typeof UserServiceDefinition;
export const UserServiceDefinition = {
  name: "UserService",
  fullName: "mercury.api.v2.UserService",
  methods: {
    getUser: {
      name: "GetUser",
      requestType: GetUserRequest,
      requestStream: false,
      responseType: GetUserResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([8, 117, 115, 101, 114, 110, 97, 109, 101])],
          578365826: [
            new Uint8Array([
              25,
              18,
              23,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              117,
              115,
              101,
              114,
              47,
              123,
              117,
              115,
              101,
              114,
              110,
              97,
              109,
              101,
              125,
            ]),
          ],
        },
      },
    },
    updateUser: {
      name: "UpdateUser",
      requestType: UpdateUserRequest,
      requestStream: false,
      responseType: UpdateUserResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([8, 117, 115, 101, 114, 110, 97, 109, 101])],
          578365826: [
            new Uint8Array([
              29,
              58,
              1,
              42,
              34,
              24,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              117,
              115,
              101,
              114,
              115,
              47,
              123,
              117,
              115,
              101,
              114,
              110,
              97,
              109,
              101,
              125,
            ]),
          ],
        },
      },
    },
    listUserAccessTokens: {
      name: "ListUserAccessTokens",
      requestType: ListUserAccessTokensRequest,
      requestStream: false,
      responseType: ListUserAccessTokensResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              40,
              18,
              38,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              117,
              115,
              101,
              114,
              115,
              47,
              123,
              117,
              115,
              101,
              114,
              110,
              97,
              109,
              101,
              125,
              47,
              97,
              99,
              99,
              101,
              115,
              115,
              95,
              116,
              111,
              107,
              101,
              110,
              115,
            ]),
          ],
        },
      },
    },
    createUserAccessToken: {
      name: "CreateUserAccessToken",
      requestType: CreateUserAccessTokenRequest,
      requestStream: false,
      responseType: CreateUserAccessTokenResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([8, 117, 115, 101, 114, 110, 97, 109, 101])],
          578365826: [
            new Uint8Array([
              43,
              58,
              1,
              42,
              34,
              38,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              117,
              115,
              101,
              114,
              115,
              47,
              123,
              117,
              115,
              101,
              114,
              110,
              97,
              109,
              101,
              125,
              47,
              97,
              99,
              99,
              101,
              115,
              115,
              95,
              116,
              111,
              107,
              101,
              110,
              115,
            ]),
          ],
        },
      },
    },
    deleteUserAccessToken: {
      name: "DeleteUserAccessToken",
      requestType: DeleteUserAccessTokenRequest,
      requestStream: false,
      responseType: DeleteUserAccessTokenResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [
            new Uint8Array([
              21,
              117,
              115,
              101,
              114,
              110,
              97,
              109,
              101,
              44,
              97,
              99,
              99,
              101,
              115,
              115,
              95,
              116,
              111,
              107,
              101,
              110,
            ]),
          ],
          578365826: [
            new Uint8Array([
              55,
              42,
              53,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              117,
              115,
              101,
              114,
              115,
              47,
              123,
              117,
              115,
              101,
              114,
              110,
              97,
              109,
              101,
              125,
              47,
              97,
              99,
              99,
              101,
              115,
              115,
              95,
              116,
              111,
              107,
              101,
              110,
              115,
              47,
              123,
              97,
              99,
              99,
              101,
              115,
              115,
              95,
              116,
              111,
              107,
              101,
              110,
              125,
            ]),
          ],
        },
      },
    },
  },
} as const;

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
