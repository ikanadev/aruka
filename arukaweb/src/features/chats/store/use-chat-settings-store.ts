import { create } from "zustand";
import { persist } from "zustand/middleware";

interface ChatSettingsStore {
	modelId: string | null;
	customPrompt: string;
	saveSettings: (settings: { modelId: string; customPrompt: string }) => void;
}

export const useChatSettingsStore = create<ChatSettingsStore>()(
	persist(
		(set) => ({
			modelId: null,
			customPrompt: "",
			saveSettings: (settings) =>
				set({ modelId: settings.modelId, customPrompt: settings.customPrompt }),
		}),
		{ name: "chat-settings" },
	),
);
