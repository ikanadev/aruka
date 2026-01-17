import { create } from "zustand";

interface FirstMessageStore {
  firstMessage: string | null;
  setFirstMessage: (message: string) => void;
  clearFirstMessage: () => void;
}

export const useFirstMessageStore = create<FirstMessageStore>((set) => ({
  firstMessage: null,
  setFirstMessage: (message) => set({ firstMessage: message }),
  clearFirstMessage: () => set({ firstMessage: null }),
}));
