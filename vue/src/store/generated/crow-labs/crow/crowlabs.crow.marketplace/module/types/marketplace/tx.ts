/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "crowlabs.crow.marketplace";

export interface MsgCreateListing {
  creator: string;
  title: string;
  description: string;
  minPrice: string;
}

export interface MsgCreateListingResponse {}

const baseMsgCreateListing: object = {
  creator: "",
  title: "",
  description: "",
  minPrice: "",
};

export const MsgCreateListing = {
  encode(message: MsgCreateListing, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.minPrice !== "") {
      writer.uint32(34).string(message.minPrice);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateListing {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateListing } as MsgCreateListing;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.minPrice = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateListing {
    const message = { ...baseMsgCreateListing } as MsgCreateListing;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.minPrice !== undefined && object.minPrice !== null) {
      message.minPrice = String(object.minPrice);
    } else {
      message.minPrice = "";
    }
    return message;
  },

  toJSON(message: MsgCreateListing): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.minPrice !== undefined && (obj.minPrice = message.minPrice);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateListing>): MsgCreateListing {
    const message = { ...baseMsgCreateListing } as MsgCreateListing;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.minPrice !== undefined && object.minPrice !== null) {
      message.minPrice = object.minPrice;
    } else {
      message.minPrice = "";
    }
    return message;
  },
};

const baseMsgCreateListingResponse: object = {};

export const MsgCreateListingResponse = {
  encode(
    _: MsgCreateListingResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateListingResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateListingResponse,
    } as MsgCreateListingResponse;
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

  fromJSON(_: any): MsgCreateListingResponse {
    const message = {
      ...baseMsgCreateListingResponse,
    } as MsgCreateListingResponse;
    return message;
  },

  toJSON(_: MsgCreateListingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateListingResponse>
  ): MsgCreateListingResponse {
    const message = {
      ...baseMsgCreateListingResponse,
    } as MsgCreateListingResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateListing(request: MsgCreateListing): Promise<MsgCreateListingResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateListing(request: MsgCreateListing): Promise<MsgCreateListingResponse> {
    const data = MsgCreateListing.encode(request).finish();
    const promise = this.rpc.request(
      "crowlabs.crow.marketplace.Msg",
      "CreateListing",
      data
    );
    return promise.then((data) =>
      MsgCreateListingResponse.decode(new Reader(data))
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
