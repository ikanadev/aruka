import { ChatService } from "@connect/chat/v1/services_pb";
import { ProviderService } from "@connect/provider/v1/services_pb";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

const transport = createConnectTransport({
  baseUrl: 'http://localhost:5000',
});

export const chatClient = createClient(ChatService, transport);
export const providerClient = createClient(ProviderService, transport);
