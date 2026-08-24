import { Buffer } from "node:buffer";

export function encodeFrame(payload: string): Buffer {
  const body = Buffer.from(payload, "utf8");
  return Buffer.concat([
    Buffer.from(`Content-Length: ${body.length}\r\n\r\n`, "ascii"),
    body,
  ]);
}

export class FrameDecoder {
  private buffer = Buffer.alloc(0);

  push(chunk: Buffer): string[] {
    this.buffer = Buffer.concat([this.buffer, chunk]);
    const messages: string[] = [];
    while (true) {
      const separator = this.buffer.indexOf(Buffer.from("\r\n\r\n"));
      if (separator < 0) return messages;
      const header = this.buffer.subarray(0, separator).toString("ascii");
      const match = /^Content-Length:\s*(\d+)$/im.exec(header.trim());
      if (!match) throw new Error("invalid aosp-index frame header");
      const length = Number(match[1]);
      const start = separator + 4;
      if (this.buffer.length < start + length) return messages;
      messages.push(this.buffer.subarray(start, start + length).toString("utf8"));
      this.buffer = this.buffer.subarray(start + length);
    }
  }
}
