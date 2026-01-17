import { providerClient } from "@common/utils/clients";
import { useQuery } from "@tanstack/react-query";
import { commontQueryKeys } from "./common-query-keys";
import { ModelStatus } from "@connect/models/v1/model_pb";

export function useProviders(status: ModelStatus = ModelStatus.ACTIVE) {
  const query = useQuery({
    queryFn: async () => providerClient.listProviders({ status }),
    queryKey: commontQueryKeys.providers(),
  });

  const providers = query.data?.providers || [];

  return {
    providers,
    loadingProviders: query.isLoading,
    fetchingProviders: query.isFetching,
  };
}
