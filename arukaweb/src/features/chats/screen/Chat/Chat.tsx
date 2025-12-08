import { SendIcon } from "@assets/icons/send";
import { chatClient } from "@common/utils/clients";
import { useChatMessages } from "@features/chats/data/use-chat-messages";
import { useFirstMessageStore } from "@features/chats/store/useFirstMessageStore";
import {
  ActionIcon,
  Box,
  Card,
  Container,
  Group,
  Textarea,
} from "@mantine/core";
import { Fragment, useCallback, useEffect } from "react";
import styles from "./styles.module.css";
import { MessageRole, type Message } from "@connect/models/v1/message_pb";
import { UserMessage } from "@features/chats/components/UserMessage/UserMesage";
import { AIMessage } from "@features/chats/components/AIMessage/AIMessage";
import { useForm } from "@mantine/form";
import { useQueryClient } from "@tanstack/react-query";
import { chatQueryKeys } from "@features/chats/data/chat-query-keys";
import type { ChatMessageResponse } from "@connect/chat/v1/chat_message_pb";
import type { ChatMessagesResponse } from "@connect/chat/v1/chat_messages_pb";

interface Props {
  chatId: string;
}

export function Chat(props: Props) {
  const { chatId } = props;
  const queryClient = useQueryClient();
  const { messages } = useChatMessages(chatId);
  const { firstMessage, clearFirstMessage } = useFirstMessageStore();
  const form = useForm({
    mode: "uncontrolled",
    initialValues: { userText: "" },
    validate: {
      userText: (value) => (value ? null : "Please enter a message"),
    },
  });

  const handleNewMessage = useCallback(
    async (message: string) => {
      const res = chatClient.chatMessage({
        chatId,
        content: [
          { content: { case: "textContent", value: { text: message } } },
        ],
      });
      let responseText = "";
      for await (const delta of res) {
        console.log(delta);
        responseText += delta.delta;
      }
      queryClient.setQueryData(
        chatQueryKeys.chatMessages(chatId),
        (prev: ChatMessagesResponse): ChatMessagesResponse => ({
          ...prev,
          messages: [
            ...prev.messages,
            {
              id: `${Date.now()}`,
              role: MessageRole.ASSISTANT,
              content: [
                {
                  $typeName: "models.v1.MessageContent",
                  content: {
                    case: "textContent",
                    value: {
                      text: responseText,
                      $typeName: "models.v1.MessageTextContent",
                    },
                  },
                },
              ],
              $typeName: "models.v1.Message",
            },
          ],
        }),
      );
    },
    [chatId, queryClient],
  );

  console.log(messages);

  const handleSubmit = (values: typeof form.values) => {
    queryClient.setQueryData(
      chatQueryKeys.chatMessages(chatId),
      (prev: ChatMessagesResponse): ChatMessagesResponse => ({
        ...prev,
        messages: [
          ...prev.messages,
          {
            id: `${Date.now()}`,
            role: MessageRole.USER,
            content: [
              {
                $typeName: "models.v1.MessageContent",
                content: {
                  case: "textContent",
                  value: {
                    text: values.userText,
                    $typeName: "models.v1.MessageTextContent",
                  },
                },
              },
            ],
            $typeName: "models.v1.Message",
          },
        ],
      }),
    );
    handleNewMessage(values.userText);
    form.reset();
  };

  useEffect(() => {
    if (firstMessage === null) return;
    console.log("firstMessage req:", firstMessage);
    handleNewMessage(firstMessage).then(() => {
      clearFirstMessage();
    });
  }, [firstMessage, handleNewMessage, clearFirstMessage]);

  return (
    <Box px="md" className={styles.container}>
      {messages.map((message) => (
        <Fragment key={message.id}>
          {message.role === MessageRole.USER && (
            <UserMessage content={message.content} />
          )}
          {message.role === MessageRole.ASSISTANT && (
            <AIMessage content={message.content} />
          )}
        </Fragment>
      ))}

      <Box className={styles.cardContainer}>
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
