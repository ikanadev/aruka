import { ChatService } from "@connect/arukabe/chat/v1/services_pb";
import { ProviderService } from "@connect/arukabe/provider/v1/services_pb";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

const transport = createConnectTransport({
  baseUrl: "http://192.168.0.9:5000",
});

export const chatClient = createClient(ChatService, transport);
export const providerClient = createClient(ProviderService, transport);
