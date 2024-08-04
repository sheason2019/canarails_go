import useApi from '@/common/use-api';
import { useParams } from '@modern-js/runtime/router';
import useSWR from 'swr';

export default function useAppVariants() {
  const { id } = useParams();
  const api = useApi();

  return useSWR(
    ['app-variants', id],
    () =>
      api.GET('/api/app-variant', {
        params: {
          query: {
            appId: Number(id),
          },
        },
      }),
    {
      suspense: true,
    },
  );
}
