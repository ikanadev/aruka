import { useLayoutEffect, useRef } from "react";
import { useChatsScrollStore } from "../store/useChatsScrollStore";

export function useChatScroll(chatId: string) {
  const hasStoredScroll = useChatsScrollStore((s) => chatId in s.scrollMap);
  const chatScroll = useChatsScrollStore((s) => s.scrollMap[chatId] ?? 0);
  const setChatScroll = useChatsScrollStore((s) => s.setChatScroll);
  const ref = useRef<HTMLDivElement>(null);
  const isUserScrolling = useRef(false);
  const isRestoringScroll = useRef(false);

  // Scroll handler
  const onScroll = (e: React.UIEvent<HTMLDivElement, UIEvent>) => {
    if (isRestoringScroll.current) return;
    isUserScrolling.current = true;
    console.log('Set scroll ONSCROLL:', chatId, e.currentTarget.scrollTop);
    setChatScroll(chatId, e.currentTarget.scrollTop);
  };

  const handleSetScroll = (scroll: number) => {
    console.log('Set scroll', chatId, scroll);
    setChatScroll(chatId, scroll);
  };

  useLayoutEffect(() => {
    if (ref.current && !isUserScrolling.current) {
      isRestoringScroll.current = true;
      ref.current.scrollTo({
        top: chatScroll,
        behavior: "instant",
      });
      requestAnimationFrame(() => {
        isRestoringScroll.current = false;
      });
    }
    isUserScrolling.current = false;
  }, [chatScroll]);

  return {
    ref,
    chatScroll,
    hasStoredScroll,
    setChatScroll: handleSetScroll,
    onScroll,
  };
}
