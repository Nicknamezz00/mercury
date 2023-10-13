/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Api } from "../protobuf/api";
import { Enum, Type } from "../protobuf/type";
import { UInt32Value } from "../protobuf/wrappers";
import { Authentication } from "./auth";
import { Backend } from "./backend";
import { Billing } from "./billing";
import { Publishing } from "./client";
import { Context } from "./context";
import { Control } from "./control";
import { Documentation } from "./documentation";
import { Endpoint } from "./endpoint";
import { Http } from "./http";
import { LogDescriptor } from "./log";
import { Logging } from "./logging";
import { MetricDescriptor } from "./metric";
import { MonitoredResourceDescriptor } from "./monitored_resource";
import { Monitoring } from "./monitoring";
import { Quota } from "./quota";
import { SourceInfo } from "./source_info";
import { SystemParameters } from "./system_parameter";
import { Usage } from "./usage";

export const protobufPackage = "google.api";

/**
 * `Service` is the root object of Google API service configuration (service
 * config). It describes the basic information about a logical service,
 * such as the service name and the user-facing title, and delegates other
 * aspects to sub-sections. Each sub-section is either a proto message or a
 * repeated proto message that configures a specific aspect, such as auth.
 * For more information, see each proto message definition.
 *
 * Example:
 *
 *     type: google.api.Service
 *     name: calendar.googleapis.com
 *     title: Google Calendar API
 *     apis:
 *     - name: google.calendar.v3.Calendar
 *
 *     visibility:
 *       rules:
 *       - selector: "google.calendar.v3.*"
 *         restriction: PREVIEW
 *     backend:
 *       rules:
 *       - selector: "google.calendar.v3.*"
 *         address: calendar.example.com
 *
 *     authentication:
 *       providers:
 *       - id: google_calendar_auth
 *         jwks_uri: https://www.googleapis.com/oauth2/v1/certs
 *         issuer: https://securetoken.google.com
 *       rules:
 *       - selector: "*"
 *         requirements:
 *           provider_id: google_calendar_auth
 */
export interface Service {
  /**
   * The service name, which is a DNS-like logical identifier for the
   * service, such as `calendar.googleapis.com`. The service name
   * typically goes through DNS verification to make sure the owner
   * of the service also owns the DNS name.
   */
  name: string;
  /**
   * The product title for this service, it is the name displayed in Google
   * Cloud Console.
   */
  title: string;
  /** The Google project that owns this service. */
  producerProjectId: string;
  /**
   * A unique ID for a specific instance of this message, typically assigned
   * by the client for tracking purpose. Must be no longer than 63 characters
   * and only lower case letters, digits, '.', '_' and '-' are allowed. If
   * empty, the server may choose to generate one instead.
   */
  id: string;
  /**
   * A list of API interfaces exported by this service. Only the `name` field
   * of the [google.protobuf.Api][google.protobuf.Api] needs to be provided by
   * the configuration author, as the remaining fields will be derived from the
   * IDL during the normalization process. It is an error to specify an API
   * interface here which cannot be resolved against the associated IDL files.
   */
  apis: Api[];
  /**
   * A list of all proto message types included in this API service.
   * Types referenced directly or indirectly by the `apis` are automatically
   * included.  Messages which are not referenced but shall be included, such as
   * types used by the `google.protobuf.Any` type, should be listed here by
   * name by the configuration author. Example:
   *
   *     types:
   *     - name: google.protobuf.Int32
   */
  types: Type[];
  /**
   * A list of all enum types included in this API service.  Enums referenced
   * directly or indirectly by the `apis` are automatically included.  Enums
   * which are not referenced but shall be included should be listed here by
   * name by the configuration author. Example:
   *
   *     enums:
   *     - name: google.someapi.v1.SomeEnum
   */
  enums: Enum[];
  /** Additional API documentation. */
  documentation?:
    | Documentation
    | undefined;
  /** API backend configuration. */
  backend?:
    | Backend
    | undefined;
  /** HTTP configuration. */
  http?:
    | Http
    | undefined;
  /** Quota configuration. */
  quota?:
    | Quota
    | undefined;
  /** Auth configuration. */
  authentication?:
    | Authentication
    | undefined;
  /** Context configuration. */
  context?:
    | Context
    | undefined;
  /** Configuration controlling usage of this service. */
  usage?:
    | Usage
    | undefined;
  /**
   * Configuration for network endpoints.  If this is empty, then an endpoint
   * with the same name as the service is automatically generated to service all
   * defined APIs.
   */
  endpoints: Endpoint[];
  /** Configuration for the service control plane. */
  control?:
    | Control
    | undefined;
  /** Defines the logs used by this service. */
  logs: LogDescriptor[];
  /** Defines the metrics used by this service. */
  metrics: MetricDescriptor[];
  /**
   * Defines the monitored resources used by this service. This is required
   * by the [Service.monitoring][google.api.Service.monitoring] and
   * [Service.logging][google.api.Service.logging] configurations.
   */
  monitoredResources: MonitoredResourceDescriptor[];
  /** Billing configuration. */
  billing?:
    | Billing
    | undefined;
  /** Logging configuration. */
  logging?:
    | Logging
    | undefined;
  /** Monitoring configuration. */
  monitoring?:
    | Monitoring
    | undefined;
  /** System parameter configuration. */
  systemParameters?:
    | SystemParameters
    | undefined;
  /** Output only. The source information for this configuration if available. */
  sourceInfo?:
    | SourceInfo
    | undefined;
  /**
   * Settings for [Google Cloud Client
   * libraries](https://cloud.google.com/apis/docs/cloud-client-libraries)
   * generated from APIs defined as protocol buffers.
   */
  publishing?:
    | Publishing
    | undefined;
  /**
   * Obsolete. Do not use.
   *
   * This field has no semantic meaning. The service config compiler always
   * sets this field to `3`.
   */
  configVersion?: number | undefined;
}

function createBaseService(): Service {
  return {
    name: "",
    title: "",
    producerProjectId: "",
    id: "",
    apis: [],
    types: [],
    enums: [],
    documentation: undefined,
    backend: undefined,
    http: undefined,
    quota: undefined,
    authentication: undefined,
    context: undefined,
    usage: undefined,
    endpoints: [],
    control: undefined,
    logs: [],
    metrics: [],
    monitoredResources: [],
    billing: undefined,
    logging: undefined,
    monitoring: undefined,
    systemParameters: undefined,
    sourceInfo: undefined,
    publishing: undefined,
    configVersion: undefined,
  };
}

export const Service = {
  encode(message: Service, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.producerProjectId !== "") {
      writer.uint32(178).string(message.producerProjectId);
    }
    if (message.id !== "") {
      writer.uint32(266).string(message.id);
    }
    for (const v of message.apis) {
      Api.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.types) {
      Type.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.enums) {
      Enum.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.documentation !== undefined) {
      Documentation.encode(message.documentation, writer.uint32(50).fork()).ldelim();
    }
    if (message.backend !== undefined) {
      Backend.encode(message.backend, writer.uint32(66).fork()).ldelim();
    }
    if (message.http !== undefined) {
      Http.encode(message.http, writer.uint32(74).fork()).ldelim();
    }
    if (message.quota !== undefined) {
      Quota.encode(message.quota, writer.uint32(82).fork()).ldelim();
    }
    if (message.authentication !== undefined) {
      Authentication.encode(message.authentication, writer.uint32(90).fork()).ldelim();
    }
    if (message.context !== undefined) {
      Context.encode(message.context, writer.uint32(98).fork()).ldelim();
    }
    if (message.usage !== undefined) {
      Usage.encode(message.usage, writer.uint32(122).fork()).ldelim();
    }
    for (const v of message.endpoints) {
      Endpoint.encode(v!, writer.uint32(146).fork()).ldelim();
    }
    if (message.control !== undefined) {
      Control.encode(message.control, writer.uint32(170).fork()).ldelim();
    }
    for (const v of message.logs) {
      LogDescriptor.encode(v!, writer.uint32(186).fork()).ldelim();
    }
    for (const v of message.metrics) {
      MetricDescriptor.encode(v!, writer.uint32(194).fork()).ldelim();
    }
    for (const v of message.monitoredResources) {
      MonitoredResourceDescriptor.encode(v!, writer.uint32(202).fork()).ldelim();
    }
    if (message.billing !== undefined) {
      Billing.encode(message.billing, writer.uint32(210).fork()).ldelim();
    }
    if (message.logging !== undefined) {
      Logging.encode(message.logging, writer.uint32(218).fork()).ldelim();
    }
    if (message.monitoring !== undefined) {
      Monitoring.encode(message.monitoring, writer.uint32(226).fork()).ldelim();
    }
    if (message.systemParameters !== undefined) {
      SystemParameters.encode(message.systemParameters, writer.uint32(234).fork()).ldelim();
    }
    if (message.sourceInfo !== undefined) {
      SourceInfo.encode(message.sourceInfo, writer.uint32(298).fork()).ldelim();
    }
    if (message.publishing !== undefined) {
      Publishing.encode(message.publishing, writer.uint32(362).fork()).ldelim();
    }
    if (message.configVersion !== undefined) {
      UInt32Value.encode({ value: message.configVersion! }, writer.uint32(162).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Service {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseService();
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
          if (tag !== 18) {
            break;
          }

          message.title = reader.string();
          continue;
        case 22:
          if (tag !== 178) {
            break;
          }

          message.producerProjectId = reader.string();
          continue;
        case 33:
          if (tag !== 266) {
            break;
          }

          message.id = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.apis.push(Api.decode(reader, reader.uint32()));
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.types.push(Type.decode(reader, reader.uint32()));
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.enums.push(Enum.decode(reader, reader.uint32()));
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.documentation = Documentation.decode(reader, reader.uint32());
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.backend = Backend.decode(reader, reader.uint32());
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.http = Http.decode(reader, reader.uint32());
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.quota = Quota.decode(reader, reader.uint32());
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.authentication = Authentication.decode(reader, reader.uint32());
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.context = Context.decode(reader, reader.uint32());
          continue;
        case 15:
          if (tag !== 122) {
            break;
          }

          message.usage = Usage.decode(reader, reader.uint32());
          continue;
        case 18:
          if (tag !== 146) {
            break;
          }

          message.endpoints.push(Endpoint.decode(reader, reader.uint32()));
          continue;
        case 21:
          if (tag !== 170) {
            break;
          }

          message.control = Control.decode(reader, reader.uint32());
          continue;
        case 23:
          if (tag !== 186) {
            break;
          }

          message.logs.push(LogDescriptor.decode(reader, reader.uint32()));
          continue;
        case 24:
          if (tag !== 194) {
            break;
          }

          message.metrics.push(MetricDescriptor.decode(reader, reader.uint32()));
          continue;
        case 25:
          if (tag !== 202) {
            break;
          }

          message.monitoredResources.push(MonitoredResourceDescriptor.decode(reader, reader.uint32()));
          continue;
        case 26:
          if (tag !== 210) {
            break;
          }

          message.billing = Billing.decode(reader, reader.uint32());
          continue;
        case 27:
          if (tag !== 218) {
            break;
          }

          message.logging = Logging.decode(reader, reader.uint32());
          continue;
        case 28:
          if (tag !== 226) {
            break;
          }

          message.monitoring = Monitoring.decode(reader, reader.uint32());
          continue;
        case 29:
          if (tag !== 234) {
            break;
          }

          message.systemParameters = SystemParameters.decode(reader, reader.uint32());
          continue;
        case 37:
          if (tag !== 298) {
            break;
          }

          message.sourceInfo = SourceInfo.decode(reader, reader.uint32());
          continue;
        case 45:
          if (tag !== 362) {
            break;
          }

          message.publishing = Publishing.decode(reader, reader.uint32());
          continue;
        case 20:
          if (tag !== 162) {
            break;
          }

          message.configVersion = UInt32Value.decode(reader, reader.uint32()).value;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Service {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      title: isSet(object.title) ? String(object.title) : "",
      producerProjectId: isSet(object.producerProjectId) ? String(object.producerProjectId) : "",
      id: isSet(object.id) ? String(object.id) : "",
      apis: Array.isArray(object?.apis) ? object.apis.map((e: any) => Api.fromJSON(e)) : [],
      types: Array.isArray(object?.types) ? object.types.map((e: any) => Type.fromJSON(e)) : [],
      enums: Array.isArray(object?.enums) ? object.enums.map((e: any) => Enum.fromJSON(e)) : [],
      documentation: isSet(object.documentation) ? Documentation.fromJSON(object.documentation) : undefined,
      backend: isSet(object.backend) ? Backend.fromJSON(object.backend) : undefined,
      http: isSet(object.http) ? Http.fromJSON(object.http) : undefined,
      quota: isSet(object.quota) ? Quota.fromJSON(object.quota) : undefined,
      authentication: isSet(object.authentication) ? Authentication.fromJSON(object.authentication) : undefined,
      context: isSet(object.context) ? Context.fromJSON(object.context) : undefined,
      usage: isSet(object.usage) ? Usage.fromJSON(object.usage) : undefined,
      endpoints: Array.isArray(object?.endpoints) ? object.endpoints.map((e: any) => Endpoint.fromJSON(e)) : [],
      control: isSet(object.control) ? Control.fromJSON(object.control) : undefined,
      logs: Array.isArray(object?.logs) ? object.logs.map((e: any) => LogDescriptor.fromJSON(e)) : [],
      metrics: Array.isArray(object?.metrics) ? object.metrics.map((e: any) => MetricDescriptor.fromJSON(e)) : [],
      monitoredResources: Array.isArray(object?.monitoredResources)
        ? object.monitoredResources.map((e: any) => MonitoredResourceDescriptor.fromJSON(e))
        : [],
      billing: isSet(object.billing) ? Billing.fromJSON(object.billing) : undefined,
      logging: isSet(object.logging) ? Logging.fromJSON(object.logging) : undefined,
      monitoring: isSet(object.monitoring) ? Monitoring.fromJSON(object.monitoring) : undefined,
      systemParameters: isSet(object.systemParameters) ? SystemParameters.fromJSON(object.systemParameters) : undefined,
      sourceInfo: isSet(object.sourceInfo) ? SourceInfo.fromJSON(object.sourceInfo) : undefined,
      publishing: isSet(object.publishing) ? Publishing.fromJSON(object.publishing) : undefined,
      configVersion: isSet(object.configVersion) ? Number(object.configVersion) : undefined,
    };
  },

  toJSON(message: Service): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.title !== undefined && (obj.title = message.title);
    message.producerProjectId !== undefined && (obj.producerProjectId = message.producerProjectId);
    message.id !== undefined && (obj.id = message.id);
    if (message.apis) {
      obj.apis = message.apis.map((e) => e ? Api.toJSON(e) : undefined);
    } else {
      obj.apis = [];
    }
    if (message.types) {
      obj.types = message.types.map((e) => e ? Type.toJSON(e) : undefined);
    } else {
      obj.types = [];
    }
    if (message.enums) {
      obj.enums = message.enums.map((e) => e ? Enum.toJSON(e) : undefined);
    } else {
      obj.enums = [];
    }
    message.documentation !== undefined &&
      (obj.documentation = message.documentation ? Documentation.toJSON(message.documentation) : undefined);
    message.backend !== undefined && (obj.backend = message.backend ? Backend.toJSON(message.backend) : undefined);
    message.http !== undefined && (obj.http = message.http ? Http.toJSON(message.http) : undefined);
    message.quota !== undefined && (obj.quota = message.quota ? Quota.toJSON(message.quota) : undefined);
    message.authentication !== undefined &&
      (obj.authentication = message.authentication ? Authentication.toJSON(message.authentication) : undefined);
    message.context !== undefined && (obj.context = message.context ? Context.toJSON(message.context) : undefined);
    message.usage !== undefined && (obj.usage = message.usage ? Usage.toJSON(message.usage) : undefined);
    if (message.endpoints) {
      obj.endpoints = message.endpoints.map((e) => e ? Endpoint.toJSON(e) : undefined);
    } else {
      obj.endpoints = [];
    }
    message.control !== undefined && (obj.control = message.control ? Control.toJSON(message.control) : undefined);
    if (message.logs) {
      obj.logs = message.logs.map((e) => e ? LogDescriptor.toJSON(e) : undefined);
    } else {
      obj.logs = [];
    }
    if (message.metrics) {
      obj.metrics = message.metrics.map((e) => e ? MetricDescriptor.toJSON(e) : undefined);
    } else {
      obj.metrics = [];
    }
    if (message.monitoredResources) {
      obj.monitoredResources = message.monitoredResources.map((e) =>
        e ? MonitoredResourceDescriptor.toJSON(e) : undefined
      );
    } else {
      obj.monitoredResources = [];
    }
    message.billing !== undefined && (obj.billing = message.billing ? Billing.toJSON(message.billing) : undefined);
    message.logging !== undefined && (obj.logging = message.logging ? Logging.toJSON(message.logging) : undefined);
    message.monitoring !== undefined &&
      (obj.monitoring = message.monitoring ? Monitoring.toJSON(message.monitoring) : undefined);
    message.systemParameters !== undefined &&
      (obj.systemParameters = message.systemParameters ? SystemParameters.toJSON(message.systemParameters) : undefined);
    message.sourceInfo !== undefined &&
      (obj.sourceInfo = message.sourceInfo ? SourceInfo.toJSON(message.sourceInfo) : undefined);
    message.publishing !== undefined &&
      (obj.publishing = message.publishing ? Publishing.toJSON(message.publishing) : undefined);
    message.configVersion !== undefined && (obj.configVersion = message.configVersion);
    return obj;
  },

  create(base?: DeepPartial<Service>): Service {
    return Service.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<Service>): Service {
    const message = createBaseService();
    message.name = object.name ?? "";
    message.title = object.title ?? "";
    message.producerProjectId = object.producerProjectId ?? "";
    message.id = object.id ?? "";
    message.apis = object.apis?.map((e) => Api.fromPartial(e)) || [];
    message.types = object.types?.map((e) => Type.fromPartial(e)) || [];
    message.enums = object.enums?.map((e) => Enum.fromPartial(e)) || [];
    message.documentation = (object.documentation !== undefined && object.documentation !== null)
      ? Documentation.fromPartial(object.documentation)
      : undefined;
    message.backend = (object.backend !== undefined && object.backend !== null)
      ? Backend.fromPartial(object.backend)
      : undefined;
    message.http = (object.http !== undefined && object.http !== null) ? Http.fromPartial(object.http) : undefined;
    message.quota = (object.quota !== undefined && object.quota !== null) ? Quota.fromPartial(object.quota) : undefined;
    message.authentication = (object.authentication !== undefined && object.authentication !== null)
      ? Authentication.fromPartial(object.authentication)
      : undefined;
    message.context = (object.context !== undefined && object.context !== null)
      ? Context.fromPartial(object.context)
      : undefined;
    message.usage = (object.usage !== undefined && object.usage !== null) ? Usage.fromPartial(object.usage) : undefined;
    message.endpoints = object.endpoints?.map((e) => Endpoint.fromPartial(e)) || [];
    message.control = (object.control !== undefined && object.control !== null)
      ? Control.fromPartial(object.control)
      : undefined;
    message.logs = object.logs?.map((e) => LogDescriptor.fromPartial(e)) || [];
    message.metrics = object.metrics?.map((e) => MetricDescriptor.fromPartial(e)) || [];
    message.monitoredResources = object.monitoredResources?.map((e) => MonitoredResourceDescriptor.fromPartial(e)) ||
      [];
    message.billing = (object.billing !== undefined && object.billing !== null)
      ? Billing.fromPartial(object.billing)
      : undefined;
    message.logging = (object.logging !== undefined && object.logging !== null)
      ? Logging.fromPartial(object.logging)
      : undefined;
    message.monitoring = (object.monitoring !== undefined && object.monitoring !== null)
      ? Monitoring.fromPartial(object.monitoring)
      : undefined;
    message.systemParameters = (object.systemParameters !== undefined && object.systemParameters !== null)
      ? SystemParameters.fromPartial(object.systemParameters)
      : undefined;
    message.sourceInfo = (object.sourceInfo !== undefined && object.sourceInfo !== null)
      ? SourceInfo.fromPartial(object.sourceInfo)
      : undefined;
    message.publishing = (object.publishing !== undefined && object.publishing !== null)
      ? Publishing.fromPartial(object.publishing)
      : undefined;
    message.configVersion = object.configVersion ?? undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
