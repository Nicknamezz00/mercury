/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "mercury.api.v2";

export interface Resource {
  id: number;
  createdTimestamp?: Date | undefined;
  filename: string;
  externalLink: string;
  type: string;
  size: number;
  relatedMessageId?: number | undefined;
}

export interface ListResourcesRequest {
}

export interface ListResourcesResponse {
  resources: Resource[];
}

export interface CreateResourceRequest {
  filename: string;
  externalLink: string;
  type: string;
  messageId?: number | undefined;
}

export interface CreateResourceResponse {
  resource?: Resource | undefined;
}

export interface UpdateResourceRequest {
  id: number;
  resource?: Resource | undefined;
  updateMask: string[];
}

export interface UpdateResourceResponse {
  resource?: Resource | undefined;
}

export interface DeleteResourceRequest {
  id: number;
}

export interface DeleteResourceResponse {
}

function createBaseResource(): Resource {
  return {
    id: 0,
    createdTimestamp: undefined,
    filename: "",
    externalLink: "",
    type: "",
    size: 0,
    relatedMessageId: undefined,
  };
}

export const Resource = {
  encode(message: Resource, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.createdTimestamp !== undefined) {
      Timestamp.encode(toTimestamp(message.createdTimestamp), writer.uint32(18).fork()).ldelim();
    }
    if (message.filename !== "") {
      writer.uint32(26).string(message.filename);
    }
    if (message.externalLink !== "") {
      writer.uint32(34).string(message.externalLink);
    }
    if (message.type !== "") {
      writer.uint32(42).string(message.type);
    }
    if (message.size !== 0) {
      writer.uint32(48).int64(message.size);
    }
    if (message.relatedMessageId !== undefined) {
      writer.uint32(56).int32(message.relatedMessageId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Resource {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResource();
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

          message.createdTimestamp = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.filename = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.externalLink = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.type = reader.string();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.size = longToNumber(reader.int64() as Long);
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.relatedMessageId = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Resource {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      createdTimestamp: isSet(object.createdTimestamp) ? fromJsonTimestamp(object.createdTimestamp) : undefined,
      filename: isSet(object.filename) ? String(object.filename) : "",
      externalLink: isSet(object.externalLink) ? String(object.externalLink) : "",
      type: isSet(object.type) ? String(object.type) : "",
      size: isSet(object.size) ? Number(object.size) : 0,
      relatedMessageId: isSet(object.relatedMessageId) ? Number(object.relatedMessageId) : undefined,
    };
  },

  toJSON(message: Resource): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.createdTimestamp !== undefined && (obj.createdTimestamp = message.createdTimestamp.toISOString());
    message.filename !== undefined && (obj.filename = message.filename);
    message.externalLink !== undefined && (obj.externalLink = message.externalLink);
    message.type !== undefined && (obj.type = message.type);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.relatedMessageId !== undefined && (obj.relatedMessageId = Math.round(message.relatedMessageId));
    return obj;
  },

  create(base?: DeepPartial<Resource>): Resource {
    return Resource.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<Resource>): Resource {
    const message = createBaseResource();
    message.id = object.id ?? 0;
    message.createdTimestamp = object.createdTimestamp ?? undefined;
    message.filename = object.filename ?? "";
    message.externalLink = object.externalLink ?? "";
    message.type = object.type ?? "";
    message.size = object.size ?? 0;
    message.relatedMessageId = object.relatedMessageId ?? undefined;
    return message;
  },
};

function createBaseListResourcesRequest(): ListResourcesRequest {
  return {};
}

export const ListResourcesRequest = {
  encode(_: ListResourcesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListResourcesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListResourcesRequest();
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

  fromJSON(_: any): ListResourcesRequest {
    return {};
  },

  toJSON(_: ListResourcesRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<ListResourcesRequest>): ListResourcesRequest {
    return ListResourcesRequest.fromPartial(base ?? {});
  },

  fromPartial(_: DeepPartial<ListResourcesRequest>): ListResourcesRequest {
    const message = createBaseListResourcesRequest();
    return message;
  },
};

function createBaseListResourcesResponse(): ListResourcesResponse {
  return { resources: [] };
}

export const ListResourcesResponse = {
  encode(message: ListResourcesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.resources) {
      Resource.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListResourcesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListResourcesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.resources.push(Resource.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListResourcesResponse {
    return {
      resources: Array.isArray(object?.resources) ? object.resources.map((e: any) => Resource.fromJSON(e)) : [],
    };
  },

  toJSON(message: ListResourcesResponse): unknown {
    const obj: any = {};
    if (message.resources) {
      obj.resources = message.resources.map((e) => e ? Resource.toJSON(e) : undefined);
    } else {
      obj.resources = [];
    }
    return obj;
  },

  create(base?: DeepPartial<ListResourcesResponse>): ListResourcesResponse {
    return ListResourcesResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<ListResourcesResponse>): ListResourcesResponse {
    const message = createBaseListResourcesResponse();
    message.resources = object.resources?.map((e) => Resource.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreateResourceRequest(): CreateResourceRequest {
  return { filename: "", externalLink: "", type: "", messageId: undefined };
}

export const CreateResourceRequest = {
  encode(message: CreateResourceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.filename !== "") {
      writer.uint32(10).string(message.filename);
    }
    if (message.externalLink !== "") {
      writer.uint32(18).string(message.externalLink);
    }
    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }
    if (message.messageId !== undefined) {
      writer.uint32(32).int32(message.messageId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateResourceRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateResourceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.filename = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.externalLink = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.type = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.messageId = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateResourceRequest {
    return {
      filename: isSet(object.filename) ? String(object.filename) : "",
      externalLink: isSet(object.externalLink) ? String(object.externalLink) : "",
      type: isSet(object.type) ? String(object.type) : "",
      messageId: isSet(object.messageId) ? Number(object.messageId) : undefined,
    };
  },

  toJSON(message: CreateResourceRequest): unknown {
    const obj: any = {};
    message.filename !== undefined && (obj.filename = message.filename);
    message.externalLink !== undefined && (obj.externalLink = message.externalLink);
    message.type !== undefined && (obj.type = message.type);
    message.messageId !== undefined && (obj.messageId = Math.round(message.messageId));
    return obj;
  },

  create(base?: DeepPartial<CreateResourceRequest>): CreateResourceRequest {
    return CreateResourceRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<CreateResourceRequest>): CreateResourceRequest {
    const message = createBaseCreateResourceRequest();
    message.filename = object.filename ?? "";
    message.externalLink = object.externalLink ?? "";
    message.type = object.type ?? "";
    message.messageId = object.messageId ?? undefined;
    return message;
  },
};

function createBaseCreateResourceResponse(): CreateResourceResponse {
  return { resource: undefined };
}

export const CreateResourceResponse = {
  encode(message: CreateResourceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.resource !== undefined) {
      Resource.encode(message.resource, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateResourceResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateResourceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.resource = Resource.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateResourceResponse {
    return { resource: isSet(object.resource) ? Resource.fromJSON(object.resource) : undefined };
  },

  toJSON(message: CreateResourceResponse): unknown {
    const obj: any = {};
    message.resource !== undefined && (obj.resource = message.resource ? Resource.toJSON(message.resource) : undefined);
    return obj;
  },

  create(base?: DeepPartial<CreateResourceResponse>): CreateResourceResponse {
    return CreateResourceResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<CreateResourceResponse>): CreateResourceResponse {
    const message = createBaseCreateResourceResponse();
    message.resource = (object.resource !== undefined && object.resource !== null)
      ? Resource.fromPartial(object.resource)
      : undefined;
    return message;
  },
};

function createBaseUpdateResourceRequest(): UpdateResourceRequest {
  return { id: 0, resource: undefined, updateMask: [] };
}

export const UpdateResourceRequest = {
  encode(message: UpdateResourceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.resource !== undefined) {
      Resource.encode(message.resource, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.updateMask) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateResourceRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateResourceRequest();
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

          message.resource = Resource.decode(reader, reader.uint32());
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

  fromJSON(object: any): UpdateResourceRequest {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      resource: isSet(object.resource) ? Resource.fromJSON(object.resource) : undefined,
      updateMask: Array.isArray(object?.updateMask) ? object.updateMask.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: UpdateResourceRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.resource !== undefined && (obj.resource = message.resource ? Resource.toJSON(message.resource) : undefined);
    if (message.updateMask) {
      obj.updateMask = message.updateMask.map((e) => e);
    } else {
      obj.updateMask = [];
    }
    return obj;
  },

  create(base?: DeepPartial<UpdateResourceRequest>): UpdateResourceRequest {
    return UpdateResourceRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UpdateResourceRequest>): UpdateResourceRequest {
    const message = createBaseUpdateResourceRequest();
    message.id = object.id ?? 0;
    message.resource = (object.resource !== undefined && object.resource !== null)
      ? Resource.fromPartial(object.resource)
      : undefined;
    message.updateMask = object.updateMask?.map((e) => e) || [];
    return message;
  },
};

function createBaseUpdateResourceResponse(): UpdateResourceResponse {
  return { resource: undefined };
}

export const UpdateResourceResponse = {
  encode(message: UpdateResourceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.resource !== undefined) {
      Resource.encode(message.resource, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateResourceResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateResourceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.resource = Resource.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateResourceResponse {
    return { resource: isSet(object.resource) ? Resource.fromJSON(object.resource) : undefined };
  },

  toJSON(message: UpdateResourceResponse): unknown {
    const obj: any = {};
    message.resource !== undefined && (obj.resource = message.resource ? Resource.toJSON(message.resource) : undefined);
    return obj;
  },

  create(base?: DeepPartial<UpdateResourceResponse>): UpdateResourceResponse {
    return UpdateResourceResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UpdateResourceResponse>): UpdateResourceResponse {
    const message = createBaseUpdateResourceResponse();
    message.resource = (object.resource !== undefined && object.resource !== null)
      ? Resource.fromPartial(object.resource)
      : undefined;
    return message;
  },
};

function createBaseDeleteResourceRequest(): DeleteResourceRequest {
  return { id: 0 };
}

export const DeleteResourceRequest = {
  encode(message: DeleteResourceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteResourceRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteResourceRequest();
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

  fromJSON(object: any): DeleteResourceRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: DeleteResourceRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  create(base?: DeepPartial<DeleteResourceRequest>): DeleteResourceRequest {
    return DeleteResourceRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<DeleteResourceRequest>): DeleteResourceRequest {
    const message = createBaseDeleteResourceRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseDeleteResourceResponse(): DeleteResourceResponse {
  return {};
}

export const DeleteResourceResponse = {
  encode(_: DeleteResourceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteResourceResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteResourceResponse();
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

  fromJSON(_: any): DeleteResourceResponse {
    return {};
  },

  toJSON(_: DeleteResourceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<DeleteResourceResponse>): DeleteResourceResponse {
    return DeleteResourceResponse.fromPartial(base ?? {});
  },

  fromPartial(_: DeepPartial<DeleteResourceResponse>): DeleteResourceResponse {
    const message = createBaseDeleteResourceResponse();
    return message;
  },
};

export type ResourceServiceDefinition = typeof ResourceServiceDefinition;
export const ResourceServiceDefinition = {
  name: "ResourceService",
  fullName: "mercury.api.v2.ResourceService",
  methods: {
    listResources: {
      name: "ListResources",
      requestType: ListResourcesRequest,
      requestStream: false,
      responseType: ListResourcesResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([19, 18, 17, 47, 97, 112, 105, 47, 118, 50, 47, 114, 101, 115, 111, 117, 114, 99, 101, 115]),
          ],
        },
      },
    },
    createResource: {
      name: "CreateResource",
      requestType: CreateResourceRequest,
      requestStream: false,
      responseType: CreateResourceResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([19, 34, 17, 47, 97, 112, 105, 47, 118, 50, 47, 114, 101, 115, 111, 117, 114, 99, 101, 115]),
          ],
        },
      },
    },
    updateResource: {
      name: "UpdateResource",
      requestType: UpdateResourceRequest,
      requestStream: false,
      responseType: UpdateResourceResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([2, 105, 100])],
          578365826: [
            new Uint8Array([
              24,
              26,
              22,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              114,
              101,
              115,
              111,
              117,
              114,
              99,
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
    deleteResource: {
      name: "DeleteResource",
      requestType: DeleteResourceRequest,
      requestStream: false,
      responseType: DeleteResourceResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([2, 105, 100])],
          578365826: [
            new Uint8Array([
              24,
              18,
              22,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              114,
              101,
              115,
              111,
              117,
              114,
              99,
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
