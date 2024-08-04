import useApi from '@/common/use-api';
import { useParams } from '@modern-js/runtime/router';
import useSWR from 'swr';

export default function useApp() {
  const api = useApi();

  const { id } = useParams();
  return useSWR(
    ['app', id],
    () =>
      api.GET('/api/app/{id}', {
        params: {
          path: { id: Number(id) },
        },
      }),
    {
      suspense: true,
    },
  );
}
