import { SendIcon } from "@assets/icons/send";
import { chatClient } from "@common/utils/clients";
import { type Message, MessageRole } from "@connect/models/v1/message_pb";
import { AIMessage } from "@features/chats/components/AIMessage/AIMessage";
import { UserMessage } from "@features/chats/components/UserMessage/UserMesage";
import { useChatMessages } from "@features/chats/data/use-chat-messages";
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
import { Fragment, useCallback, useEffect, useLayoutEffect, useRef, useState } from "react";
import styles from "./styles.module.css";
import { INPUT_BOTTOM_DISTANCE } from "./utils";
import { useRouter } from "@tanstack/react-router";

const ScrollType = {
  LastMessage: "LastMessage",
  NewMessage: "NewMessage",
  RestoreChatPosition: "RestoreChatPosition",
} as const;
type ScrollType = (typeof ScrollType)[keyof typeof ScrollType];

const chatsScrollMap = new Map<string, number>();

interface Props {
  chatId: string;
}

export function Chat(props: Props) {
  const { chatId } = props;
  const router = useRouter();
  const inputContainerRef = useRef<HTMLDivElement>(null);
  const messagesRef = useRef<HTMLDivElement>(null);
  const containerRef = useRef<HTMLDivElement>(null);
  const { firstMessage, clearFirstMessage } = useFirstMessageStore();
  const { messages, addMessageToCache, isFetchedMessages } = useChatMessages(chatId);
  const [generatedResponse, setGeneratedResponse] = useState<Message | null>(null);
  const [responseError, setResponseError] = useState<string | null>(null);
  const [scrollType, setScrollType] = useState<ScrollType | null>(null);

  const form = useForm({
    initialValues: { userText: firstMessage || "" },
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
    setResponseError(null);
    handleNewMessage(messageStr);
  };

  const handleNewMessage = useCallback(
    async (message: string) => {
      try {
        const res = chatClient.chatMessage({
          chatId,
          content: [{ content: { case: "textContent", value: { text: message } } }],
        });
        let responseText = "";
        for await (const delta of res) {
          responseText += delta.delta;
          setGeneratedResponse(createTextMessage(responseText, MessageRole.ASSISTANT));
        }
      } catch (e) {
        setResponseError((e as Error)?.message ?? "Something went wrong");
      }
    },
    [chatId],
  );

  const handleSubmit = useCallback(
    (values: typeof form.values) => {
      addMessageToCache(createTextMessage(values.userText, MessageRole.USER));
      if (generatedResponse) {
        addMessageToCache(generatedResponse);
        setGeneratedResponse(null);
      }
      handleNewMessage(values.userText);
      form.setValues({ userText: "" });
      setScrollType(ScrollType.NewMessage);
    },
    [addMessageToCache, form, generatedResponse, handleNewMessage],
  );

  // Auto-send message when we're comming from a /new chat
  useEffect(() => {
    if (firstMessage === null || !isFetchedMessages) return;
    handleSubmit(form.values);
    clearFirstMessage();
  }, [firstMessage, clearFirstMessage, form.values, handleSubmit, isFetchedMessages]);

  // Auto-scroll to saved position when changing chats
  useEffect(() => {
    if (!chatsScrollMap.has(chatId)) return;
    console.log("Effect:", ScrollType.RestoreChatPosition);
    setScrollType(ScrollType.RestoreChatPosition);
  }, [chatId]);

  // Auto-scroll to bottom on initial chat load
  useEffect(() => {
    if (!isFetchedMessages) return;
    if (chatsScrollMap.has(chatId)) return;
    if (messages.length <= 1) return; // 1 means first user message
    console.log("Effect:", ScrollType.LastMessage);
    setScrollType(ScrollType.LastMessage);
  }, [chatId, isFetchedMessages, messages]);

  // Save last chat to query cache
  useEffect(() => {
    const unsubscribe = router.subscribe("onBeforeNavigate", () => {
      if (generatedResponse === null) return;
      addMessageToCache(generatedResponse);
      setGeneratedResponse(null);
    });
    return unsubscribe;
  }, [chatId, generatedResponse, router, addMessageToCache]);

  // Listen to scroll changes and save it to restore position later
  useEffect(() => {
    const el = containerRef.current;
    if (!el) return;
    const handleScrollEnd = () => chatsScrollMap.set(chatId, el.scrollTop);
    el.addEventListener("scrollend", handleScrollEnd);
    return () => el.removeEventListener("scrollend", handleScrollEnd);
  }, [chatId]);

  useLayoutEffect(() => {
    if (!scrollType) return;
    let top = 0;
    switch (scrollType) {
      case ScrollType.NewMessage:
        top = messagesRef.current?.scrollHeight ?? 0;
        top = top === 0 ? top : top - 1; // There are issues with exact measurement
        break;
      case ScrollType.RestoreChatPosition:
        top = chatsScrollMap.get(chatId) ?? 0;
        break;
      case ScrollType.LastMessage:
        top =
          (messagesRef.current?.scrollHeight ?? 0) -
          window.innerHeight +
          (inputContainerRef.current?.clientHeight ?? 0) +
          INPUT_BOTTOM_DISTANCE +
          12;
        top = Math.max(top, 0);
    }
    console.log("scrolling to: ", scrollType, top);
    containerRef.current?.scrollTo({
      top: top,
      behavior: "instant",
    });
    setScrollType(null);
  }, [scrollType, chatId]);

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
