import { SendIcon } from "@assets/icons/send";
import { MessageRole } from "@connect/arukabe/models/v1/message_pb";
import { AIMessage } from "@features/chats/components/AIMessage/AIMessage";
import { UserMessage } from "@features/chats/components/UserMessage/UserMesage";
import { useChatMessages } from "@features/chats/data/use-chat-messages";
import { useChatMessagesStore } from "@features/chats/store/chat-messages-store";
import { useFirstMessageStore } from "@features/chats/store/useFirstMessageStore";
import { createTextMessage } from "@features/chats/utils/create_text_message";
import {
  ActionIcon,
  Alert,
  Box,
  Button,
  Card,
  Container,
  Flex,
  Group,
  Textarea,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { Fragment, useCallback, useEffect } from "react";
import styles from "./styles.module.css";
import { useChatScroll } from "./use-chat-scroll";
import { useStreamingResponse } from "./use-streaming-response";
import { INPUT_BOTTOM_DISTANCE } from "./utils";

interface Props {
  chatId: string;
}

export function Chat(props: Props) {
  const { chatId } = props;
  const { firstMessage, clearFirstMessage } = useFirstMessageStore();
  const { messages, isFetchedMessages } = useChatMessages(chatId);
  const addMessage = useChatMessagesStore((s) => s.addMessage);
  const { generatedResponse, responseError, sendMessage, commitResponse, clearError } =
    useStreamingResponse(chatId);
  const { containerRef, messagesRef, inputContainerRef, scrollToNewMessage } = useChatScroll({
    chatId,
    isFetchedMessages,
    messagesCount: messages.length,
  });

  const form = useForm({
    initialValues: { userText: "" },
    validate: {
      userText: (value) => (value ? null : "Please enter a message"),
    },
  });

  const retry = () => {
    if (messages.length <= 0) return;
    if (!responseError) return;
    const lastMessage = messages[messages.length - 1];
    const messageStr = lastMessage.content
      .filter((c) => c.content.case === "textContent")
      .map((c) => c.content.value?.text ?? "")
      .join("");
    clearError();
    sendMessage(messageStr);
  };

  const handleSubmit = useCallback(
    (values: typeof form.values) => {
      commitResponse();
      addMessage(chatId, createTextMessage(values.userText, MessageRole.USER));
      sendMessage(values.userText);
      form.setValues({ userText: "" });
      scrollToNewMessage();
    },
    [addMessage, chatId, commitResponse, form, sendMessage, scrollToNewMessage],
  );

  // Auto-send message when coming from /new
  useEffect(() => {
    if (!firstMessage || !isFetchedMessages) return;
    handleSubmit({ userText: firstMessage });
    clearFirstMessage();
  }, [firstMessage, clearFirstMessage, handleSubmit, isFetchedMessages]);

  return (
    <Box px="md" className={styles.container} ref={containerRef} id="ccc">
      <Flex ref={messagesRef} direction="column" gap="lg" pt="xl">
        {messages.map((message) => (
          <Fragment key={message.id}>
            {message.role === MessageRole.USER && <UserMessage content={message.content} />}
            {message.role === MessageRole.ASSISTANT && <AIMessage content={message.content} />}
          </Fragment>
        ))}
      </Flex>
      <Box mih="110dvh" mt="xl">
        {generatedResponse && <AIMessage content={generatedResponse?.content} />}
        {responseError && (
          <Container size="md">
            <Alert color="red" title="Response error">
              {responseError}
              <Flex justify="flex-end" pb="xs">
                <Button variant="outline" size="sm" color="gray" onClick={retry}>
                  Retry
                </Button>
              </Flex>
            </Alert>
          </Container>
        )}
        <Box h={120} />
      </Box>

      <Box className={styles.cardContainer} bottom={INPUT_BOTTOM_DISTANCE} ref={inputContainerRef}>
        <Container>
          <form onSubmit={form.onSubmit(handleSubmit)}>
            <Card shadow="sm" className={styles.card}>
              <Textarea
                variant="unstyled"
                placeholder="Write something..."
                size="md"
                minRows={1}
                maxRows={8}
                onKeyDown={(e) => {
                  if (e.key === "Enter" && e.ctrlKey) {
                    e.preventDefault();
                    form.onSubmit(handleSubmit)();
                  }
                }}
                {...form.getInputProps("userText")}
              />
              <Group justify="flex-end">
                <ActionIcon size="xl" radius="sm" type="submit">
                  <SendIcon width={24} height={24} />
                </ActionIcon>
              </Group>
            </Card>
          </form>
        </Container>
      </Box>
    </Box>
  );
}
