import type { Message } from "@connect/arukabe/models/v1/message_pb";
import { create } from "zustand";

interface ChatMessagesState {
  messagesByChat: Record<string, Message[]>;
  setMessages: (chatId: string, messages: Message[]) => void;
  addMessage: (chatId: string, message: Message) => void;
  updateLastMessage: (chatId: string, message: Message) => void;
}

export const useChatMessagesStore = create<ChatMessagesState>((set) => ({
  messagesByChat: {},
  setMessages: (chatId, messages) =>
    set((state) => ({
      messagesByChat: { ...state.messagesByChat, [chatId]: messages },
    })),
  addMessage: (chatId, message) =>
    set((state) => ({
      messagesByChat: {
        ...state.messagesByChat,
        [chatId]: [...(state.messagesByChat[chatId] ?? []), message],
      },
    })),
  updateLastMessage: (chatId, message) =>
    set((state) => {
      const current = state.messagesByChat[chatId] ?? [];
      if (current.length === 0) {
        return { messagesByChat: { ...state.messagesByChat, [chatId]: [message] } };
      }
      return {
        messagesByChat: {
          ...state.messagesByChat,
          [chatId]: [...current.slice(0, -1), message],
        },
      };
    }),
}));

// Selectors
const EMPTY_MESSAGES: Message[] = [];
export const selectChatMessages = (chatId: string) => (state: ChatMessagesState) =>
  state.messagesByChat[chatId] ?? EMPTY_MESSAGES;
