/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { RowStatus, rowStatusFromJSON, rowStatusToJSON } from "./common";

export const protobufPackage = "mercury.api.v2";

export enum Visibility {
  VISIBILITY_UNSPECIFIED = 0,
  PRIVATE = 1,
  PROTECTED = 2,
  PUBLIC = 3,
  UNRECOGNIZED = -1,
}

export function visibilityFromJSON(object: any): Visibility {
  switch (object) {
    case 0:
    case "VISIBILITY_UNSPECIFIED":
      return Visibility.VISIBILITY_UNSPECIFIED;
    case 1:
    case "PRIVATE":
      return Visibility.PRIVATE;
    case 2:
    case "PROTECTED":
      return Visibility.PROTECTED;
    case 3:
    case "PUBLIC":
      return Visibility.PUBLIC;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Visibility.UNRECOGNIZED;
  }
}

export function visibilityToJSON(object: Visibility): string {
  switch (object) {
    case Visibility.VISIBILITY_UNSPECIFIED:
      return "VISIBILITY_UNSPECIFIED";
    case Visibility.PRIVATE:
      return "PRIVATE";
    case Visibility.PROTECTED:
      return "PROTECTED";
    case Visibility.PUBLIC:
      return "PUBLIC";
    case Visibility.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Message {
  id: number;
  rowStatus: RowStatus;
  creatorId: number;
  creatorTimestamp: number;
  updatedTimestamp: number;
  content: string;
  visibility: Visibility;
  pinned: boolean;
}

export interface CreateMessageRequest {
  content: string;
  visibility: Visibility;
}

export interface CreateMessageResponse {
  message?: Message | undefined;
}

export interface ListMessagesRequest {
  page: number;
  pageSize: number;
  filter: string;
}

export interface ListMessagesResponse {
  messages: Message[];
}

export interface GetMessageRequest {
  id: number;
}

export interface GetMessageResponse {
  message?: Message | undefined;
}

export interface CreateMessageCommentRequest {
  /** id to create comment for. */
  id: number;
  create?: CreateMessageRequest | undefined;
}

export interface CreateMessageCommentResponse {
  message?: Message | undefined;
}

export interface ListMessageCommentsRequest {
  id: number;
}

export interface ListMessageCommentsResponse {
  messages: Message[];
}

function createBaseMessage(): Message {
  return {
    id: 0,
    rowStatus: 0,
    creatorId: 0,
    creatorTimestamp: 0,
    updatedTimestamp: 0,
    content: "",
    visibility: 0,
    pinned: false,
  };
}

export const Message = {
  encode(message: Message, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.rowStatus !== 0) {
      writer.uint32(16).int32(message.rowStatus);
    }
    if (message.creatorId !== 0) {
      writer.uint32(24).int32(message.creatorId);
    }
    if (message.creatorTimestamp !== 0) {
      writer.uint32(32).int64(message.creatorTimestamp);
    }
    if (message.updatedTimestamp !== 0) {
      writer.uint32(40).int64(message.updatedTimestamp);
    }
    if (message.content !== "") {
      writer.uint32(50).string(message.content);
    }
    if (message.visibility !== 0) {
      writer.uint32(56).int32(message.visibility);
    }
    if (message.pinned === true) {
      writer.uint32(64).bool(message.pinned);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Message {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMessage();
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
          if (tag !== 16) {
            break;
          }

          message.rowStatus = reader.int32() as any;
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.creatorId = reader.int32();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.creatorTimestamp = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.updatedTimestamp = longToNumber(reader.int64() as Long);
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.content = reader.string();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.visibility = reader.int32() as any;
          continue;
        case 8:
          if (tag !== 64) {
            break;
          }

          message.pinned = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Message {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      rowStatus: isSet(object.rowStatus) ? rowStatusFromJSON(object.rowStatus) : 0,
      creatorId: isSet(object.creatorId) ? Number(object.creatorId) : 0,
      creatorTimestamp: isSet(object.creatorTimestamp) ? Number(object.creatorTimestamp) : 0,
      updatedTimestamp: isSet(object.updatedTimestamp) ? Number(object.updatedTimestamp) : 0,
      content: isSet(object.content) ? String(object.content) : "",
      visibility: isSet(object.visibility) ? visibilityFromJSON(object.visibility) : 0,
      pinned: isSet(object.pinned) ? Boolean(object.pinned) : false,
    };
  },

  toJSON(message: Message): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.rowStatus !== undefined && (obj.rowStatus = rowStatusToJSON(message.rowStatus));
    message.creatorId !== undefined && (obj.creatorId = Math.round(message.creatorId));
    message.creatorTimestamp !== undefined && (obj.creatorTimestamp = Math.round(message.creatorTimestamp));
    message.updatedTimestamp !== undefined && (obj.updatedTimestamp = Math.round(message.updatedTimestamp));
    message.content !== undefined && (obj.content = message.content);
    message.visibility !== undefined && (obj.visibility = visibilityToJSON(message.visibility));
    message.pinned !== undefined && (obj.pinned = message.pinned);
    return obj;
  },

  create(base?: DeepPartial<Message>): Message {
    return Message.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<Message>): Message {
    const message = createBaseMessage();
    message.id = object.id ?? 0;
    message.rowStatus = object.rowStatus ?? 0;
    message.creatorId = object.creatorId ?? 0;
    message.creatorTimestamp = object.creatorTimestamp ?? 0;
    message.updatedTimestamp = object.updatedTimestamp ?? 0;
    message.content = object.content ?? "";
    message.visibility = object.visibility ?? 0;
    message.pinned = object.pinned ?? false;
    return message;
  },
};

function createBaseCreateMessageRequest(): CreateMessageRequest {
  return { content: "", visibility: 0 };
}

export const CreateMessageRequest = {
  encode(message: CreateMessageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.content !== "") {
      writer.uint32(10).string(message.content);
    }
    if (message.visibility !== 0) {
      writer.uint32(16).int32(message.visibility);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateMessageRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateMessageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.content = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.visibility = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateMessageRequest {
    return {
      content: isSet(object.content) ? String(object.content) : "",
      visibility: isSet(object.visibility) ? visibilityFromJSON(object.visibility) : 0,
    };
  },

  toJSON(message: CreateMessageRequest): unknown {
    const obj: any = {};
    message.content !== undefined && (obj.content = message.content);
    message.visibility !== undefined && (obj.visibility = visibilityToJSON(message.visibility));
    return obj;
  },

  create(base?: DeepPartial<CreateMessageRequest>): CreateMessageRequest {
    return CreateMessageRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<CreateMessageRequest>): CreateMessageRequest {
    const message = createBaseCreateMessageRequest();
    message.content = object.content ?? "";
    message.visibility = object.visibility ?? 0;
    return message;
  },
};

function createBaseCreateMessageResponse(): CreateMessageResponse {
  return { message: undefined };
}

export const CreateMessageResponse = {
  encode(message: CreateMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.message !== undefined) {
      Message.encode(message.message, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateMessageResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.message = Message.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateMessageResponse {
    return { message: isSet(object.message) ? Message.fromJSON(object.message) : undefined };
  },

  toJSON(message: CreateMessageResponse): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message ? Message.toJSON(message.message) : undefined);
    return obj;
  },

  create(base?: DeepPartial<CreateMessageResponse>): CreateMessageResponse {
    return CreateMessageResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<CreateMessageResponse>): CreateMessageResponse {
    const message = createBaseCreateMessageResponse();
    message.message = (object.message !== undefined && object.message !== null)
      ? Message.fromPartial(object.message)
      : undefined;
    return message;
  },
};

function createBaseListMessagesRequest(): ListMessagesRequest {
  return { page: 0, pageSize: 0, filter: "" };
}

export const ListMessagesRequest = {
  encode(message: ListMessagesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.page !== 0) {
      writer.uint32(8).int32(message.page);
    }
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    if (message.filter !== "") {
      writer.uint32(26).string(message.filter);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListMessagesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListMessagesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.page = reader.int32();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.pageSize = reader.int32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.filter = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListMessagesRequest {
    return {
      page: isSet(object.page) ? Number(object.page) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      filter: isSet(object.filter) ? String(object.filter) : "",
    };
  },

  toJSON(message: ListMessagesRequest): unknown {
    const obj: any = {};
    message.page !== undefined && (obj.page = Math.round(message.page));
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.filter !== undefined && (obj.filter = message.filter);
    return obj;
  },

  create(base?: DeepPartial<ListMessagesRequest>): ListMessagesRequest {
    return ListMessagesRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<ListMessagesRequest>): ListMessagesRequest {
    const message = createBaseListMessagesRequest();
    message.page = object.page ?? 0;
    message.pageSize = object.pageSize ?? 0;
    message.filter = object.filter ?? "";
    return message;
  },
};

function createBaseListMessagesResponse(): ListMessagesResponse {
  return { messages: [] };
}

export const ListMessagesResponse = {
  encode(message: ListMessagesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.messages) {
      Message.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListMessagesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListMessagesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.messages.push(Message.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListMessagesResponse {
    return { messages: Array.isArray(object?.messages) ? object.messages.map((e: any) => Message.fromJSON(e)) : [] };
  },

  toJSON(message: ListMessagesResponse): unknown {
    const obj: any = {};
    if (message.messages) {
      obj.messages = message.messages.map((e) => e ? Message.toJSON(e) : undefined);
    } else {
      obj.messages = [];
    }
    return obj;
  },

  create(base?: DeepPartial<ListMessagesResponse>): ListMessagesResponse {
    return ListMessagesResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<ListMessagesResponse>): ListMessagesResponse {
    const message = createBaseListMessagesResponse();
    message.messages = object.messages?.map((e) => Message.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetMessageRequest(): GetMessageRequest {
  return { id: 0 };
}

export const GetMessageRequest = {
  encode(message: GetMessageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMessageRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMessageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMessageRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: GetMessageRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  create(base?: DeepPartial<GetMessageRequest>): GetMessageRequest {
    return GetMessageRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<GetMessageRequest>): GetMessageRequest {
    const message = createBaseGetMessageRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseGetMessageResponse(): GetMessageResponse {
  return { message: undefined };
}

export const GetMessageResponse = {
  encode(message: GetMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.message !== undefined) {
      Message.encode(message.message, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMessageResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.message = Message.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMessageResponse {
    return { message: isSet(object.message) ? Message.fromJSON(object.message) : undefined };
  },

  toJSON(message: GetMessageResponse): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message ? Message.toJSON(message.message) : undefined);
    return obj;
  },

  create(base?: DeepPartial<GetMessageResponse>): GetMessageResponse {
    return GetMessageResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<GetMessageResponse>): GetMessageResponse {
    const message = createBaseGetMessageResponse();
    message.message = (object.message !== undefined && object.message !== null)
      ? Message.fromPartial(object.message)
      : undefined;
    return message;
  },
};

function createBaseCreateMessageCommentRequest(): CreateMessageCommentRequest {
  return { id: 0, create: undefined };
}

export const CreateMessageCommentRequest = {
  encode(message: CreateMessageCommentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.create !== undefined) {
      CreateMessageRequest.encode(message.create, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateMessageCommentRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateMessageCommentRequest();
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

          message.create = CreateMessageRequest.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateMessageCommentRequest {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      create: isSet(object.create) ? CreateMessageRequest.fromJSON(object.create) : undefined,
    };
  },

  toJSON(message: CreateMessageCommentRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.create !== undefined &&
      (obj.create = message.create ? CreateMessageRequest.toJSON(message.create) : undefined);
    return obj;
  },

  create(base?: DeepPartial<CreateMessageCommentRequest>): CreateMessageCommentRequest {
    return CreateMessageCommentRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<CreateMessageCommentRequest>): CreateMessageCommentRequest {
    const message = createBaseCreateMessageCommentRequest();
    message.id = object.id ?? 0;
    message.create = (object.create !== undefined && object.create !== null)
      ? CreateMessageRequest.fromPartial(object.create)
      : undefined;
    return message;
  },
};

function createBaseCreateMessageCommentResponse(): CreateMessageCommentResponse {
  return { message: undefined };
}

export const CreateMessageCommentResponse = {
  encode(message: CreateMessageCommentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.message !== undefined) {
      Message.encode(message.message, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateMessageCommentResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateMessageCommentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.message = Message.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateMessageCommentResponse {
    return { message: isSet(object.message) ? Message.fromJSON(object.message) : undefined };
  },

  toJSON(message: CreateMessageCommentResponse): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message ? Message.toJSON(message.message) : undefined);
    return obj;
  },

  create(base?: DeepPartial<CreateMessageCommentResponse>): CreateMessageCommentResponse {
    return CreateMessageCommentResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<CreateMessageCommentResponse>): CreateMessageCommentResponse {
    const message = createBaseCreateMessageCommentResponse();
    message.message = (object.message !== undefined && object.message !== null)
      ? Message.fromPartial(object.message)
      : undefined;
    return message;
  },
};

function createBaseListMessageCommentsRequest(): ListMessageCommentsRequest {
  return { id: 0 };
}

export const ListMessageCommentsRequest = {
  encode(message: ListMessageCommentsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListMessageCommentsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListMessageCommentsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListMessageCommentsRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: ListMessageCommentsRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  create(base?: DeepPartial<ListMessageCommentsRequest>): ListMessageCommentsRequest {
    return ListMessageCommentsRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<ListMessageCommentsRequest>): ListMessageCommentsRequest {
    const message = createBaseListMessageCommentsRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseListMessageCommentsResponse(): ListMessageCommentsResponse {
  return { messages: [] };
}

export const ListMessageCommentsResponse = {
  encode(message: ListMessageCommentsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.messages) {
      Message.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListMessageCommentsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListMessageCommentsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.messages.push(Message.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListMessageCommentsResponse {
    return { messages: Array.isArray(object?.messages) ? object.messages.map((e: any) => Message.fromJSON(e)) : [] };
  },

  toJSON(message: ListMessageCommentsResponse): unknown {
    const obj: any = {};
    if (message.messages) {
      obj.messages = message.messages.map((e) => e ? Message.toJSON(e) : undefined);
    } else {
      obj.messages = [];
    }
    return obj;
  },

  create(base?: DeepPartial<ListMessageCommentsResponse>): ListMessageCommentsResponse {
    return ListMessageCommentsResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<ListMessageCommentsResponse>): ListMessageCommentsResponse {
    const message = createBaseListMessageCommentsResponse();
    message.messages = object.messages?.map((e) => Message.fromPartial(e)) || [];
    return message;
  },
};

export type MessageServiceDefinition = typeof MessageServiceDefinition;
export const MessageServiceDefinition = {
  name: "MessageService",
  fullName: "mercury.api.v2.MessageService",
  methods: {
    createMessage: {
      name: "CreateMessage",
      requestType: CreateMessageRequest,
      requestStream: false,
      responseType: CreateMessageResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([18, 34, 16, 47, 97, 112, 105, 47, 118, 50, 47, 109, 101, 115, 115, 97, 103, 101, 115]),
          ],
        },
      },
    },
    listMessages: {
      name: "ListMessages",
      requestType: ListMessagesRequest,
      requestStream: false,
      responseType: ListMessagesResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([18, 18, 16, 47, 97, 112, 105, 47, 118, 50, 47, 109, 101, 115, 115, 97, 103, 101, 115]),
          ],
        },
      },
    },
    getMessage: {
      name: "GetMessage",
      requestType: GetMessageRequest,
      requestStream: false,
      responseType: GetMessageResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([2, 105, 100])],
          578365826: [
            new Uint8Array([
              23,
              18,
              21,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              109,
              101,
              115,
              115,
              97,
              103,
              101,
              115,
              47,
              123,
              105,
              100,
              125,
            ]),
          ],
        },
      },
    },
    createMessageComment: {
      name: "CreateMessageComment",
      requestType: CreateMessageCommentRequest,
      requestStream: false,
      responseType: CreateMessageCommentResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([2, 105, 100])],
          578365826: [
            new Uint8Array([
              32,
              34,
              30,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              109,
              101,
              115,
              115,
              97,
              103,
              101,
              115,
              47,
              123,
              105,
              100,
              125,
              47,
              99,
              111,
              109,
              109,
              101,
              110,
              116,
              115,
            ]),
          ],
        },
      },
    },
    listMessageComments: {
      name: "ListMessageComments",
      requestType: ListMessageCommentsRequest,
      requestStream: false,
      responseType: ListMessageCommentsResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([2, 105, 100])],
          578365826: [
            new Uint8Array([
              32,
              18,
              30,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              109,
              101,
              115,
              115,
              97,
              103,
              101,
              115,
              47,
              123,
              105,
              100,
              125,
              47,
              99,
              111,
              109,
              109,
              101,
              110,
              116,
              115,
            ]),
          ],
        },
      },
    },
  },
} as const;

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
