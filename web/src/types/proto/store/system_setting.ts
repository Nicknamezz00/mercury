/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "mercury.store";

export enum SystemSettingKey {
  SYSTEM_SETTING_KEY_UNSPECIFIED = 0,
  BACKUP_CONFIG = 1,
  UNRECOGNIZED = -1,
}

export function systemSettingKeyFromJSON(object: any): SystemSettingKey {
  switch (object) {
    case 0:
    case "SYSTEM_SETTING_KEY_UNSPECIFIED":
      return SystemSettingKey.SYSTEM_SETTING_KEY_UNSPECIFIED;
    case 1:
    case "BACKUP_CONFIG":
      return SystemSettingKey.BACKUP_CONFIG;
    case -1:
    case "UNRECOGNIZED":
    default:
      return SystemSettingKey.UNRECOGNIZED;
  }
}

export function systemSettingKeyToJSON(object: SystemSettingKey): string {
  switch (object) {
    case SystemSettingKey.SYSTEM_SETTING_KEY_UNSPECIFIED:
      return "SYSTEM_SETTING_KEY_UNSPECIFIED";
    case SystemSettingKey.BACKUP_CONFIG:
      return "BACKUP_CONFIG";
    case SystemSettingKey.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface BackupConfig {
  enabled: boolean;
  /**
   * For more information about cron expression
   * see: https://pkg.go.dev/github.com/robfig/cron?utm_source=godoc#hdr-CRON_Expression_Format
   */
  cron: string;
  /** number of backups to keep. */
  maxKeep: number;
}

function createBaseBackupConfig(): BackupConfig {
  return { enabled: false, cron: "", maxKeep: 0 };
}

export const BackupConfig = {
  encode(message: BackupConfig, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.enabled === true) {
      writer.uint32(8).bool(message.enabled);
    }
    if (message.cron !== "") {
      writer.uint32(18).string(message.cron);
    }
    if (message.maxKeep !== 0) {
      writer.uint32(24).int32(message.maxKeep);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BackupConfig {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBackupConfig();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.enabled = reader.bool();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.cron = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.maxKeep = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): BackupConfig {
    return {
      enabled: isSet(object.enabled) ? Boolean(object.enabled) : false,
      cron: isSet(object.cron) ? String(object.cron) : "",
      maxKeep: isSet(object.maxKeep) ? Number(object.maxKeep) : 0,
    };
  },

  toJSON(message: BackupConfig): unknown {
    const obj: any = {};
    message.enabled !== undefined && (obj.enabled = message.enabled);
    message.cron !== undefined && (obj.cron = message.cron);
    message.maxKeep !== undefined && (obj.maxKeep = Math.round(message.maxKeep));
    return obj;
  },

  create(base?: DeepPartial<BackupConfig>): BackupConfig {
    return BackupConfig.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<BackupConfig>): BackupConfig {
    const message = createBaseBackupConfig();
    message.enabled = object.enabled ?? false;
    message.cron = object.cron ?? "";
    message.maxKeep = object.maxKeep ?? 0;
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
