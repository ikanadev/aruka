import { useEffect, useLayoutEffect, useRef, useState } from "react";
import { INPUT_BOTTOM_DISTANCE } from "./utils";

const ScrollType = {
  LastMessage: "LastMessage",
  NewMessage: "NewMessage",
  RestoreChatPosition: "RestoreChatPosition",
} as const;
type ScrollType = (typeof ScrollType)[keyof typeof ScrollType];

const chatsScrollMap = new Map<string, number>();

interface UseChatScrollOptions {
  chatId: string;
  isFetchedMessages: boolean;
  messagesCount: number;
}

export function useChatScroll(options: UseChatScrollOptions) {
  const { chatId, isFetchedMessages, messagesCount } = options;
  const containerRef = useRef<HTMLDivElement>(null);
  const messagesRef = useRef<HTMLDivElement>(null);
  const inputContainerRef = useRef<HTMLDivElement>(null);
  const [scrollType, setScrollType] = useState<ScrollType | null>(null);

  // Restore scroll position on revisit
  useEffect(() => {
    if (!chatsScrollMap.has(chatId)) return;
    setScrollType(ScrollType.RestoreChatPosition);
  }, [chatId]);

  // Scroll to last message on first load
  useEffect(() => {
    if (!isFetchedMessages) return;
    if (chatsScrollMap.has(chatId)) return;
    if (messagesCount <= 1) return;
    setScrollType(ScrollType.LastMessage);
  }, [chatId, isFetchedMessages, messagesCount]);

  // Track scroll position
  useEffect(() => {
    const el = containerRef.current;
    if (!el) return;
    const handleScrollEnd = () => chatsScrollMap.set(chatId, el.scrollTop);
    el.addEventListener("scrollend", handleScrollEnd);
    return () => el.removeEventListener("scrollend", handleScrollEnd);
  }, [chatId]);

  // Execute scroll
  useLayoutEffect(() => {
    if (!scrollType) return;
    let top = 0;
    switch (scrollType) {
      case ScrollType.NewMessage:
        top = messagesRef.current?.scrollHeight ?? 0;
        top = top === 0 ? top : top - 1;
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
    containerRef.current?.scrollTo({ top, behavior: "instant" });
    setScrollType(null);
  }, [scrollType, chatId]);

  const scrollToNewMessage = () => setScrollType(ScrollType.NewMessage);

  return {
    containerRef,
    messagesRef,
    inputContainerRef,
    scrollToNewMessage,
  };
}
