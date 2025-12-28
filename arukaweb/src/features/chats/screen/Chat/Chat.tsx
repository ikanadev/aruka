import { SendIcon } from "@assets/icons/send";
import { chatClient } from "@common/utils/clients";
import type { ChatMessagesResponse } from "@connect/chat/v1/chat_messages_pb";
import { type Message, MessageRole } from "@connect/models/v1/message_pb";
import { AIMessage } from "@features/chats/components/AIMessage/AIMessage";
import { UserMessage } from "@features/chats/components/UserMessage/UserMesage";
import { chatQueryKeys } from "@features/chats/data/chat-query-keys";
import { useChatMessages } from "@features/chats/data/use-chat-messages";
import { useFirstMessageStore } from "@features/chats/store/useFirstMessageStore";
import { createTextMessage } from "@features/chats/utils/create_text_message";
import {
  ActionIcon,
  Box,
  Card,
  Container,
  Flex,
  Group,
  Textarea,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { useQueryClient } from "@tanstack/react-query";
import { Fragment, useEffect, useLayoutEffect, useRef, useState } from "react";
import styles from "./styles.module.css";
import { INPUT_BOTTOM_DISTANCE } from "./utils";
import { useRouter } from "@tanstack/react-router";

export const chatsScrollMap: Record<string, number> = {};

interface Props {
  chatId: string;
}

export function Chat(props: Props) {
  const { chatId } = props;
  const router = useRouter();
  const queryClient = useQueryClient();
  const inputContainerRef = useRef<HTMLDivElement>(null);
  const messagesRef = useRef<HTMLDivElement>(null);
  const containerRef = useRef<HTMLDivElement>(null);
  const { firstMessage, clearFirstMessage } = useFirstMessageStore();
  const { messages, isFetchedMessages } = useChatMessages(chatId);
  const [generatedResponse, setGeneratedResponse] = useState<Message | null>(
    null,
  );

  const form = useForm({
    initialValues: { userText: firstMessage || "" },
    validate: {
      userText: (value) => (value ? null : "Please enter a message"),
    },
  });

  const newMessageScroll = () => {
    if (!messagesRef.current || !containerRef.current) return;
    containerRef.current.scrollTo({
      top: messagesRef.current.scrollHeight,
      behavior: "instant",
    });
  };

  const handleNewMessage = async (message: string) => {
    const res = chatClient.chatMessage({
      chatId,
      content: [{ content: { case: "textContent", value: { text: message } } }],
    });
    let responseText = "";
    for await (const delta of res) {
      responseText += delta.delta;
      setGeneratedResponse(
        createTextMessage(responseText, MessageRole.ASSISTANT),
      );
    }
  };

  const handleSubmit = (values: typeof form.values) => {
    queryClient.setQueryData(
      chatQueryKeys.chatMessages(chatId),
      (prev: ChatMessagesResponse | undefined): ChatMessagesResponse => {
        if (prev === undefined) {
          return {
            $typeName: "chat.v1.ChatMessagesResponse",
            messages: [createTextMessage(values.userText, MessageRole.USER)],
          };
        }
        const newMessages = [...prev.messages];
        if (generatedResponse) {
          newMessages.push(generatedResponse);
          setGeneratedResponse(null);
        }
        newMessages.push(createTextMessage(values.userText, MessageRole.USER));
        return {
          ...prev,
          messages: newMessages,
        };
      },
    );
    newMessageScroll();
    handleNewMessage(values.userText);
    form.setValues({ userText: "" });
  };

  // Auto-send message when we're comming from a /new chat
  useEffect(() => {
    if (firstMessage === null || !isFetchedMessages) return;
    handleSubmit(form.values);
    clearFirstMessage();
  }, [
    firstMessage,
    clearFirstMessage,
    form.values,
    handleSubmit,
    isFetchedMessages,
  ]);

  // Auto-scroll to saved position when changing chats
  useLayoutEffect(() => {
    if (!(chatId in chatsScrollMap)) return;
    if (!containerRef.current) return;
    containerRef.current.scrollTo({
      top: chatsScrollMap[chatId],
      behavior: "instant",
    });
  }, [chatId]);

  // Auto-scroll to bottom on initial chat load
  useLayoutEffect(() => {
    if (!isFetchedMessages) return;
    if (
      !containerRef.current ||
      !messagesRef.current ||
      !inputContainerRef.current
    )
      return;
    if (chatId in chatsScrollMap) return;
    containerRef.current.scrollTo({
      top:
        messagesRef.current.scrollHeight -
        window.innerHeight +
        inputContainerRef.current.clientHeight +
        INPUT_BOTTOM_DISTANCE +
        12,
      behavior: "instant",
    });
  }, [chatId, isFetchedMessages]);

  // Save last chat to query cache
  useEffect(() => {
    const unsubscribe = router.subscribe("onBeforeNavigate", () => {
      if (generatedResponse === null) return;
      queryClient.setQueryData(
        chatQueryKeys.chatMessages(chatId),
        (
          prev: ChatMessagesResponse | undefined,
        ): ChatMessagesResponse | undefined => {
          if (prev === undefined) return undefined;
          return {
            ...prev,
            messages: [...prev.messages, generatedResponse],
          };
        },
      );
      setGeneratedResponse(null);
    });
    return unsubscribe;
  }, [chatId, generatedResponse]);

  // Listen to scroll changes and save it to restore position later
  useEffect(() => {
    const el = containerRef.current;
    if (!el) return;
    const handleScrollEnd = () => {
      chatsScrollMap[chatId] = el.scrollTop;
    };
    el.addEventListener("scrollend", handleScrollEnd);
    return () => el.removeEventListener("scrollend", handleScrollEnd);
  }, [chatId]);

  return (
    <Box px="md" className={styles.container} ref={containerRef}>
      <Flex ref={messagesRef} direction="column" gap="lg" pt="xl">
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
      </Flex>
      <Box mih="100dvh">
        {generatedResponse && (
          <AIMessage content={generatedResponse?.content} />
        )}
        <Box h={120} />
      </Box>

      <Box
        className={styles.cardContainer}
        bottom={INPUT_BOTTOM_DISTANCE}
        ref={inputContainerRef}
      >
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
