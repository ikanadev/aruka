import { ActionIcon, Box, Card, Container, Flex, Group, Textarea } from "@mantine/core";

import styles from './styles.module.css';
import { SendIcon } from "@assets/icons/send";
import { useChatMessages } from "@features/chats/data/use-chat-messages";
import { chatClient } from "@common/utils/clients";
import { useFirstMessageStore } from "@features/chats/store/useFirstMessageStore";
import { useCallback, useEffect } from "react";

interface Props {
  chatId: string;
}

export function Chat(props: Props) {
  const { chatId } = props;
  const { messages } = useChatMessages(chatId);
  const { firstMessage, clearFirstMessage } = useFirstMessageStore();

  const handleNewMessage = useCallback(async (message: string) => {
    const res = chatClient.chatMessage({
      chatId,
      content: [{ content: { case: "textContent", value: { text: message } } }],
    });
    for await (let delta of res) {
      console.log(delta);
    }
  }, [chatId]);

  console.log(messages);

  useEffect(() => {
    if (firstMessage === null) return;
    handleNewMessage(firstMessage).then(() => {
      clearFirstMessage();
    });
  }, [firstMessage, handleNewMessage, clearFirstMessage]);

  return (
    <Flex px="md" className={styles.container}>
      <Box className={styles.cardContainer}>
        <Container>
          <Card shadow="sm" className={styles.card}>
            <Textarea variant="unstyled" placeholder="Write something..." size="md" minRows={1} maxRows={8} />
            <Group justify="flex-end">
              <ActionIcon size="xl" radius="sm" type="submit">
                <SendIcon width={24} height={24} />
              </ActionIcon>
            </Group>
          </Card>
        </Container>
      </Box>
    </Flex>
  );
}
