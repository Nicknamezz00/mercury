/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "mercury.api.v2";

export interface Tag {
  name: string;
  creatorId: number;
}

export interface ListTagsRequest {
  creatorId: number;
}

export interface ListTagsResponse {
  tags: Tag[];
}

export interface UpsertTagRequest {
  name: string;
}

export interface UpsertTagResponse {
  tag?: Tag | undefined;
}

export interface DeleteTagRequest {
  tag?: Tag | undefined;
}

export interface DeleteTagResponse {
}

function createBaseTag(): Tag {
  return { name: "", creatorId: 0 };
}

export const Tag = {
  encode(message: Tag, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.creatorId !== 0) {
      writer.uint32(16).int32(message.creatorId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Tag {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTag();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.creatorId = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Tag {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      creatorId: isSet(object.creatorId) ? Number(object.creatorId) : 0,
    };
  },

  toJSON(message: Tag): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.creatorId !== undefined && (obj.creatorId = Math.round(message.creatorId));
    return obj;
  },

  create(base?: DeepPartial<Tag>): Tag {
    return Tag.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<Tag>): Tag {
    const message = createBaseTag();
    message.name = object.name ?? "";
    message.creatorId = object.creatorId ?? 0;
    return message;
  },
};

function createBaseListTagsRequest(): ListTagsRequest {
  return { creatorId: 0 };
}

export const ListTagsRequest = {
  encode(message: ListTagsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creatorId !== 0) {
      writer.uint32(8).int32(message.creatorId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListTagsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListTagsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.creatorId = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListTagsRequest {
    return { creatorId: isSet(object.creatorId) ? Number(object.creatorId) : 0 };
  },

  toJSON(message: ListTagsRequest): unknown {
    const obj: any = {};
    message.creatorId !== undefined && (obj.creatorId = Math.round(message.creatorId));
    return obj;
  },

  create(base?: DeepPartial<ListTagsRequest>): ListTagsRequest {
    return ListTagsRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<ListTagsRequest>): ListTagsRequest {
    const message = createBaseListTagsRequest();
    message.creatorId = object.creatorId ?? 0;
    return message;
  },
};

function createBaseListTagsResponse(): ListTagsResponse {
  return { tags: [] };
}

export const ListTagsResponse = {
  encode(message: ListTagsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.tags) {
      Tag.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListTagsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListTagsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tags.push(Tag.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListTagsResponse {
    return { tags: Array.isArray(object?.tags) ? object.tags.map((e: any) => Tag.fromJSON(e)) : [] };
  },

  toJSON(message: ListTagsResponse): unknown {
    const obj: any = {};
    if (message.tags) {
      obj.tags = message.tags.map((e) => e ? Tag.toJSON(e) : undefined);
    } else {
      obj.tags = [];
    }
    return obj;
  },

  create(base?: DeepPartial<ListTagsResponse>): ListTagsResponse {
    return ListTagsResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<ListTagsResponse>): ListTagsResponse {
    const message = createBaseListTagsResponse();
    message.tags = object.tags?.map((e) => Tag.fromPartial(e)) || [];
    return message;
  },
};

function createBaseUpsertTagRequest(): UpsertTagRequest {
  return { name: "" };
}

export const UpsertTagRequest = {
  encode(message: UpsertTagRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpsertTagRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpsertTagRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpsertTagRequest {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: UpsertTagRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create(base?: DeepPartial<UpsertTagRequest>): UpsertTagRequest {
    return UpsertTagRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UpsertTagRequest>): UpsertTagRequest {
    const message = createBaseUpsertTagRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseUpsertTagResponse(): UpsertTagResponse {
  return { tag: undefined };
}

export const UpsertTagResponse = {
  encode(message: UpsertTagResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tag !== undefined) {
      Tag.encode(message.tag, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpsertTagResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpsertTagResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tag = Tag.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpsertTagResponse {
    return { tag: isSet(object.tag) ? Tag.fromJSON(object.tag) : undefined };
  },

  toJSON(message: UpsertTagResponse): unknown {
    const obj: any = {};
    message.tag !== undefined && (obj.tag = message.tag ? Tag.toJSON(message.tag) : undefined);
    return obj;
  },

  create(base?: DeepPartial<UpsertTagResponse>): UpsertTagResponse {
    return UpsertTagResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UpsertTagResponse>): UpsertTagResponse {
    const message = createBaseUpsertTagResponse();
    message.tag = (object.tag !== undefined && object.tag !== null) ? Tag.fromPartial(object.tag) : undefined;
    return message;
  },
};

function createBaseDeleteTagRequest(): DeleteTagRequest {
  return { tag: undefined };
}

export const DeleteTagRequest = {
  encode(message: DeleteTagRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tag !== undefined) {
      Tag.encode(message.tag, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteTagRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteTagRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tag = Tag.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteTagRequest {
    return { tag: isSet(object.tag) ? Tag.fromJSON(object.tag) : undefined };
  },

  toJSON(message: DeleteTagRequest): unknown {
    const obj: any = {};
    message.tag !== undefined && (obj.tag = message.tag ? Tag.toJSON(message.tag) : undefined);
    return obj;
  },

  create(base?: DeepPartial<DeleteTagRequest>): DeleteTagRequest {
    return DeleteTagRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<DeleteTagRequest>): DeleteTagRequest {
    const message = createBaseDeleteTagRequest();
    message.tag = (object.tag !== undefined && object.tag !== null) ? Tag.fromPartial(object.tag) : undefined;
    return message;
  },
};

function createBaseDeleteTagResponse(): DeleteTagResponse {
  return {};
}

export const DeleteTagResponse = {
  encode(_: DeleteTagResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteTagResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteTagResponse();
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

  fromJSON(_: any): DeleteTagResponse {
    return {};
  },

  toJSON(_: DeleteTagResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<DeleteTagResponse>): DeleteTagResponse {
    return DeleteTagResponse.fromPartial(base ?? {});
  },

  fromPartial(_: DeepPartial<DeleteTagResponse>): DeleteTagResponse {
    const message = createBaseDeleteTagResponse();
    return message;
  },
};

export type TagServiceDefinition = typeof TagServiceDefinition;
export const TagServiceDefinition = {
  name: "TagService",
  fullName: "mercury.api.v2.TagService",
  methods: {
    listTags: {
      name: "ListTags",
      requestType: ListTagsRequest,
      requestStream: false,
      responseType: ListTagsResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [new Uint8Array([14, 18, 12, 47, 97, 112, 105, 47, 118, 50, 47, 116, 97, 103, 115])],
        },
      },
    },
    upsertTag: {
      name: "UpsertTag",
      requestType: UpsertTagRequest,
      requestStream: false,
      responseType: UpsertTagResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [new Uint8Array([14, 34, 12, 47, 97, 112, 105, 47, 118, 50, 47, 116, 97, 103, 115])],
        },
      },
    },
    deleteTag: {
      name: "DeleteTag",
      requestType: DeleteTagRequest,
      requestStream: false,
      responseType: DeleteTagResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [new Uint8Array([14, 42, 12, 47, 97, 112, 105, 47, 118, 50, 47, 116, 97, 103, 115])],
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
