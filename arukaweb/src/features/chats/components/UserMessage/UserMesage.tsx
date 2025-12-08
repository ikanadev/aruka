import type { MessageContent } from "@connect/models/v1/message_content_pb";
import { ActionIcon, Box, Container, Group } from "@mantine/core";
import { Fragment, useState } from "react";

interface Props {
  content: MessageContent[];
}

export function UserMessage(props: Props) {
  const { content } = props;
  const [expanded, setExpanded] = useState(false);
  return (
    <Box>
      <Container size={expanded ? "xl" : "md"}>
        <Box>
          {content.map((msg, index) => (
            // biome-ignore lint/suspicious/noArrayIndexKey: not critical
            <Fragment key={index}>
              {msg.content.case === "textContent" && (
                <p>{msg.content.value.text}</p>
              )}
              {/* Handle more message types */}
            </Fragment>
          ))}
        </Box>
      </Container>
      <Container size="md">
        <Group justify="flex-end">
          <ActionIcon onClick={() => setExpanded((prev) => !prev)}>
            X
          </ActionIcon>
        </Group>
      </Container>
    </Box>
  );
}
