import { Box, Container, Typography } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { Fragment } from "react";

import type { MessageContent } from "@connect/arukabe/models/v1/message_content_pb";
import { parser } from "@features/chats/utils/parser";
import { MessageActions } from "@features/chats/components/MessageActions/MessageActions";

interface Props {
  content: MessageContent[];
}

export function AIMessage(props: Props) {
  const { content } = props;
  const [expanded, { toggle }] = useDisclosure(false);

  return (
    <Box>
      <Container size={expanded ? "xl" : "md"} mb="xs">
        <Box>
          {content.map((msg, index) => (
            // biome-ignore lint/suspicious/noArrayIndexKey: not critical
            <Fragment key={index}>
              {msg.content.case === "textContent" && (
                <Typography>
                  <div
                    // biome-ignore lint/security/noDangerouslySetInnerHtml: markdown generated
                    dangerouslySetInnerHTML={{
                      __html: parser.render(msg.content.value.text),
                    }}
                  />
                </Typography>
              )}
              {/* Handle more message types */}
            </Fragment>
          ))}
        </Box>
      </Container>
      <MessageActions toggleExpand={toggle} />
    </Box>
  );
}
