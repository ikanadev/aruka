import { create } from "zustand";

interface ChatsScrollStore {
  scrollMap: Record<string, number>;
  setChatScroll: (chatId: string, scroll: number) => void;
}

export const useChatsScrollStore = create<ChatsScrollStore>((set) => ({
  scrollMap: {},
  setChatScroll: (chatId, scroll) =>
    set((state) => ({ scrollMap: { ...state.scrollMap, [chatId]: scroll } })),
}));
