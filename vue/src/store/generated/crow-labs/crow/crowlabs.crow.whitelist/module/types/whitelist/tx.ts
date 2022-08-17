/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "crowlabs.crow.whitelist";

export interface MsgWhitelistUser {
  creator: string;
  address: string;
}

export interface MsgWhitelistUserResponse {}

export interface MsgWhitelistProducer {
  creator: string;
  address: string;
}

export interface MsgWhitelistProducerResponse {}

const baseMsgWhitelistUser: object = { creator: "", address: "" };

export const MsgWhitelistUser = {
  encode(message: MsgWhitelistUser, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWhitelistUser {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWhitelistUser } as MsgWhitelistUser;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWhitelistUser {
    const message = { ...baseMsgWhitelistUser } as MsgWhitelistUser;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: MsgWhitelistUser): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgWhitelistUser>): MsgWhitelistUser {
    const message = { ...baseMsgWhitelistUser } as MsgWhitelistUser;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseMsgWhitelistUserResponse: object = {};

export const MsgWhitelistUserResponse = {
  encode(
    _: MsgWhitelistUserResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgWhitelistUserResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWhitelistUserResponse,
    } as MsgWhitelistUserResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgWhitelistUserResponse {
    const message = {
      ...baseMsgWhitelistUserResponse,
    } as MsgWhitelistUserResponse;
    return message;
  },

  toJSON(_: MsgWhitelistUserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgWhitelistUserResponse>
  ): MsgWhitelistUserResponse {
    const message = {
      ...baseMsgWhitelistUserResponse,
    } as MsgWhitelistUserResponse;
    return message;
  },
};

const baseMsgWhitelistProducer: object = { creator: "", address: "" };

export const MsgWhitelistProducer = {
  encode(
    message: MsgWhitelistProducer,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWhitelistProducer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWhitelistProducer } as MsgWhitelistProducer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWhitelistProducer {
    const message = { ...baseMsgWhitelistProducer } as MsgWhitelistProducer;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: MsgWhitelistProducer): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgWhitelistProducer>): MsgWhitelistProducer {
    const message = { ...baseMsgWhitelistProducer } as MsgWhitelistProducer;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseMsgWhitelistProducerResponse: object = {};

export const MsgWhitelistProducerResponse = {
  encode(
    _: MsgWhitelistProducerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgWhitelistProducerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWhitelistProducerResponse,
    } as MsgWhitelistProducerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgWhitelistProducerResponse {
    const message = {
      ...baseMsgWhitelistProducerResponse,
    } as MsgWhitelistProducerResponse;
    return message;
  },

  toJSON(_: MsgWhitelistProducerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgWhitelistProducerResponse>
  ): MsgWhitelistProducerResponse {
    const message = {
      ...baseMsgWhitelistProducerResponse,
    } as MsgWhitelistProducerResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  WhitelistUser(request: MsgWhitelistUser): Promise<MsgWhitelistUserResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  WhitelistProducer(
    request: MsgWhitelistProducer
  ): Promise<MsgWhitelistProducerResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  WhitelistUser(request: MsgWhitelistUser): Promise<MsgWhitelistUserResponse> {
    const data = MsgWhitelistUser.encode(request).finish();
    const promise = this.rpc.request(
      "crowlabs.crow.whitelist.Msg",
      "WhitelistUser",
      data
    );
    return promise.then((data) =>
      MsgWhitelistUserResponse.decode(new Reader(data))
    );
  }

  WhitelistProducer(
    request: MsgWhitelistProducer
  ): Promise<MsgWhitelistProducerResponse> {
    const data = MsgWhitelistProducer.encode(request).finish();
    const promise = this.rpc.request(
      "crowlabs.crow.whitelist.Msg",
      "WhitelistProducer",
      data
    );
    return promise.then((data) =>
      MsgWhitelistProducerResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
