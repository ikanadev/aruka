import { Alert, Box, Container } from "@mantine/core";
import { Fragment } from "react";
import { useDisclosure } from "@mantine/hooks";

import type { MessageContent } from "@connect/models/v1/message_content_pb";
import { MessageActions } from "@features/chats/components/MessageActions/MessageActions";

interface Props {
  content: MessageContent[];
}

export function UserMessage(props: Props) {
  const { content } = props;
  const [expanded, { toggle }] = useDisclosure(false);

  return (
    <Box>
      <Container size={expanded ? "xl" : "md"} mb="xs">
        <Alert variant="default">
          {content.map((msg, index) => (
            // biome-ignore lint/suspicious/noArrayIndexKey: not critical
            <Fragment key={index}>
              {msg.content.case === "textContent" && <p>{msg.content.value.text}</p>}
              {/* Handle more message types */}
            </Fragment>
          ))}
        </Alert>
      </Container>
      <MessageActions toggleExpand={toggle} />
    </Box>
  );
}
