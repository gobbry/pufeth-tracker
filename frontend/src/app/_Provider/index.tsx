import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Provider as UIProvider } from "@/components/ui/provider";

const queryClient = new QueryClient();

export const Provider = ({ children }: { children: React.ReactNode }) => {
  return (
    <UIProvider>
      <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
    </UIProvider>
  );
};
