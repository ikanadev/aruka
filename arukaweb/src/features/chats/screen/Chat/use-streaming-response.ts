import { chatClient } from "@common/utils/clients";
import type { Message } from "@connect/arukabe/models/v1/message_pb";
import { MessageRole } from "@connect/arukabe/models/v1/message_pb";
import { useChatMessagesStore } from "@features/chats/store/chat-messages-store";
import { createTextMessage } from "@features/chats/utils/create_text_message";
import { useRouter } from "@tanstack/react-router";
import { useCallback, useEffect, useState } from "react";

export function useStreamingResponse(chatId: string) {
  const router = useRouter();
  const addMessage = useChatMessagesStore((s) => s.addMessage);
  const [generatedResponse, setGeneratedResponse] = useState<Message | null>(null);
  const [responseError, setResponseError] = useState<string | null>(null);

  // Save generated response to store before navigating away
  useEffect(() => {
    const unsubscribe = router.subscribe("onBeforeNavigate", () => {
      if (generatedResponse === null) return;
      addMessage(chatId, generatedResponse);
      setGeneratedResponse(null);
    });
    return unsubscribe;
  }, [chatId, generatedResponse, router, addMessage]);

  const sendMessage = useCallback(
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

  const commitResponse = useCallback(() => {
    if (generatedResponse === null) return;
    addMessage(chatId, generatedResponse);
    setGeneratedResponse(null);
  }, [addMessage, chatId, generatedResponse]);

  const clearError = useCallback(() => setResponseError(null), []);

  return {
    generatedResponse,
    responseError,
    sendMessage,
    commitResponse,
    clearError,
  };
}
