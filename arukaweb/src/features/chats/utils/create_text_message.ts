import { generateId } from "@common/utils/generate_id";
import type { Message, MessageRole } from "@connect/arukabe/models/v1/message_pb";

export function createTextMessage(message: string, role: MessageRole): Message {
  return {
    $typeName: "models.v1.Message",
    id: generateId(),
    content: [
      {
        $typeName: "models.v1.MessageContent",
        content: {
          case: "textContent",
          value: {
            $typeName: "models.v1.MessageTextContent",
            text: message,
          },
        },
      },
    ],
    role,
    createdAt: {
      $typeName: "google.protobuf.Timestamp",
      seconds: BigInt(Math.round(Date.now() / 1000)),
      nanos: 0,
    },
  };
}
